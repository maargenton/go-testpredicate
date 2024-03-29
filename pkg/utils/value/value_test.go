package value_test

import (
	"math"
	"testing"
	"time"

	"github.com/maargenton/go-testpredicate/pkg/utils/prettyprint"
	"github.com/maargenton/go-testpredicate/pkg/utils/value"
)

func TestAsInt(t *testing.T) {
	var inputs = []struct {
		value    interface{}
		expected int64
		success  bool
	}{
		{int8(127), 127, true},
		{int8(-128), -128, true},
		{int16(12700), 12700, true},
		{int32(-1280000), -1280000, true},

		{uint8(127), 127, true},
		{uint16(12700), 12700, true},
		{uint32(1280000), 1280000, true},

		{uint(123456), 123456, true},
		{uint(0x7FFFFFFFFFFFFFFF), 0x7FFFFFFFFFFFFFFF, true},
		{uint(0x8000000000000000), 0, false},

		{uint64(123456), 123456, true},
		{uint64(0x7FFFFFFFFFFFFFFF), 0x7FFFFFFFFFFFFFFF, true},
		{uint64(0x8000000000000000), 0, false},

		{float32(1.234), 0, false},
		{float64(1.234), 0, false},
		{"1.234", 0, false},
	}

	for _, in := range inputs {
		v, success := value.AsInt(in.value)
		if success != in.success {
			if in.success {
				t.Errorf("Expected successful conversion for value %v", in.value)
			} else {
				t.Errorf("Expected failed conversion for value %v", in.value)
			}
		}
		if in.success && v != in.expected {
			t.Errorf(
				"Converted value %T(%#v) for %T(%#v) doesn't match expected %#v",
				v, v, in.value, in.value, in.expected)
		}
	}
}

func TestAsUInt(t *testing.T) {
	var inputs = []struct {
		value    interface{}
		expected uint64
		success  bool
	}{
		{int8(127), 127, true},
		{int8(-128), 0, false},
		{int16(12700), 12700, true},
		{int16(-12700), 0, false},
		{int32(1280000), 1280000, true},
		{int32(-1280000), 0, false},
		{int64(1280000), 1280000, true},
		{int64(-1280000), 0, false},
		{int(1280000), 1280000, true},
		{int(-1280000), 0, false},

		{uint8(127), 127, true},
		{uint16(12700), 12700, true},
		{uint32(1280000), 1280000, true},

		{uint64(123456), 123456, true},
		{uint64(0x7FFFFFFFFFFFFFFF), 0x7FFFFFFFFFFFFFFF, true},
		{uint64(0x8000000000000000), 0x8000000000000000, true},
		{uint64(0xFFFFFFFFFFFFFFFF), 0xFFFFFFFFFFFFFFFF, true},

		{uint(0xFFFFFFFFFFFFFFFF), 0xFFFFFFFFFFFFFFFF, true},
		{uintptr(0xFFFFFFFFFFFFFFFF), 0xFFFFFFFFFFFFFFFF, true},

		{float32(1.234), 0, false},
		{float64(1.234), 0, false},
		{"1.234", 0, false},
	}

	for _, in := range inputs {
		v, success := value.AsUInt(in.value)
		if success != in.success {
			if in.success {
				t.Errorf("Expected successful conversion for value %v", in.value)
			} else {
				t.Errorf("Expected failed conversion for value %v", in.value)
			}
		}
		if in.success && v != in.expected {
			t.Errorf(
				"Converted value %T(%#v) for %T(%#v) doesn't match expected %#v",
				v, v, in.value, in.value, in.expected)
		}
	}
}

