// +build go1.13

package p_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/p"
	"github.com/maargenton/go-testpredicate/pkg/predicate"
)

// ---------------------------------------------------------------------------
// p.IsError()
// ---------------------------------------------------------------------------

func TestIsError(t *testing.T) {
	pred := p.IsError(io.EOF)

	validatePredicate(t, pred, &predicateExpectation{
		value:        nil,
		result:       predicate.Failed,
		descMatchers: []string{"value is an error matching"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:        io.ErrClosedPipe,
		result:       predicate.Failed,
		descMatchers: []string{"value is an error matching"},
		errMatchers:  []string{"detailed error:"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:        io.EOF,
		result:       predicate.Passed,
		descMatchers: []string{"value is an error matching"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:       123,
		result:      predicate.Invalid,
		errMatchers: []string{"/value of type .* is not an error/"},
	})
}

func TestIsErrorWithWrappedError(t *testing.T) {
	pred := p.IsError(io.EOF)
	err := fmt.Errorf("custom error, base error: %w", io.EOF)

	validatePredicate(t, pred, &predicateExpectation{
		value:  err,
		result: predicate.Passed,
	})
}

func TestIsErrorNil(t *testing.T) {
	pred := p.IsError(nil)

	validatePredicate(t, pred, &predicateExpectation{
		value:  nil,
		result: predicate.Passed,
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:       io.EOF,
		result:      predicate.Failed,
		errMatchers: []string{"detailed error:"},
	})
}
