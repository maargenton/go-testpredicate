package impl

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/maargenton/go-testpredicate/pkg/utils/predicate"
)

// IsError tests an error value to be either nil, a specific error according to
// `errors.Is()`, or an error whose message contains a specified string or
// matches a regexp. `.IsError("")` matches any error whose message contains an
// empty string, which is any non-nil error.
func IsError(expected any) (desc string, f predicate.PredicateFunc) {
	var predicateError error

	if expected == nil {
		desc = "{} is no error"
	} else {
		if _, ok := expected.(error); ok {
			desc = fmt.Sprintf("{} is error '%v'", expected)
		} else if s, ok := expected.(string); ok {
			if len(s) == 0 {
				desc = "{} is an error"
			} else {
				desc = fmt.Sprintf("{} is error containing '%v'", s)
			}
		} else if re, ok := expected.(*regexp.Regexp); ok {
			desc = fmt.Sprintf("{} is error matching /%v/", re)
		} else {
			predicateError = fmt.Errorf(
				"invalid argument of type '%T' for 'IsError()' predicate",
				expected)
		}
	}

	if predicateError != nil {
		f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
			err = predicateError
			return
		}
		return
	}

	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		var errValue, isError = v.(error)
		if !isError && v != nil {
			err = fmt.Errorf("value of type '%T' is not an error", v)
			return
		}
		if isError {
			ctx = []predicate.ContextValue{
				{Name: "message", Value: errValue.Error()},
			}
		}

		if expected == nil {
			r = errValue == nil
		} else if expectedErr, ok := expected.(error); ok {
			r = errors.Is(errValue, expectedErr)
		} else if expectedString, ok := expected.(string); ok && errValue != nil {
			r = strings.Contains(errValue.Error(), expectedString)
		} else if expectedRegexp, ok := expected.(*regexp.Regexp); ok && errValue != nil {
			r = expectedRegexp.MatchString(errValue.Error())
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
				ctx = []predicate.ContextValue{
					{Name: "message", Value: errValue.Error()},
				}
			} else {
				r = tv.Elem().Interface()
				ctx = []predicate.ContextValue{
					{Name: "target", Value: r},
					{Name: "message", Value: errValue.Error()},
				}
			}
		} else {
			err = fmt.Errorf("value of type '%T' is not an error", v)
		}
		return
	}
	return
}