func TestAsFloat(t *testing.T) {
	var inputs = []struct {
		value    interface{}
		expected float64
		success  bool
	}{
		{int8(127), 127, true},
		{int8(-128), -128, true},
		{int16(12700), 12700, true},
		{int32(-1280000), -1280000, true},
		{int64(123456), 123456, true},

		{uint8(127), 127, true},
		{uint16(12700), 12700, true},
		{uint32(123456), 123456, true},
		{uint64(12345600), 12345600, true},

		{int(0x001FFFFFFFFFFFFF), 0x001FFFFFFFFFFFFF, true},
		{int(0x0020000000000000), 0x0020000000000000, true},
		{int(0x0020000000000001), 0x0020000000000001, false},

		{int64(0x001FFFFFFFFFFFFF), 0x001FFFFFFFFFFFFF, true},
		{int64(0x0020000000000000), 0x0020000000000000, true},
		{int64(0x0020000000000001), 0x0020000000000001, false},

		{-int64(0x001FFFFFFFFFFFFF), -0x001FFFFFFFFFFFFF, true},
		{-int64(0x0020000000000000), -0x0020000000000000, true},
		{-int64(0x0020000000000001), -0x0020000000000001, false},

		{uint(0x001FFFFFFFFFFFFF), 0x001FFFFFFFFFFFFF, true},
		{uint(0x0020000000000000), 0x0020000000000000, true},
		{uint(0x0020000000000001), 0x0020000000000001, false},

		{uint64(0x001FFFFFFFFFFFFF), 0x001FFFFFFFFFFFFF, true},
		{uint64(0x0020000000000000), 0x0020000000000000, true},
		{uint64(0x0020000000000001), 0x0020000000000001, false},

		{float32(1.235), float64(float32(1.235)), true},
		{float64(1.235), 1.235, true},
		{"1.234", 0, false},
	}

	for _, in := range inputs {
		v, success := value.AsFloat(in.value)
		if success != in.success {
			if in.success {
				t.Errorf("Expected successful conversion for value %v", in.value)
			} else {
				t.Errorf("Expected failed conversion for value %v", in.value)
			}
		}
		if in.success && v != in.expected {
			t.Errorf(
				"Converted value %T(%#v) for %T(%#v) doesn't match expected %#v",
				v, v, in.value, in.value, in.expected)
		}

		if iv, integerType := value.AsInt(in.value); integerType && in.success {
			if iv != int64(in.expected) {
				t.Errorf("Value %v wasn't preserved through float conversion", v)

			}
		}
	}
}

func TestComapareOrdered(t *testing.T) {
	var inputs = []struct {
		lhs, rhs interface{}
		result   int
		err      bool
	}{
		{1, 1, 0, false},
		{1, 2, -1, false},
		{2, 1, 1, false},

		{1, uint64(0xF000000000000000), -1, false},
		{uint64(0xF000000000000000), uint64(0xF000000000000000), 0, false},
		{uint64(0xF000000000000000), 1, 1, false},
		{-1, uint64(0xF000000000000000), 0, true},

		{1, 1.5, -1, false},
		{2, 1.5, 1, false},
		{2, 2.0, 0, false},

		{"aa", "aaa", -1, false},
		{"abc", "abd", -1, false},
		{"abc", "abc", 0, false},
		{"abd", "abc", 1, false},

		{"aa", []byte{97, 98}, -1, false},
		{"aa", []byte{97, 97}, 0, false},
		{"aa", []byte{97, 96}, 1, false},
		{"aa", []int{97, 98}, -1, false},
		{"aa", []int{97, 97}, 0, false},
		{"aa", []int{97, 96}, 1, false},

		{[]byte{1, 2, 3}, []byte{1, 2, 4}, -1, false},
		{[]byte{1, 2, 3}, []byte{1, 2, 3}, 0, false},
		{[]byte{1, 2, 3}, []byte{1, 2, 2}, 1, false},

		{123, struct{ a int }{123}, 0, true},

		{time.Unix(123, 0), time.Unix(124, 0), -1, false},
		{time.Unix(124, 0), time.Unix(123, 0), 1, false},
		{time.Unix(124, 0), time.Unix(124, 0), 0, false},
		{time.Unix(124, 0), 124, 0, true},
		{124, time.Unix(124, 0), 0, true},

		{
			[]int{123, 456, 789},
			[]interface{}{123, struct{ a int }{456}, 789},
			0, true,
		},
		{struct{ a int }{123}, struct{ a int }{123}, 0, true},
	}

	for _, input := range inputs {
		r, err := value.CompareOrdered(input.lhs, input.rhs)
		if input.err && err == nil {
			t.Errorf(
				"\nexpected CompareOrdered(a, b) to return an error"+
					"\na: %v"+
					"\nb: %v",
				prettyprint.FormatValue(input.lhs),
				prettyprint.FormatValue(input.rhs))
		} else if !input.err && err != nil {
			t.Errorf(
				"\nCompareOrdered(a, b) returned an error"+
					"\nerror: %v"+
					"\na: %v"+
					"\nb: %v",
				err,
				prettyprint.FormatValue(input.lhs),
				prettyprint.FormatValue(input.rhs))
		}

		if r != input.result {
			t.Errorf(
				"\nexpected CompareOrdered(a, b) = %d, actual = %v"+
					"\na: %v"+
					"\nb: %v",
				input.result, r,
				prettyprint.FormatValue(input.lhs),
				prettyprint.FormatValue(input.rhs),
			)
		}
	}
}

