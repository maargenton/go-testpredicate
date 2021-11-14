package impl_test

import (
	"fmt"
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/utils/predicate/impl"
)

func TestIsError(t *testing.T) {
	var sentinel = fmt.Errorf("sentinel")
	var err = fmt.Errorf("wrapper: %w", sentinel)
	var other = fmt.Errorf("other")

	verifyPredicate(t, pr(impl.IsError(nil)), expectation{value: nil, pass: true})
	verifyPredicate(t, pr(impl.IsError(nil)), expectation{value: other, pass: false})
	verifyPredicate(t, pr(impl.IsError(nil)), expectation{
		value:    123,
		errorMsg: "value of type 'int' is not an error",
	})

	verifyPredicate(t, pr(impl.IsError(sentinel)), expectation{value: err, pass: true})
	verifyPredicate(t, pr(impl.IsError(sentinel)), expectation{value: sentinel, pass: true})
	verifyPredicate(t, pr(impl.IsError(sentinel)), expectation{value: other, pass: false})
	verifyPredicate(t, pr(impl.IsError(sentinel)), expectation{
		value:    123,
		errorMsg: "value of type 'int' is not an error",
	})
}

type MyError struct {
	Code int
}

func (err *MyError) Error() string {
	return fmt.Sprintf("MyError(%v)", err.Code)
}

var _ error = (*MyError)(nil)

func TestAsError(t *testing.T) {
	var sentinel = &MyError{Code: 123}
	var err = fmt.Errorf("wrapper: %w", sentinel)
	var other = fmt.Errorf("other")

	var target *MyError

	verifyTransform(t, tr(impl.AsError(&target)), expectation{
		value:  err,
		result: sentinel,
	})
	verifyTransform(t, tr(impl.AsError(&target)), expectation{
		value:  sentinel,
		result: sentinel,
	})
	verifyTransform(t, tr(impl.AsError(&target)), expectation{
		value:    other,
		result:   nil,
		errorMsg: "value of type '*errors.errorString' is not a '*impl_test.MyError'",
	})
	verifyTransform(t, tr(impl.AsError(target)), expectation{
		value:    err,
		result:   nil,
		errorMsg: "target must be non-nil pointer to an error type, not a '*impl_test.MyError'",
	})
	verifyTransform(t, tr(impl.AsError(&target)), expectation{
		value:    123,
		result:   nil,
		errorMsg: "value of type 'int' is not an error",
	})
}
