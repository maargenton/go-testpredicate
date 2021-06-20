package impl_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/internal/predicate/impl"
)

func TestPanic(t *testing.T) {
	verifyPredicate(t, pr(impl.Panics()), expectation{
		value: func() { panic(123) },
		pass:  true,
	})
	verifyPredicate(t, pr(impl.Panics()), expectation{
		value: func() {},
		pass:  false,
	})
	verifyPredicate(t, pr(impl.Panics()), expectation{
		value:    func() { panic(nil) },
		pass:     false,
		errorMsg: "value() panicked with a nil value",
	})
	verifyPredicate(t, pr(impl.Panics()), expectation{
		value:    123,
		pass:     false,
		errorMsg: "value of type 'int' is not callable",
	})
}

func TestPanicsAndRecoveredValue(t *testing.T) {
	verifyTransform(t, tr(impl.PanicsAndRecoveredValue()), expectation{
		value:  func() { panic(123) },
		result: 123,
	})
	verifyTransform(t, tr(impl.PanicsAndRecoveredValue()), expectation{
		value:    func() {},
		result:   nil,
		errorMsg: "value() did not panic",
	})
	verifyTransform(t, tr(impl.PanicsAndRecoveredValue()), expectation{
		value:    123,
		result:   nil,
		errorMsg: "value of type 'int' is not callable",
	})
}