func TestComapareUnordered(t *testing.T) {
	var inputs = []struct {
		lhs, rhs interface{}
		result   bool
		err      bool
	}{
		{1, 1, true, false},
		{1, 2, false, false},
		{2, 1, false, false},

		{1, uint64(0xF000000000000000), false, false},
		{uint64(0xF000000000000000), uint64(0xF000000000000000), true, false},
		{uint64(0xF000000000000000), 1, false, false},
		{-1, uint64(0xF000000000000000), false, true},

		{1, 1.5, false, false},
		{2, 1.5, false, false},
		{2, 2.0, true, false},

		{"aa", "aaa", false, false},
		{"abc", "abd", false, false},
		{"abc", "abc", true, false},
		{"abd", "abc", false, false},

		{"aa", []byte{97, 98}, false, false},
		{"aa", []byte{97, 97}, true, false},
		{"aa", []byte{97, 96}, false, false},
		{"aa", []int{97, 98}, false, false},
		{"aa", []int{97, 97}, true, false},
		{"aa", []int{97, 96}, false, false},

		{[]byte{1, 2, 3}, []byte{1, 2}, false, false},
		{[]byte{1, 2, 3}, []byte{1, 2, 4}, false, false},
		{[]byte{1, 2, 3}, []byte{1, 2, 3}, true, false},
		{[]byte{1, 2, 3}, []byte{1, 2, 2}, false, false},

		{123, struct{ a int }{123}, false, true},

		{
			[]int{123, 456, 789},
			[]interface{}{123, struct{ a int }{456}, 789},
			false, true,
		},

		{struct{ a int }{123}, struct{ a int }{122}, false, false},
		{struct{ a int }{123}, struct{ a int }{123}, true, false},
		{struct{ a int }{123}, struct{ a int }{124}, false, false},
	}

	for _, input := range inputs {
		r, err := value.CompareUnordered(input.lhs, input.rhs)
		if input.err && err == nil {
			t.Errorf(
				"\nexpected CompareOrdered(%#+v, %#+v) to return an error",
				input.lhs, input.rhs)
		} else if !input.err && err != nil {
			t.Errorf(
				"\nCompareOrdered(%#+v, %#+v) returned error,\n%v",
				input.lhs, input.rhs, err)
		}

		if r != input.result {
			t.Errorf(
				"\nexpected CompareOrdered(%#+v, %#+v) = %v,\nactual = %v",
				input.lhs, input.rhs, input.result, r)
		}
	}
}

func TestMaxAbsoluteDifference(t *testing.T) {
	var inputs = []struct {
		lhs, rhs interface{}
		diff     float64
	}{
		{[3]float64{1, 2, 3}, [3]float64{1, 2, 3}, 0},
		{[3]float64{1, 2, 3}, [3]float64{1.1, 2, 3}, 0.1},
		{[3]float64{1, 2, 3}, [3]float64{1.1, 2.2, 3.3}, 0.3},

		{[]float64{1, 2, 3}, []float64{1, 2, 3}, 0},
		{[]float64{1, 2, 3}, []float64{1.1, 2, 3}, 0.1},
		{[]float64{1, 2, 3}, []float64{1.1, 2.2, 3.3}, 0.3},

		{
			[][]float64{
				{1, 2, 3},
				{2, 3, 4}},
			[][]float64{
				{1, 2, 3},
				{2, 3, 4},
			},
			0,
		},

		{
			[][]float64{
				{1, 2, 3},
				{2, 3, 4}},
			[][]float64{
				{1.1, 2.2, 3.3},
				{2.2, 3.3, 4.4},
			},
			0.4,
		},
	}

	for _, input := range inputs {
		r, err := value.MaxAbsoluteDifference(input.lhs, input.rhs)
		if err != nil {
			t.Error(err)
		}
		if math.Abs(r-input.diff) > 0.00001 {
			t.Errorf("expected difference: %v\nactual:%v", input.diff, r)
		}
	}
}

func TestMaxAbsoluteDifferenceErrors(t *testing.T) {
	var inputs = []struct {
		lhs, rhs interface{}
		err      string
	}{
		{
			[]float64{1, 2, 3},
			[]float64{1, 2, 3, 4},
			"value length (3 and 4) mismatched",
		},
		{
			[]interface{}{1, 2, 3},
			[]interface{}{1, "2", 3},
			"failed to compare values at index 1, value of type 'string' cannot be converted to float",
		},
		{
			[]interface{}{"1", 2, 3},
			[]interface{}{1, 2, 3},
			"failed to compare values at index 0, value of type 'string' cannot be converted to float",
		},
	}

	for _, input := range inputs {
		_, err := value.MaxAbsoluteDifference(input.lhs, input.rhs)
		if err == nil {
			t.Errorf(
				"\nexpected error for MaxAbsoluteDifference(%#v, %#v)",
				input.lhs, input.rhs)
		}
		if err.Error() != input.err {
			t.Errorf("\nunexpected error for MaxAbsoluteDifference(...):\n%v",
				err)
		}
	}
}
