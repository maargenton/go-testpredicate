package impl

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/maargenton/go-testpredicate/pkg/internal/predicate"
	"github.com/maargenton/go-testpredicate/pkg/prettyprint"
)

// IsEqualSet tests if two containers contain the same set of values,
// independently of order.
func IsEqualSet(rhs interface{}) (desc string, f predicate.PredicateFunc) {
	desc = fmt.Sprintf("set({}) == %v", prettyprint.FormatValue(rhs))
	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		lhsSet, err := ReflectSet(v)
		if err != nil {
			return
		}
		rhsSet, err := ReflectSet(rhs)
		if err != nil {
			return
		}

		extra := lhsSet.Minus(rhsSet)
		missing := rhsSet.Minus(lhsSet)
		r = len(extra) == 0 && len(missing) == 0
		if !r {
			ctx = []predicate.ContextValue{
				{Name: "extra values", Value: FormatSetValues(extra), Pre: true},
				{Name: "missing values", Value: FormatSetValues(missing), Pre: true},
			}
		}
		return
	}
	return
}

// IsDisjointSetFrom tests if two containers contain no common values
func IsDisjointSetFrom(rhs interface{}) (desc string, f predicate.PredicateFunc) {
	desc = fmt.Sprintf("set({}) ∩ %v == ∅", prettyprint.FormatValue(rhs))
	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		lhsSet, err := ReflectSet(v)
		if err != nil {
			return
		}
		rhsSet, err := ReflectSet(rhs)
		if err != nil {
			return
		}

		intersection := lhsSet.Intersect(rhsSet)
		r = len(intersection) == 0
		if !r {
			ctx = []predicate.ContextValue{
				{Name: "common values", Value: FormatSetValues(intersection), Pre: true},
			}
		}
		return
	}
	return
}

// IsSubsetOf tests if the value under test is a subset of the reference value.
// Both values must be containers and are treated as unordered sets.
func IsSubsetOf(rhs interface{}) (desc string, f predicate.PredicateFunc) {
	desc = fmt.Sprintf("set({}) ⊂ %v", prettyprint.FormatValue(rhs))
	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		lhsSet, err := ReflectSet(v)
		if err != nil {
			return
		}
		rhsSet, err := ReflectSet(rhs)
		if err != nil {
			return
		}

		diff := lhsSet.Minus(rhsSet)
		r = len(diff) == 0
		if !r {
			ctx = []predicate.ContextValue{
				{Name: "extra values", Value: FormatSetValues(diff), Pre: true},
			}
		}
		return
	}
	return
}

// IsSupersetOf tests if the value under test is a superset of the reference
// value. Both values must be containers and are treated as unordered sets.
func IsSupersetOf(rhs interface{}) (desc string, f predicate.PredicateFunc) {
	desc = fmt.Sprintf("set({}) ⊃ %v", prettyprint.FormatValue(rhs))
	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		lhsSet, err := ReflectSet(v)
		if err != nil {
			return
		}
		rhsSet, err := ReflectSet(rhs)
		if err != nil {
			return
		}

		diff := rhsSet.Minus(lhsSet)
		r = len(diff) == 0
		if !r {
			ctx = []predicate.ContextValue{
				{Name: "missing values", Value: FormatSetValues(diff), Pre: true},
			}
		}
		return
	}
	return
}

// ---------------------------------------------------------------------------
// Helper functions to manipulate collections as unordered sets

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

// func (lhs Set) Union(rhs Set) (r Set) {
// 	r = Set{}
// 	for v := range lhs {
// 		r[v] = struct{}{}
// 	}
// 	for v := range rhs {
// 		r[v] = struct{}{}
// 	}
// 	return
// }

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

// Helper functions to manipulate collections as unordered sets
// ---------------------------------------------------------------------------
