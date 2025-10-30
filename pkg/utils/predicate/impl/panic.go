package impl

import (
	"fmt"
	"reflect"

	"github.com/maargenton/go-testpredicate/pkg/utils/predicate"
)

// Panics verifies that the value under test is a callable function that panics.
// Since version > 1.5.0, `panic(nil)` is no longer a special case, reflecting
// the change of behavior in Go 1.21.
func Panics() (desc string, f predicate.PredicateFunc) {
	desc = "{}() panics"
	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		var fct, ok = v.(func())
		if !ok {
			return false, nil, fmt.Errorf(
				"value of type '%v' is not callable",
				reflect.TypeOf(v))
		}
		var panicked, _ = recoverWrapper(fct)
		return panicked, nil, nil
	}
	return
}

// PanicsAndRecoveredValue verifies that the value under test is a callable
// function that panics, and captures the recovered value for further
// evaluation.
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
