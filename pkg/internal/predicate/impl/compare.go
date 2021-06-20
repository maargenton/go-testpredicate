package impl

import (
	"fmt"
	"reflect"

	"github.com/maargenton/go-testpredicate/pkg/internal/predicate"
	"github.com/maargenton/go-testpredicate/pkg/prettyprint"
	"github.com/maargenton/go-testpredicate/pkg/value"
)

// IsTrue tests if a value is true
func IsTrue() (desc string, f predicate.PredicateFunc) {
	desc = "{} is true"
	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		if b, ok := v.(bool); ok {
			return b, nil, nil
		}
		return false, nil, fmt.Errorf(
			"value of type '%v' is never true",
			reflect.TypeOf(v))
	}
	return
}

// IsFalse tests if a value is false
func IsFalse() (desc string, f predicate.PredicateFunc) {
	desc = "{} is false"
	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		if b, ok := v.(bool); ok {
			return !b, nil, nil
		}
		return false, nil, fmt.Errorf(
			"value of type '%v' is never false",
			reflect.TypeOf(v))
	}
	return
}

// IsNil tests if a value is either a nil literal or a nillable type set to nil
func IsNil() (desc string, f predicate.PredicateFunc) {
	desc = "{} is nil"
	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		if v == nil {
			return true, nil, nil
		}
		var vv = reflect.ValueOf(v)
		switch vv.Kind() {
		case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr,
			reflect.UnsafePointer, reflect.Interface, reflect.Slice:

			return vv.IsNil(), nil, nil
		default:
			return false, nil, fmt.Errorf(
				"value of type '%v' is never nil",
				vv.Type())
		}
	}
	return
}

// IsNil tests if a value is neither a nil literal nor a nillable type set to
// nil
func IsNotNil() (desc string, f predicate.PredicateFunc) {
	desc = "{} is not nil"
	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		if v == nil {
			return false, nil, nil
		}
		var vv = reflect.ValueOf(v)
		switch vv.Kind() {
		case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr,
			reflect.UnsafePointer, reflect.Interface, reflect.Slice:

			return !vv.IsNil(), nil, nil
		default:
			return false, nil, fmt.Errorf(
				"value of type '%v' can never be nil",
				vv.Type())
		}
	}
	return
}

// IsEqualTo tests if a value is equatable and equal to the specified value.
func IsEqualTo(rhs interface{}) (desc string, f predicate.PredicateFunc) {
	desc = fmt.Sprintf("{} == %v", prettyprint.FormatValue(rhs))
	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		eq, err := value.CompareUnordered(v, rhs)
		return eq, nil, err
	}
	return
}

// IsNotEqualTo tests if a value is equatable but different from the specified
// value.
func IsNotEqualTo(rhs interface{}) (desc string, f predicate.PredicateFunc) {
	desc = fmt.Sprintf("{} != %v", prettyprint.FormatValue(rhs))
	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		eq, err := value.CompareUnordered(v, rhs)
		return !eq && err == nil, nil, err
	}
	return
}

// ---------------------------------------------------------------------------
// Aliases

// Eq tests if a value is equatable and equal to the specified value.
func Eq(rhs interface{}) (desc string, f predicate.PredicateFunc) {
	return IsEqualTo(rhs)
}

// Ne tests if a value is equatable but different from the specified
// value.
func Ne(rhs interface{}) (desc string, f predicate.PredicateFunc) {
	return IsNotEqualTo(rhs)
}

// Aliases
// ---------------------------------------------------------------------------
