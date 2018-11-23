package pred

import (
	"fmt"

	"github.com/marcus999/go-testpredicate"
	"github.com/marcus999/go-testpredicate/utils"
)

// Lt is a shorter alias for IsLessThan
func Lt(rhs interface{}) testpredicate.Predicate {
	return LessThan(rhs)
}

// LessThan returns a predicate that check if a value is strictly less than
// the specified value
func LessThan(rhs interface{}) testpredicate.Predicate {
	return testpredicate.MakeBoolPredicate(
		fmt.Sprintf("value < %v", utils.FormatValue(rhs)),
		func(value interface{}) (bool, error) {
			r, err := utils.CompareOrdered(value, rhs)
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
		func(value interface{}) (bool, error) {
			r, err := utils.CompareOrdered(value, rhs)
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
		func(value interface{}) (bool, error) {
			r, err := utils.CompareOrdered(value, rhs)
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
		func(value interface{}) (bool, error) {
			r, err := utils.CompareOrdered(value, rhs)
			if err != nil {
				return false, err
			}
			return r >= 0, nil
		})
}

// CloseTo return a predicate to check if a numeric value is almost equal to
// the reference value, withing the specified tolerance
func CloseTo(rhs, tolerance float64) testpredicate.Predicate {
	return testpredicate.MakeBoolPredicate(
		fmt.Sprintf("value ≈ %v ± %v", rhs, tolerance),
		func(value interface{}) (bool, error) {
			v, ok := utils.ValueAsFloat(value)
			if !ok {
				return false, fmt.Errorf(
					"value %v of type %T cannot be converted to float",
					utils.FormatValue(value), value)
			}

			return (v >= rhs-tolerance && v <= rhs+tolerance), nil
		})
}
