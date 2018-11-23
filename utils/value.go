package utils

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

// Synopsis:
//
// ValueAsInt
// ValueAsFloat
// ValueAsUInt
//   Convert value to a prefered int or float type if possible
//
// CompareOrdered
// CompareUnordered
//   Compare two values to either find an order between them, or just check
//   for equality. Both functions are lenient, using numeric value conversion
//   and first difference for sequence comparison.

// ---------------------------------------------------------------------------
// Helper functions to normalize numeric values into comparable type
// ---------------------------------------------------------------------------

// ValueAsInt return the value as an int64 if possible
func ValueAsInt(value interface{}) (int64, bool) {
	switch v := value.(type) {
	case int:
		return int64(v), true
	case int8:
		return int64(v), true
	case int16:
		return int64(v), true
	case int32:
		return int64(v), true
	case int64:
		return v, true

	case uint:
		if uint64(v) > (^uint64(0) >> 1) {
			return 0, false
		}
		return int64(v), true
	case uint8:
		return int64(v), true
	case uint16:
		return int64(v), true
	case uint32:
		return int64(v), true
	case uint64:
		if v > (^uint64(0) >> 1) {
			return 0, false
		}
		return int64(v), true

	default:
		return 0, false
	}
}

// ValueAsUInt return the value as an uint64 if possible
func ValueAsUInt(value interface{}) (uint64, bool) {
	switch v := value.(type) {
	case int:
		if v < 0 {
			return 0, false
		}
		return uint64(v), true
	case int8:
		if v < 0 {
			return 0, false
		}
		return uint64(v), true
	case int16:
		if v < 0 {
			return 0, false
		}
		return uint64(v), true
	case int32:
		if v < 0 {
			return 0, false
		}
		return uint64(v), true
	case int64:
		if v < 0 {
			return 0, false
		}
		return uint64(v), true

	case uint:
		return uint64(v), true
	case uint8:
		return uint64(v), true
	case uint16:
		return uint64(v), true
	case uint32:
		return uint64(v), true
	case uint64:
		return v, true
	case uintptr:
		return uint64(v), true

	default:
		return 0, false
	}
}

// ValueAsFloat return the value as a flaot64 if possible
func ValueAsFloat(value interface{}) (float64, bool) {
	switch v := value.(type) {
	case float32:
		return float64(v), true
	case float64:
		return v, true

	case int8:
		return float64(v), true
	case int16:
		return float64(v), true
	case int32:
		return float64(v), true

	case uint8:
		return float64(v), true
	case uint16:
		return float64(v), true
	case uint32:
		return float64(v), true

	case int:
		if v >= 0 && uint64(v) > uint64(1<<53) ||
			v < 0 && uint64(-v) > uint64(1<<53) {
			return 0, false
		}
		return float64(v), true
	case int64:
		if v >= 0 && uint64(v) > uint64(1<<53) ||
			v < 0 && uint64(-v) > uint64(1<<53) {
			return 0, false
		}
		return float64(v), true

	case uint:
		if uint64(v) > uint64(1<<53) {
			return 0, false
		}
		return float64(v), true
	case uint64:
		if v > uint64(1<<53) {
			return 0, false
		}
		return float64(v), true

	default:
		return 0, false
	}
}

// ---------------------------------------------------------------------------
// ---------------------------------------------------------------------------

// CompareOrdered returns the result of the comparison of two values of compatible
// type that can be ordered, including ints, floats, strings, byte slices,
// and slices of comparable orderable values. Values of different tyype or
// values of ther same type that cannot be ordered return an error.
func CompareOrdered(lhs, rhs interface{}) (int, error) {

	if lhsInt, ok := ValueAsInt(lhs); ok {
		if rhsInt, ok := ValueAsInt(rhs); ok {
			return compareInt(lhsInt, rhsInt), nil
		}
	}

	if lhsUInt, ok := ValueAsUInt(lhs); ok {
		if rhsUInt, ok := ValueAsUInt(rhs); ok {
			return compareUInt(lhsUInt, rhsUInt), nil
		}
	}

	if lhsFloat, ok := ValueAsFloat(lhs); ok {
		if rhsFloat, ok := ValueAsFloat(rhs); ok {
			return compareFloat(lhsFloat, rhsFloat), nil
		}
	}

	if lhsStr, ok := lhs.(string); ok {
		if rhsStr, ok := rhs.(string); ok {
			return strings.Compare(lhsStr, rhsStr), nil
		}
	}

	if lhsBytes, ok := lhs.([]byte); ok {
		if rhsBytes, ok := rhs.([]byte); ok {
			return bytes.Compare(lhsBytes, rhsBytes), nil
		}
	}

	if isSliceComparable(lhs) && isSliceComparable(rhs) {
		return compareOrderedSlices(lhs, rhs)
	}

	ta := reflect.TypeOf(lhs)
	tb := reflect.TypeOf(rhs)
	if ta == tb {
		return 0, fmt.Errorf("values of type %v are not order comparable", ta)
	}
	return 0, fmt.Errorf("values of type %v and %v are not order comparable", ta, tb)
}

