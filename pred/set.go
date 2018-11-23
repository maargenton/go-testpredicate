package pred

import (
	"fmt"
	"reflect"

	"github.com/marcus999/go-testpredicate"
	"github.com/marcus999/go-testpredicate/utils"
)

// IsSubsetOf returns a predicate that checks
func IsSubsetOf(rhs interface{}) testpredicate.Predicate {
	return testpredicate.MakePredicate(
		fmt.Sprintf("value is subset of %v", utils.FormatValue(rhs)),
		func(value interface{}) (testpredicate.PredicateResult, error) {

			lhsSet, err := reflectSet(value)
			if err != nil {
				return testpredicate.PredicateInvalid, err
			}

			rhsSet, err := reflectSet(rhs)
			if err != nil {
				return testpredicate.PredicateInvalid, err
			}

			diff := diffSet(lhsSet, rhsSet)
			if len(diff) == 0 {
				return testpredicate.PredicatePassed, nil
			}

			return testpredicate.PredicateFailed, fmt.Errorf(
				"extra values: %v", formatSetValues(diff))
		})
}

// IsSupersetOf returns a predicate that checks
func IsSupersetOf(rhs interface{}) testpredicate.Predicate {
	return testpredicate.MakePredicate(
		fmt.Sprintf("value is superset of %v", utils.FormatValue(rhs)),
		func(value interface{}) (testpredicate.PredicateResult, error) {

			lhsSet, err := reflectSet(value)
			if err != nil {
				return testpredicate.PredicateInvalid, err
			}

			rhsSet, err := reflectSet(rhs)
			if err != nil {
				return testpredicate.PredicateInvalid, err
			}

			diff := diffSet(rhsSet, lhsSet)
			if len(diff) == 0 {
				return testpredicate.PredicatePassed, nil
			}

			return testpredicate.PredicateFailed, fmt.Errorf(
				"missing values: %v", formatSetValues(diff))
		})
}

// IsDisjointSetFrom returns a predicate that checks
func IsDisjointSetFrom(rhs interface{}) testpredicate.Predicate {
	return testpredicate.MakePredicate(
		fmt.Sprintf("value is disjoint from %v", utils.FormatValue(rhs)),
		func(value interface{}) (testpredicate.PredicateResult, error) {

			lhsSet, err := reflectSet(value)
			if err != nil {
				return testpredicate.PredicateInvalid, err
			}

			rhsSet, err := reflectSet(rhs)
			if err != nil {
				return testpredicate.PredicateInvalid, err
			}

			inter := intersectSet(lhsSet, rhsSet)
			if len(inter) == 0 {
				return testpredicate.PredicatePassed, nil
			}

			return testpredicate.PredicateFailed, fmt.Errorf(
				"common values: %v", formatSetValues(inter))
		})
}

// IsEqualSet returns a predicate that checks
func IsEqualSet(rhs interface{}) testpredicate.Predicate {
	return testpredicate.MakePredicate(
		fmt.Sprintf("value is equal set as %v", utils.FormatValue(rhs)),
		func(value interface{}) (testpredicate.PredicateResult, error) {

			lhsSet, err := reflectSet(value)
			if err != nil {
				return testpredicate.PredicateInvalid, err
			}

			rhsSet, err := reflectSet(rhs)
			if err != nil {
				return testpredicate.PredicateInvalid, err
			}

			extra := diffSet(lhsSet, rhsSet)
			missing := diffSet(rhsSet, lhsSet)
			if len(extra) == 0 && len(missing) == 0 {
				return testpredicate.PredicatePassed, nil
			}

			return testpredicate.PredicateFailed, fmt.Errorf(
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
			"value %v of type %T is not a indexable collection", utils.FormatValue(value), value)
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
			values += utils.FormatValue(v)
		} else {
			values += ", ..."
			break
		}
	}

	return values
}
