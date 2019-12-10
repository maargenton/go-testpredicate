// Package value defines helper functions to go from generic interface{} to
// specific numeric types, and compare them as either ordered or unordered
// values.
package value

import (
	"bytes"
	"fmt"
	"math"
	"reflect"
	"strings"
	"time"
)

// Synopsis:
//
// AsInt
// AsFloat
// AsUInt
//   Convert value to a preferred int or float type if possible
//
// CompareOrdered
// CompareUnordered
//   Compare two values to either find an order between them, or just check
//   for equality. Both functions are lenient, using numeric value conversion
//   and first difference for sequence comparison.
//
// MaxAbsoluteDifference
//   Returns the maximum absolute difference of all the components. Both
//   values must have the same shape and must be composed of values convertible
//   to float64 for comparison

// ---------------------------------------------------------------------------
// Helper functions to normalize numeric values into comparable type
// ---------------------------------------------------------------------------

// AsInt return the value as an int64 if possible
func AsInt(value interface{}) (int64, bool) {
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

// AsUInt return the value as an uint64 if possible
func AsUInt(value interface{}) (uint64, bool) {
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

// AsFloat return the value as a flaot64 if possible
func AsFloat(value interface{}) (float64, bool) {
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
// and slices of comparable orderable values. Values of different type or
// values of the same type that cannot be ordered return an error.
func CompareOrdered(lhs, rhs interface{}) (int, error) {

	if lhsInt, ok := AsInt(lhs); ok {
		if rhsInt, ok := AsInt(rhs); ok {
			return compareInt(lhsInt, rhsInt), nil
		}
	}

	if lhsUInt, ok := AsUInt(lhs); ok {
		if rhsUInt, ok := AsUInt(rhs); ok {
			return compareUInt(lhsUInt, rhsUInt), nil
		}
	}

	if lhsFloat, ok := AsFloat(lhs); ok {
		if rhsFloat, ok := AsFloat(rhs); ok {
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

	if lhsTime, ok := lhs.(time.Time); ok {
		if rhsTime, ok := rhs.(time.Time); ok {
			dt := lhsTime.Sub(rhsTime)
			return compareInt(int64(dt), 0), nil
		}
	}

	if isSliceComparable(lhs) && isSliceComparable(rhs) {
		return compareOrderedSlices(lhs, rhs)
	}

	ta := reflect.TypeOf(lhs)
	tb := reflect.TypeOf(rhs)
	if ta == tb {
		return 0, fmt.Errorf("values of type '%v' are not order comparable", ta)
	}
	return 0, fmt.Errorf("values of type '%v' and '%v' are not order comparable", ta, tb)
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

	if lhsInt, ok := AsInt(lhs); ok {
		if rhsInt, ok := AsInt(rhs); ok {
			return lhsInt == rhsInt, nil
		}
	}

	if lhsUInt, ok := AsUInt(lhs); ok {
		if rhsUInt, ok := AsUInt(rhs); ok {
			return lhsUInt == rhsUInt, nil
		}
	}

	if lhsFloat, ok := AsFloat(lhs); ok {
		if rhsFloat, ok := AsFloat(rhs); ok {
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
			"values of type '%T' and '%T' are never equal", lhs, rhs)
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
			return false, fmt.Errorf(
				"comparison of values at index %v failed, %v", i, err)
		}
		if !result {
			return false, nil
		}
	}

	return true, nil
}

// MaxAbsoluteDifference returns the maximum absolute difference of all the
// components. Both values must have the same shape and must be composed of
// values convertible to float64 for comparison
func MaxAbsoluteDifference(lhs, rhs interface{}) (float64, error) {

	a := reflect.ValueOf(lhs)
	b := reflect.ValueOf(rhs)
	ka := a.Kind()
	kb := b.Kind()
	if (ka == reflect.Slice || ka == reflect.Array) &&
		(kb == reflect.Slice || kb == reflect.Array) {

		na := a.Len()
		nb := b.Len()

		if na != nb {
			return 0, fmt.Errorf("value length (%v and %v) mismatched", na, nb)
		}

		max := 0.0
		for i := 0; i < na; i++ {
			va := a.Index(i).Interface()
			vb := b.Index(i).Interface()
			d, err := MaxAbsoluteDifference(va, vb)
			if err != nil {
				return 0, fmt.Errorf(
					"failed to compare values at index %v, %v", i, err)
			}

			if d > max {
				max = d
			}
		}

		return max, nil
	}

	if fa, ok := AsFloat(lhs); ok {
		if fb, ok := AsFloat(rhs); ok {
			return math.Abs(fa - fb), nil
		}
		return 0, fmt.Errorf(
			"value of type '%T' cannot be converted to float", rhs)
	}
	return 0, fmt.Errorf(
		"value of type '%T' cannot be converted to float", lhs)
}
