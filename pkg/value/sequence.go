package value

import (
	"fmt"
	"reflect"
)

// IsSequenceType checks if a value is a kind of sequence
func IsSequenceType(v reflect.Value) bool {
	var k = v.Kind()
	return k == reflect.Array || k == reflect.Slice || k == reflect.String
}

// PreCheckSubsequence checks if a pair of values are suitable for use with
// IndexOfSubsequence().
func PreCheckSubsequence(v1, v2 reflect.Value) error {

	if !IsSequenceType(v1) {
		return fmt.Errorf("value of type '%T' is not a sequence",
			v1.Interface())
	}
	if !IsSequenceType(v2) {
		return fmt.Errorf("value of type '%T' is not a sequence",
			v2.Interface())
	}
	return nil
}

// IndexOfSubsequence finds the index of a sub-sequence `sub` in the sequence
// `seq`, or returns -1. Both arguments must be suitable sequences.
func IndexOfSubsequence(seq, sub reflect.Value) int {
	l1, l2 := seq.Len(), sub.Len()
	for i := 0; i <= l1-l2; i++ {
		allEq := true
		for j := 0; j < l2; j++ {
			v1, v2 := seq.Index(i+j), sub.Index(j)
			eq, _ := CompareUnordered(v1.Interface(), v2.Interface())
			if !eq {
				allEq = false
				break
			}
		}
		if allEq {
			return i
		}
	}

	return -1
}
