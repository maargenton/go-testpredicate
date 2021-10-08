package value_test

import (
	"reflect"
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/utils/value"
)

func TestIsSequenceType(t *testing.T) {
	var tcs = []struct {
		v        interface{}
		expected bool
	}{
		{"abc", true},
		{[]int{1, 2, 3}, true},
		{[3]int{}, true},
		{123, false},
		{struct{}{}, false},
	}

	for _, tc := range tcs {
		actual := value.IsSequenceType(reflect.ValueOf(tc.v))
		if actual != tc.expected {
			t.Errorf("\nexpected: %v\nactual:   %v\nvalue:    %#+v",
				tc.expected, actual, tc.v)
		}
	}
}

func TestPreCheckSubsequence(t *testing.T) {
	a := reflect.ValueOf([]int{1, 2, 3})
	b := reflect.ValueOf(123)

	if err := value.PreCheckSubsequence(a, b); err == nil {
		t.Errorf("expected error, none produced")
	} else {
		s := err.Error()
		if s != "value of type 'int' is not a sequence" {
			t.Errorf("error mismatch, was: %v", err)
		}
	}

	if err := value.PreCheckSubsequence(b, a); err == nil {
		t.Errorf("expected error, none produced")
	} else {
		s := err.Error()
		if s != "value of type 'int' is not a sequence" {
			t.Errorf("error mismatch, was: %v", err)
		}
	}

	if err := value.PreCheckSubsequence(a, a); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestIndexOfSubsequence(t *testing.T) {
	a := reflect.ValueOf([]int{1, 2, 3, 4, 5})
	b := reflect.ValueOf([]int{3, 4})
	c := reflect.ValueOf([]int{3, 4, 5})
	d := reflect.ValueOf([]int{3, 4, 5, 6})

	if i := value.IndexOfSubsequence(a, b); i != 2 {
		t.Errorf("\nfailed to find sub-sequence, index: %v", i)
	}
	if i := value.IndexOfSubsequence(a, c); i != 2 {
		t.Errorf("\nfailed to find sub-sequence, index: %v", i)
	}
	if i := value.IndexOfSubsequence(a, d); i != -1 {
		t.Errorf("\nfound unexpected sub-sequence, index: %v", i)
	}
}
