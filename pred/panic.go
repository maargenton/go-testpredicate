package pred

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/maargenton/go-testpredicate"
	"github.com/maargenton/go-testpredicate/pkg/prettyprint"
	"github.com/maargenton/go-testpredicate/utils"
)

// Panics returns a predicate that evaluates value as a callable
// function and expects it to panic.
func Panics() testpredicate.Predicate {
	return testpredicate.MakePredicate(
		"fct() panics",
		func(value interface{}) (testpredicate.PredicateResult, error) {

			var fct, ok = value.(func())
			if !ok {
				return testpredicate.PredicateInvalid, fmt.Errorf(
					"error: value of type '%v' is not callable",
					reflect.TypeOf(value))
			}

			var result = recoverWrapper(fct)
			if result == nil {
				return testpredicate.PredicateFailed, fmt.Errorf(
					"failure: call to fct() did not panic")
			}
			return testpredicate.PredicatePassed, nil
		})
}

// PanicsAndResult returns a predicate that evaluates value as a callable
// function, expect it to panic, and evaluate the panic value against the
// nested predicate
func PanicsAndResult(p testpredicate.Predicate) testpredicate.Predicate {
	return testpredicate.MakePredicate(
		"fct() panics with "+strings.Replace(p.String(), "value", "result", -1),
		func(value interface{}) (testpredicate.PredicateResult, error) {

			var fct, ok = value.(func())
			if !ok {
				return testpredicate.PredicateInvalid, fmt.Errorf(
					"error: value of type '%v' is not callable",
					reflect.TypeOf(value))
			}

			var result = recoverWrapper(fct)
			if result == nil {
				return testpredicate.PredicateFailed, fmt.Errorf(
					"failure: call to fct() did not panic")
			}
			r, err := p.Evaluate(result)
			err = utils.WrapError(err, "panic: %v", prettyprint.FormatValue(result))
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
