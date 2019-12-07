package pred

import (
	"fmt"

	"github.com/maargenton/go-testpredicate"
	"github.com/maargenton/go-testpredicate/pkg/value"
)

// Lt is a shorter alias for IsLessThan
func Lt(rhs interface{}) testpredicate.Predicate {
	return LessThan(rhs)
}

// LessThan returns a predicate that check if a value is strictly less than
// the specified value
func LessThan(rhs interface{}) testpredicate.Predicate {
	return testpredicate.MakeBoolPredicate(
		fmt.Sprintf("value < %v", rhs),
		func(v interface{}) (bool, error) {
			r, err := value.CompareOrdered(v, rhs)
			if err != nil {
				return false, err
			}
			return r < 0, nil
		})
}

// Le is a shorter alias for IsLessOrEqualTo
func Le(rhs interface{}) testpredicate.Predicate {
	return LessOrEqualTo(rhs)
}

// LessOrEqualTo returns a predicate that check if a value is less or equal to
// the specified value
func LessOrEqualTo(rhs interface{}) testpredicate.Predicate {
	return testpredicate.MakeBoolPredicate(
		fmt.Sprintf("value <= %v", rhs),
		func(v interface{}) (bool, error) {
			r, err := value.CompareOrdered(v, rhs)
			if err != nil {
				return false, err
			}
			return r <= 0, nil
		})
}

// Gt is a shorter alias for IsGreaterThan
func Gt(rhs interface{}) testpredicate.Predicate {
	return GreaterThan(rhs)
}

// GreaterThan returns a predicate that check if a value is strictly greater
// than the specified value
func GreaterThan(rhs interface{}) testpredicate.Predicate {
	return testpredicate.MakeBoolPredicate(
		fmt.Sprintf("value > %v", rhs),
		func(v interface{}) (bool, error) {
			r, err := value.CompareOrdered(v, rhs)
			if err != nil {
				return false, err
			}
			return r > 0, nil
		})
}

// Ge is a shorter alias for IsGreaterOrEqualTo
func Ge(rhs interface{}) testpredicate.Predicate {
	return GreaterOrEqualTo(rhs)
}

// GreaterOrEqualTo returns a predicate that check if a value is greater or equal
// to the specified value
func GreaterOrEqualTo(rhs interface{}) testpredicate.Predicate {
	return testpredicate.MakeBoolPredicate(
		fmt.Sprintf("value >= %v", rhs),
		func(v interface{}) (bool, error) {
			r, err := value.CompareOrdered(v, rhs)
			if err != nil {
				return false, err
			}
			return r >= 0, nil
		})
}

// CloseTo return a predicate to check if a value is almost equal to
// the reference value, within the specified tolerance. Value and reference
// value can numeric, or slice or array of numeric values, of equal size
func CloseTo(rhs interface{}, tolerance float64) testpredicate.Predicate {
	return testpredicate.MakeBoolPredicate(
		fmt.Sprintf("value ≈ %v ± %v", rhs, tolerance),

		func(v interface{}) (bool, error) {
			d, err := value.MaxAbsoluteDifference(v, rhs)
			if err != nil {
				return false, err
			}
			if d > tolerance {
				return false, nil
			}

			return true, nil
		})
}