func compareInt(lhs, rhs int64) int {
	if lhs == rhs {
		return 0
	}
	if lhs < rhs {
		return -1
	}
	return 1
}

func compareUInt(lhs, rhs uint64) int {
	if lhs == rhs {
		return 0
	}
	if lhs < rhs {
		return -1
	}
	return 1
}

func compareFloat(lhs, rhs float64) int {
	if lhs == rhs {
		return 0
	}
	if lhs < rhs {
		return -1
	}
	return 1
}

func isSliceComparable(value interface{}) bool {
	k := reflect.ValueOf(value).Kind()
	if k == reflect.Slice || k == reflect.Array || k == reflect.String {
		return true
	}
	return false
}

func compareOrderedSlices(lhs, rhs interface{}) (int, error) {

	a := reflect.ValueOf(lhs)
	b := reflect.ValueOf(rhs)
	na := a.Len()
	nb := b.Len()
	for i := 0; i < na && i < nb; i++ {
		va := a.Index(i).Interface()
		vb := b.Index(i).Interface()

		result, err := CompareOrdered(va, vb)
		if err != nil {
			return 0, fmt.Errorf("comparison of values at index %v failed, %v", i, err)
		}
		if result != 0 {
			return result, nil
		}
	}

	return compareInt(int64(na), int64(nb)), nil
}

// ---------------------------------------------------------------------------
// ---------------------------------------------------------------------------

// CompareUnordered returns the result of the equality comparison of two values
// of compatible type, including ints, floats, strings and slices of
// comparable values. Other values are compared with reflect.DeepEqual().
// Values of different tyype that cannot be compared return an error.
func CompareUnordered(lhs, rhs interface{}) (bool, error) {

	if lhsInt, ok := ValueAsInt(lhs); ok {
		if rhsInt, ok := ValueAsInt(rhs); ok {
			return lhsInt == rhsInt, nil
		}
	}

	if lhsUInt, ok := ValueAsUInt(lhs); ok {
		if rhsUInt, ok := ValueAsUInt(rhs); ok {
			return lhsUInt == rhsUInt, nil
		}
	}

	if lhsFloat, ok := ValueAsFloat(lhs); ok {
		if rhsFloat, ok := ValueAsFloat(rhs); ok {
			return lhsFloat == rhsFloat, nil
		}
	}

	if lhsStr, ok := lhs.(string); ok {
		if rhsStr, ok := rhs.(string); ok {
			return lhsStr == rhsStr, nil
		}
	}

	if isSliceComparable(lhs) && isSliceComparable(rhs) {
		return compareUnorderedSlices(lhs, rhs)
	}

	v1 := reflect.ValueOf(lhs)
	v2 := reflect.ValueOf(rhs)
	if v1.Type() != v2.Type() {
		return false, fmt.Errorf(
			"values of type %T and %T are never equal", lhs, rhs)
	}

	return reflect.DeepEqual(lhs, rhs), nil
}

func compareUnorderedSlices(lhs, rhs interface{}) (bool, error) {

	a := reflect.ValueOf(lhs)
	b := reflect.ValueOf(rhs)
	na := a.Len()
	nb := b.Len()
	if na != nb {
		return false, nil
	}

	for i := 0; i < na && i < nb; i++ {
		va := a.Index(i).Interface()
		vb := b.Index(i).Interface()

		result, err := CompareUnordered(va, vb)
		if err != nil {
			return false, fmt.Errorf("comparison of values at index %v failed, %v", i, err)
		}
		if !result {
			return false, nil
		}
	}

	return true, nil
}
