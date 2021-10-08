package impl

import (
	"fmt"
	"reflect"

	"github.com/maargenton/go-testpredicate/pkg/utils/predicate"
)

// Panics verifies that the value under test is a callable function that panics.
// Special case using panic(nil) is considered an error because common recover()
// code will not catch it.
func Panics() (desc string, f predicate.PredicateFunc) {
	desc = "{}() panics"
	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		var fct, ok = v.(func())
		if !ok {
			return false, nil, fmt.Errorf(
				"value of type '%v' is not callable",
				reflect.TypeOf(v))
		}
		panicked, recoveredValue := recoverWrapper(fct)
		if panicked && recoveredValue == nil {
			return false, nil, fmt.Errorf(
				"value() panicked with a nil value")
		}
		return panicked && recoveredValue != nil, nil, nil
	}
	return
}

// PanicsAndRecoveredValue verifies that the value under test is a callable
// function that panics, and captures the recovered value for further evalation.
func PanicsAndRecoveredValue() (desc string, f predicate.TransformFunc) {
	desc = "recover({}())"
	f = func(v interface{}) (r interface{}, ctx []predicate.ContextValue, err error) {
		var fct, ok = v.(func())
		if !ok {
			return nil, nil, fmt.Errorf(
				"value of type '%v' is not callable",
				reflect.TypeOf(v))
		}
		panicked, recoverValue := recoverWrapper(fct)
		if !panicked {
			return nil, nil, fmt.Errorf(
				"value() did not panic")
		}
		return recoverValue, []predicate.ContextValue{
			{Name: "recovered", Value: recoverValue},
		}, nil
	}
	return
}

// ---------------------------------------------------------------------------
// Panic related predicated helpers

func recoverWrapper(fct func()) (panicked bool, recoverValue interface{}) {
	defer func() {
		r := recover()
		recoverValue = r
	}()
	panicked = true
	panicked = panicWrapper(fct)
	return
}

func panicWrapper(fct func()) (panicked bool) {
	panicked = true
	fct()
	panicked = false
	return
}

// Panic related predicated helpers
// ---------------------------------------------------------------------------
