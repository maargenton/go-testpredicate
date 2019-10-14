package pred

import (
	"fmt"
	"reflect"

	"github.com/maargenton/go-testpredicate"
	"github.com/maargenton/go-testpredicate/utils"
)

// IsNil tests if a value is nil
func IsNil() testpredicate.Predicate {
	return testpredicate.MakePredicate(
		"value is nil",
		func(value interface{}) (testpredicate.PredicateResult, error) {
			if value == nil {
				return testpredicate.PredicatePassed, nil
			}

			var vv = reflect.ValueOf(value)
			switch vv.Kind() {
			case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr,
				reflect.UnsafePointer, reflect.Interface, reflect.Slice:

				if vv.IsNil() {
					return testpredicate.PredicatePassed, nil
				}
				return testpredicate.PredicateFailed, nil

			default:
				return testpredicate.PredicateInvalid, fmt.Errorf(
					"value of type '%v' is never nil",
					vv.Type())
			}
		})
}

// IsNotNil tests if a value is not nil
func IsNotNil() testpredicate.Predicate {
	return testpredicate.MakePredicate(
		"value is not nil",
		func(value interface{}) (testpredicate.PredicateResult, error) {
			if value == nil {
				return testpredicate.PredicateFailed, nil
			}

			var vv = reflect.ValueOf(value)
			switch vv.Kind() {
			case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr,
				reflect.UnsafePointer, reflect.Interface, reflect.Slice:

				if vv.IsNil() {
					return testpredicate.PredicateFailed, nil
				}
				return testpredicate.PredicatePassed, nil

			default:
				return testpredicate.PredicateInvalid, fmt.Errorf(
					"value of type '%v' is never nil",
					vv.Type())
			}
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

// IsTrue tests if a value is true
func IsTrue() testpredicate.Predicate {
	return testpredicate.MakePredicate(
		"value is true",
		func(value interface{}) (testpredicate.PredicateResult, error) {

			if b, ok := value.(bool); ok {
				if b == true {
					return testpredicate.PredicatePassed, nil
				}
				return testpredicate.PredicateFailed, nil
			}

			return testpredicate.PredicateInvalid, fmt.Errorf(
				"value of type '%v' is never true",
				reflect.TypeOf(value))
		})
}

// IsFalse tests if a value is false
func IsFalse() testpredicate.Predicate {
	return testpredicate.MakePredicate(
		"value is false",
		func(value interface{}) (testpredicate.PredicateResult, error) {

			if b, ok := value.(bool); ok {
				if b == false {
					return testpredicate.PredicatePassed, nil
				}
				return testpredicate.PredicateFailed, nil
			}

			return testpredicate.PredicateInvalid, fmt.Errorf(
				"value of type '%v' is never false",
				reflect.TypeOf(value))
		})
}
