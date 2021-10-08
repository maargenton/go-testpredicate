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
