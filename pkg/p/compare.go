package p

import (
	"fmt"
	"reflect"

	"github.com/maargenton/go-testpredicate/pkg/predicate"
	"github.com/maargenton/go-testpredicate/pkg/prettyprint"
	"github.com/maargenton/go-testpredicate/pkg/value"
)

// IsNil tests if a value is nil
func IsNil() predicate.T {
	return predicate.Make(
		"value is nil",
		func(value interface{}) (predicate.Result, error) {
			if value == nil {
				return predicate.Passed, nil
			}

			var vv = reflect.ValueOf(value)
			switch vv.Kind() {
			case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr,
				reflect.UnsafePointer, reflect.Interface, reflect.Slice:

				if vv.IsNil() {
					return predicate.Passed, nil
				}
				return predicate.Failed, nil

			default:
				return predicate.Invalid, fmt.Errorf(
					"value of type '%v' is never nil",
					vv.Type())
			}
		})
}

// IsNotNil tests if a value is not nil
func IsNotNil() predicate.T {
	return predicate.Make(
		"value is not nil",
		func(value interface{}) (predicate.Result, error) {
			if value == nil {
				return predicate.Failed, nil
			}

			var vv = reflect.ValueOf(value)
			switch vv.Kind() {
			case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr,
				reflect.UnsafePointer, reflect.Interface, reflect.Slice:

				if vv.IsNil() {
					return predicate.Failed, nil
				}
				return predicate.Passed, nil

			default:
				return predicate.Invalid, fmt.Errorf(
					"value of type '%v' is never nil",
					vv.Type())
			}
		})
}

// Eq is a shorter alias for IsEqualTo
func Eq(rhs interface{}) predicate.T {
	return IsEqualTo(rhs)
}

// IsEqualTo tests if a value is comparable and equal to the reference value
func IsEqualTo(rhs interface{}) predicate.T {
	return predicate.MakeBool(
		fmt.Sprintf("value == %v", prettyprint.FormatValue(rhs)),
		func(v interface{}) (bool, error) {
			r, err := value.CompareUnordered(v, rhs)
			if err != nil {
				return false, err
			}
			return r, nil
		})
}

// Ne is a shorter alias for IsNotEqualTo
func Ne(rhs interface{}) predicate.T {
	return IsNotEqualTo(rhs)
}

// IsNotEqualTo tests if a value is comparable but different than the reference value
func IsNotEqualTo(rhs interface{}) predicate.T {
	return predicate.MakeBool(
		fmt.Sprintf("value != %v", prettyprint.FormatValue(rhs)),
		func(v interface{}) (bool, error) {
			r, err := value.CompareUnordered(v, rhs)
			if err != nil {
				return false, err
			}
			return !r, nil
		})
}

// IsTrue tests if a value is true
func IsTrue() predicate.T {
	return predicate.Make(
		"value is true",
		func(value interface{}) (predicate.Result, error) {

			if b, ok := value.(bool); ok {
				if b {
					return predicate.Passed, nil
				}
				return predicate.Failed, nil
			}

			return predicate.Invalid, fmt.Errorf(
				"value of type '%v' is never true",
				reflect.TypeOf(value))
		})
}

// IsFalse tests if a value is false
func IsFalse() predicate.T {
	return predicate.Make(
		"value is false",
		func(value interface{}) (predicate.Result, error) {

			if b, ok := value.(bool); ok {
				if !b {
					return predicate.Passed, nil
				}
				return predicate.Failed, nil
			}

			return predicate.Invalid, fmt.Errorf(
				"value of type '%v' is never false",
				reflect.TypeOf(value))
		})
}
