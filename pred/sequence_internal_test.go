package pred

import (
	"reflect"
	"testing"
)

func TestPreCheckSubsequence(t *testing.T) {
	v1 := []int{123, 455, 789}
	v2 := []int{123, 455}

	err := preCheckSubsequence(reflect.ValueOf(uint64(1)), reflect.ValueOf(v1))
	if err == nil || err.Error() != "value 0x1 of type uint64 is not a sequence" {
		t.Errorf("\nunexpected error: %v", err)
	}

	err = preCheckSubsequence(reflect.ValueOf(v1), reflect.ValueOf(uint64(1)))
	if err == nil || err.Error() != "value 0x1 of type uint64 is not a sequence" {
		t.Errorf("\nunexpected error: %v", err)
	}

	err = preCheckSubsequence(reflect.ValueOf(v1), reflect.ValueOf(v2))
	if err != nil {
		t.Errorf("\nunexpected error: %v", err)
	}
}

func TestIndexOfSubsequence(t *testing.T) {
	v1 := []int{1, 2, 3, 4, 5}

	v2 := []int{3, 4}
	i := indexOfSubsequence(reflect.ValueOf(v1), reflect.ValueOf(v2))
	if i != 2 {
		t.Errorf("index: %v", i)
	}

	v3 := []int{3, 4, 5}
	i = indexOfSubsequence(reflect.ValueOf(v1), reflect.ValueOf(v3))
	if i != 2 {
		t.Errorf("index: %v", i)
	}

	v4 := []int{3, 5}
	i = indexOfSubsequence(reflect.ValueOf(v1), reflect.ValueOf(v4))
	if i != -1 {
		t.Errorf("index: %v", i)
	}

	v5 := []int{1, 2, 3, 4, 5, 6}
	i = indexOfSubsequence(reflect.ValueOf(v1), reflect.ValueOf(v5))
	if i != -1 {
		t.Errorf("index: %v", i)
	}
}

func TestIndexOfSubsequenceWithString(t *testing.T) {
	v1 := "abcde"

	v2 := "cd"
	i := indexOfSubsequence(reflect.ValueOf(v1), reflect.ValueOf(v2))
	if i != 2 {
		t.Errorf("index: %v", i)
	}

	v3 := "cde"
	i = indexOfSubsequence(reflect.ValueOf(v1), reflect.ValueOf(v3))
	if i != 2 {
		t.Errorf("index: %v", i)
	}

	v4 := "ce"
	i = indexOfSubsequence(reflect.ValueOf(v1), reflect.ValueOf(v4))
	if i != -1 {
		t.Errorf("index: %v", i)
	}

	v5 := "abcdef"
	i = indexOfSubsequence(reflect.ValueOf(v1), reflect.ValueOf(v5))
	if i != -1 {
		t.Errorf("index: %v", i)
	}
}
