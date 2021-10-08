package value_test

import (
	"strings"
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/utils/value"
)

func TestReflectSet(t *testing.T) {
	var a = []int{1, 2, 3}
	var aa, err = value.ReflectSet(a)

	if err != nil {
		t.Errorf("\nunexpected error:\n%v", err)
	}
	if len(aa) != 3 {
		t.Errorf("\nlen(aa) = %v", len(aa))
	}
}

func TestReflectSetError(t *testing.T) {
	var a = 123
	var _, err = value.ReflectSet(a)

	if err == nil {
		t.Errorf("\nno error returned by ReflectSet on invalid type")
	}
}

func TestSetUnion(t *testing.T) {
	a, err := value.ReflectSet([]int{1, 2, 3})
	if err != nil {
		t.Errorf("\nunexpected error:\n%v", err)
	}

	b, err := value.ReflectSet([]int{3, 4, 5})
	if err != nil {
		t.Errorf("\nunexpected error:\n%v", err)
	}

	c := a.Union(b)
	if len(c) != 5 {
		t.Errorf("\nlen(c) = %v", len(c))
	}
}

func TestSetMinus(t *testing.T) {
	a, err := value.ReflectSet([]int{1, 2, 3})
	if err != nil {
		t.Errorf("\nunexpected error:\n%v", err)
	}

	b, err := value.ReflectSet([]int{3, 4, 5})
	if err != nil {
		t.Errorf("\nunexpected error:\n%v", err)
	}

	c := a.Minus(b)
	if len(c) != 2 {
		t.Errorf("\nlen(c) = %v", len(c))
	}
}

func TestSetIntersect(t *testing.T) {
	a, err := value.ReflectSet([]int{1, 2, 3})
	if err != nil {
		t.Errorf("\nunexpected error:\n%v", err)
	}

	b, err := value.ReflectSet([]int{3, 4, 5})
	if err != nil {
		t.Errorf("\nunexpected error:\n%v", err)
	}

	c := a.Intersect(b)
	if len(c) != 1 {
		t.Errorf("\nlen(c) = %v", len(c))
	}
}

func TestFormatSetValue(t *testing.T) {
	var a []int
	for i := 0; i < 50; i++ {
		a = append(a, i)
	}
	aa, _ := value.ReflectSet(a)
	s := value.FormatSetValues(aa)

	if len(s) > 60 {
		t.Errorf("\nset was not truncated:\n%v", s)
	}
	if !strings.HasSuffix(s, ", ...") {
		t.Errorf("\ntruncated set does not end with ellipsis:\n%v", s)
	}
}
