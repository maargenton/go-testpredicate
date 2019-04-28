package pred

import (
	"fmt"

	"github.com/maargenton/go-testpredicate"
	"github.com/maargenton/go-testpredicate/utils"
)

// IsNil tests if a value is nil
func IsNil() testpredicate.Predicate {
	return testpredicate.MakeBoolPredicate(
		"value is nil",
		func(value interface{}) (bool, error) {
			return value == nil, nil
		})
}

// IsNotNil tests if a value is not nil
func IsNotNil() testpredicate.Predicate {
	return testpredicate.MakeBoolPredicate(
		"value is not nil",
		func(value interface{}) (bool, error) {
			return value != nil, nil
		})
}

// Eq is a shorter alias for IsEqualTo
func Eq(rhs interface{}) testpredicate.Predicate {
	return IsEqualTo(rhs)
}

// IsEqualTo tests if a value is comparable and equal to the reference value
func IsEqualTo(rhs interface{}) testpredicate.Predicate {
	return testpredicate.MakeBoolPredicate(
		fmt.Sprintf("value == %v", utils.FormatValue(rhs)),
		func(value interface{}) (bool, error) {
			r, err := utils.CompareUnordered(value, rhs)
			if err != nil {
				return false, err
			}
			return r, nil
		})
}

// Ne is a shorter alias for IsNotEqualTo
func Ne(rhs interface{}) testpredicate.Predicate {
	return IsNotEqualTo(rhs)
}

// IsNotEqualTo tests if a value is comparable but different than the reference value
func IsNotEqualTo(rhs interface{}) testpredicate.Predicate {
	return testpredicate.MakeBoolPredicate(
		fmt.Sprintf("value != %v", utils.FormatValue(rhs)),
		func(value interface{}) (bool, error) {
			r, err := utils.CompareUnordered(value, rhs)
			if err != nil {
				return false, err
			}
			return !r, nil
		})
}
