package impl

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/maargenton/go-testpredicate/pkg/utils/predicate"
)

// IsError tests if a value is an error matching or wrapping the expected error
// (according to go 1.13 error.Is()).
func IsError(expected error) (desc string, f predicate.PredicateFunc) {
	if expected != nil {
		desc = fmt.Sprintf("{} is error '%v'", expected)
	} else {
		desc = "{} is no error"
	}
	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		if v == nil {
			r = expected == nil
		} else if errValue, ok := v.(error); ok {
			r = errors.Is(errValue, expected)
			ctx = []predicate.ContextValue{
				{Name: "error", Value: errValue.Error(), Pre: true},
			}

		} else {
			err = fmt.Errorf("value of type '%T' is not an error", v)
		}
		return
	}
	return
}

// AsError tests if a value is an error matching or wrapping the expected error
// (according to go 1.13 error.As()) and returns the unwrapped error for further
// evaluation.
func AsError(target interface{}) (desc string, f predicate.TransformFunc) {
	var v = reflect.ValueOf(target)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	desc = fmt.Sprintf("{}.As(%v)", v.Type())

	f = func(v interface{}) (r interface{}, ctx []predicate.ContextValue, err error) {
		var tv = reflect.ValueOf(target)
		var errorType = reflect.TypeOf((*error)(nil)).Elem()
		if tv.Kind() != reflect.Ptr || tv.IsNil() || !tv.Elem().Type().Implements(errorType) {
			err = fmt.Errorf("target must be non-nil pointer to an error type, not a '%T'", target)
			return
		}

		if errValue, ok := v.(error); ok {
			if !errors.As(errValue, target) {
				err = fmt.Errorf("value of type '%T' is not a '%v'", v, tv.Elem().Type())
			} else {
				r = tv.Elem().Interface()
				ctx = []predicate.ContextValue{
					{Name: "target error", Value: r},
				}
			}
		} else {
			err = fmt.Errorf("value of type '%T' is not an error", v)
		}
		return
	}
	return
}
