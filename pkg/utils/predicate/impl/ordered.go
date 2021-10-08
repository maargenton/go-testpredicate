package impl

import (
	"fmt"

	"github.com/maargenton/go-testpredicate/pkg/utils/predicate"
	"github.com/maargenton/go-testpredicate/pkg/utils/prettyprint"
	"github.com/maargenton/go-testpredicate/pkg/utils/value"
)

// ---------------------------------------------------------------------------
// Comparison predicates

// IsLessThan tests if a value is strictly less than a reference value
func IsLessThan(rhs interface{}) (desc string, f predicate.PredicateFunc) {
	desc = fmt.Sprintf("{} < %v", prettyprint.FormatValue(rhs))
	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		order, err := value.CompareOrdered(v, rhs)
		return order < 0 && err == nil, nil, err
	}
	return
}

// IsLessOrEqualTo tests if a value is less than or equal to a reference value
func IsLessOrEqualTo(rhs interface{}) (desc string, f predicate.PredicateFunc) {
	desc = fmt.Sprintf("{} <= %v", prettyprint.FormatValue(rhs))
	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		order, err := value.CompareOrdered(v, rhs)
		return order <= 0 && err == nil, nil, err
	}
	return
}

// IsGreaterThan tests if a value is strictly greater than a reference value
func IsGreaterThan(rhs interface{}) (desc string, f predicate.PredicateFunc) {
	desc = fmt.Sprintf("{} > %v", prettyprint.FormatValue(rhs))
	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		order, err := value.CompareOrdered(v, rhs)
		return order > 0 && err == nil, nil, err
	}
	return
}

// IsGreaterOrEqualTo tests if a value is greater than or equal to a reference value
func IsGreaterOrEqualTo(rhs interface{}) (desc string, f predicate.PredicateFunc) {
	desc = fmt.Sprintf("{} >= %v", prettyprint.FormatValue(rhs))
	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		order, err := value.CompareOrdered(v, rhs)
		return order >= 0 && err == nil, nil, err
	}
	return
}

// IsCloseTo tests if a value is within tolerance of a reference value
func IsCloseTo(rhs interface{}, tolerance float64) (desc string, f predicate.PredicateFunc) {
	desc = fmt.Sprintf("{} ≈ %v ± %v", prettyprint.FormatValue(rhs), tolerance)
	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		delta, err := value.MaxAbsoluteDifference(v, rhs)
		return delta <= tolerance && err == nil, []predicate.ContextValue{
			{Name: "difference", Value: delta},
		}, err
	}
	return
}

// Comparison predicates
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// Aliases

// Lt tests if a value is strictly less than a reference value
func Lt(rhs interface{}) (desc string, f predicate.PredicateFunc) {
	return IsLessThan(rhs)
}

// Le tests if a value is less than or equal to a reference value
func Le(rhs interface{}) (desc string, f predicate.PredicateFunc) {
	return IsLessOrEqualTo(rhs)
}

// Gt tests if a value is strictly greater than a reference value
func Gt(rhs interface{}) (desc string, f predicate.PredicateFunc) {
	return IsGreaterThan(rhs)
}

// Ge tests if a value is greater than or equal to a reference value
func Ge(rhs interface{}) (desc string, f predicate.PredicateFunc) {
	return IsGreaterOrEqualTo(rhs)
}

// Aliases
// ---------------------------------------------------------------------------
