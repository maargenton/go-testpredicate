package prettyprint

import (
	"fmt"
	"testing"
)

func TestMin(t *testing.T) {
	var tcs = []struct {
		a, b, v int
	}{
		{1, 2, 1},
		{2, 1, 1},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprintf("min(%v, %v) == %v", tc.a, tc.b, tc.v), func(t *testing.T) {
			v := min(tc.a, tc.b)
			if v != tc.v {
				t.Errorf("unexpected value: %v", v)
			}
		})
	}
}

func TestMax(t *testing.T) {
	var tcs = []struct {
		a, b, v int
	}{
		{1, 2, 2},
		{2, 1, 2},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprintf("max(%v, %v) == %v", tc.a, tc.b, tc.v), func(t *testing.T) {
			v := max(tc.a, tc.b)
			if v != tc.v {
				t.Errorf("unexpected value: %v", v)
			}
		})
	}
}

func TestWrapString(t *testing.T) {
	var s = "Name:     \"then the command is not run\","

	// Test no panic, no infinite loop, even when w <= 0
	wrapString(s, 0)
	wrapString(s, -1)
}
