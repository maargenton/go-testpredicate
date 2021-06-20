package value

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/maargenton/go-testpredicate/pkg/prettyprint"
)

type Set map[interface{}]struct{}

func ReflectSet(value interface{}) (s Set, err error) {
	s = Set{}
	v := reflect.ValueOf(value)
	k := v.Kind()
	if !(k == reflect.Slice || k == reflect.Array || k == reflect.String) {
		err = fmt.Errorf(
			"value of type '%T' is not an indexable collection", value)
		return
	}
	for i, n := 0, v.Len(); i < n; i++ {
		s[v.Index(i).Interface()] = struct{}{}
	}
	return
}

func (lhs Set) Union(rhs Set) (r Set) {
	r = Set{}
	for v := range lhs {
		r[v] = struct{}{}
	}
	for v := range rhs {
		r[v] = struct{}{}
	}
	return
}

func (lhs Set) Minus(rhs Set) (r Set) {
	r = Set{}
	for v := range lhs {
		if _, ok := rhs[v]; !ok {
			r[v] = struct{}{}
		}
	}
	return
}

func (lhs Set) Intersect(rhs Set) (r Set) {
	r = Set{}
	for v := range lhs {
		if _, ok := rhs[v]; ok {
			r[v] = struct{}{}
		}
	}
	return
}

func FormatSetValues(s Set) string {
	var buf strings.Builder
	for v := range s {
		if buf.Len() < 50 {
			if buf.Len() > 0 {
				buf.WriteString(", ")
			}
			buf.WriteString(prettyprint.FormatValue(v))
		} else {
			buf.WriteString(", ...")
			break
		}
	}
	return buf.String()
}
