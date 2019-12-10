package p

import (
	"fmt"
	"reflect"

	"github.com/maargenton/go-testpredicate/pkg/predicate"
	"github.com/maargenton/go-testpredicate/pkg/prettyprint"
)

// IsSubsetOf returns a predicate that checks
func IsSubsetOf(rhs interface{}) predicate.T {
	return predicate.Make(
		fmt.Sprintf("value is subset of %v", prettyprint.FormatValue(rhs)),
		func(value interface{}) (predicate.Result, error) {

			lhsSet, err := reflectSet(value)
			if err != nil {
				return predicate.Invalid, err
			}

			rhsSet, err := reflectSet(rhs)
			if err != nil {
				return predicate.Invalid, err
			}

			diff := diffSet(lhsSet, rhsSet)
			if len(diff) == 0 {
				return predicate.Passed, nil
			}

			return predicate.Failed, fmt.Errorf(
				"extra values: %v", formatSetValues(diff))
		})
}

// IsSupersetOf returns a predicate that checks
func IsSupersetOf(rhs interface{}) predicate.T {
	return predicate.Make(
		fmt.Sprintf("value is superset of %v", prettyprint.FormatValue(rhs)),
		func(value interface{}) (predicate.Result, error) {

			lhsSet, err := reflectSet(value)
			if err != nil {
				return predicate.Invalid, err
			}

			rhsSet, err := reflectSet(rhs)
			if err != nil {
				return predicate.Invalid, err
			}

			diff := diffSet(rhsSet, lhsSet)
			if len(diff) == 0 {
				return predicate.Passed, nil
			}

			return predicate.Failed, fmt.Errorf(
				"missing values: %v", formatSetValues(diff))
		})
}

// IsDisjointSetFrom returns a predicate that checks
func IsDisjointSetFrom(rhs interface{}) predicate.T {
	return predicate.Make(
		fmt.Sprintf("value is disjoint from %v", prettyprint.FormatValue(rhs)),
		func(value interface{}) (predicate.Result, error) {

			lhsSet, err := reflectSet(value)
			if err != nil {
				return predicate.Invalid, err
			}

			rhsSet, err := reflectSet(rhs)
			if err != nil {
				return predicate.Invalid, err
			}

			inter := intersectSet(lhsSet, rhsSet)
			if len(inter) == 0 {
				return predicate.Passed, nil
			}

			return predicate.Failed, fmt.Errorf(
				"common values: %v", formatSetValues(inter))
		})
}

// IsEqualSet returns a predicate that checks
func IsEqualSet(rhs interface{}) predicate.T {
	return predicate.Make(
		fmt.Sprintf("value is equal set as %v", prettyprint.FormatValue(rhs)),
		func(value interface{}) (predicate.Result, error) {

			lhsSet, err := reflectSet(value)
			if err != nil {
				return predicate.Invalid, err
			}

			rhsSet, err := reflectSet(rhs)
			if err != nil {
				return predicate.Invalid, err
			}

			extra := diffSet(lhsSet, rhsSet)
			missing := diffSet(rhsSet, lhsSet)
			if len(extra) == 0 && len(missing) == 0 {
				return predicate.Passed, nil
			}

			return predicate.Failed, fmt.Errorf(
				"extra values: %v\nmissing values: %v",
				formatSetValues(extra), formatSetValues(missing))
		})
}

// ---------------------------------------------------------------------------
// ---------------------------------------------------------------------------
// local Set type and helpers
// ---------------------------------------------------------------------------

type set map[interface{}]struct{}

func isIndexable(k reflect.Kind) bool {
	if k == reflect.Slice || k == reflect.Array || k == reflect.String {
		return true
	}
	return false
}

func reflectSet(value interface{}) (set, error) {
	v := reflect.ValueOf(value)

	if !isIndexable(v.Kind()) {
		return set{}, fmt.Errorf(
			"value of type '%T' is not a indexable collection", value)
	}

	s := make(set, v.Len())

	for i, n := 0, v.Len(); i < n; i++ {
		s[v.Index(i).Interface()] = struct{}{}
	}

	return s, nil
}

func intMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func diffSet(lhs, rhs set) set {
	n := intMax(len(lhs), len(rhs))
	s := make(set, n)

	for v := range lhs {
		if _, ok := rhs[v]; !ok {
			s[v] = struct{}{}
		}
	}

	return s
}

func intersectSet(lhs, rhs set) set {
	n := intMax(len(lhs), len(rhs))
	s := make(set, n)

	for v := range lhs {
		if _, ok := rhs[v]; ok {
			s[v] = struct{}{}
		}
	}
	return s
}

func formatSetValues(s set) string {
	values := ""
	for v := range s {
		if len(values) < 50 {
			if len(values) > 0 {
				values += ", "
			}
			values += prettyprint.FormatValue(v)
		} else {
			values += ", ..."
			break
		}
	}

	return values
}
