package impl

import (
	"fmt"

	"github.com/maargenton/go-testpredicate/pkg/internal/predicate"
	"github.com/maargenton/go-testpredicate/pkg/prettyprint"
	"github.com/maargenton/go-testpredicate/pkg/value"
)

// IsEqualSet tests if two containers contain the same set of values,
// independently of order.
func IsEqualSet(rhs interface{}) (desc string, f predicate.PredicateFunc) {
	desc = fmt.Sprintf("set({}) == %v", prettyprint.FormatValue(rhs))
	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		lhsSet, err := value.ReflectSet(v)
		if err != nil {
			return
		}
		rhsSet, err := value.ReflectSet(rhs)
		if err != nil {
			return
		}

		extra := lhsSet.Minus(rhsSet)
		missing := rhsSet.Minus(lhsSet)
		r = len(extra) == 0 && len(missing) == 0
		if !r {
			ctx = []predicate.ContextValue{
				{Name: "extra values", Value: value.FormatSetValues(extra), Pre: true},
				{Name: "missing values", Value: value.FormatSetValues(missing), Pre: true},
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
		lhsSet, err := value.ReflectSet(v)
		if err != nil {
			return
		}
		rhsSet, err := value.ReflectSet(rhs)
		if err != nil {
			return
		}

		intersection := lhsSet.Intersect(rhsSet)
		r = len(intersection) == 0
		if !r {
			ctx = []predicate.ContextValue{
				{Name: "common values", Value: value.FormatSetValues(intersection), Pre: true},
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
		lhsSet, err := value.ReflectSet(v)
		if err != nil {
			return
		}
		rhsSet, err := value.ReflectSet(rhs)
		if err != nil {
			return
		}

		diff := lhsSet.Minus(rhsSet)
		r = len(diff) == 0
		if !r {
			ctx = []predicate.ContextValue{
				{Name: "extra values", Value: value.FormatSetValues(diff), Pre: true},
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
		lhsSet, err := value.ReflectSet(v)
		if err != nil {
			return
		}
		rhsSet, err := value.ReflectSet(rhs)
		if err != nil {
			return
		}

		diff := rhsSet.Minus(lhsSet)
		r = len(diff) == 0
		if !r {
			ctx = []predicate.ContextValue{
				{Name: "missing values", Value: value.FormatSetValues(diff), Pre: true},
			}
		}
		return
	}
	return
}
