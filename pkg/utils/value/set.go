package value

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/maargenton/go-testpredicate/pkg/utils/prettyprint"
)

// Set represents a unique set of values of any kind, and provides methods to
// performat standard set operation.
type Set map[interface{}]struct{}

// ReflectSet returns a newly constructed set from a collection type, or an
// error if the value is not a collection.
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

// Union returns a new set that is the union of the receiver with another set.
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

// Minus returns a new set that is the result of subtracting another set from the
// receiver.
func (lhs Set) Minus(rhs Set) (r Set) {
	r = Set{}
	for v := range lhs {
		if _, ok := rhs[v]; !ok {
			r[v] = struct{}{}
		}
	}
	return
}

// Intersect returns a new set that is the intersection between the receiver and
// another set.
func (lhs Set) Intersect(rhs Set) (r Set) {
	r = Set{}
	for v := range lhs {
		if _, ok := rhs[v]; ok {
			r[v] = struct{}{}
		}
	}
	return
}

// FormatSetValues returns a textual representation of the receiving set,
// potentially abbreviated, for the purpose of reporting discrepancies during
// unit-test.
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
