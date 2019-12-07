package p

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/maargenton/go-testpredicate/pkg/predicate"
	"github.com/maargenton/go-testpredicate/pkg/prettyprint"
)

// Panics returns a predicate that evaluates value as a callable
// function and expects it to panic.
func Panics() predicate.T {
	return predicate.Make(
		"fct() panics",
		func(value interface{}) (predicate.Result, error) {

			var fct, ok = value.(func())
			if !ok {
				return predicate.Invalid, fmt.Errorf(
					"error: value of type '%v' is not callable",
					reflect.TypeOf(value))
			}

			var result = recoverWrapper(fct)
			if result == nil {
				return predicate.Failed, fmt.Errorf(
					"failure: call to fct() did not panic")
			}
			return predicate.Passed, nil
		})
}

// PanicsAndResult returns a predicate that evaluates value as a callable
// function, expect it to panic, and evaluate the panic value against the
// nested predicate
func PanicsAndResult(p predicate.T) predicate.T {
	return predicate.Make(
		"fct() panics with "+strings.Replace(p.String(), "value", "result", -1),
		func(value interface{}) (predicate.Result, error) {

			var fct, ok = value.(func())
			if !ok {
				return predicate.Invalid, fmt.Errorf(
					"error: value of type '%v' is not callable",
					reflect.TypeOf(value))
			}

			var result = recoverWrapper(fct)
			if result == nil {
				return predicate.Failed, fmt.Errorf(
					"failure: call to fct() did not panic")
			}
			r, err := p.Evaluate(result)
			err = predicate.WrapError(err, "panic: %v", prettyprint.FormatValue(result))
			return r, err
		})
}

func recoverWrapper(fct func()) (recoverValue interface{}) {
	defer func() {
		if r := recover(); r != nil {
			recoverValue = r
		}
	}()
	fct()
	return nil
}
