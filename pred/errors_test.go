package pred_test

import (
	"io"
	"testing"

	"github.com/maargenton/go-testpredicate"
	"github.com/maargenton/go-testpredicate/pred"
	"golang.org/x/xerrors"
)

// ---------------------------------------------------------------------------
// pred.IsNoError()
// ---------------------------------------------------------------------------

func TestIsNoError(t *testing.T) {
	p := pred.IsNoError()

	validatePredicate(t, p, &predicateExpectation{
		value:        nil,
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value is not an error"},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:        io.EOF,
		result:       testpredicate.PredicateFailed,
		descMatchers: []string{"value is not an error"},
		errMatchers:  []string{"detailed error:"},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:       123,
		result:      testpredicate.PredicateInvalid,
		errMatchers: []string{"/value of type .* is not an error/"},
	})
}

// ---------------------------------------------------------------------------
// pred.IsError()
// ---------------------------------------------------------------------------

func TestIsError(t *testing.T) {
	p := pred.IsError(io.EOF)

	validatePredicate(t, p, &predicateExpectation{
		value:        nil,
		result:       testpredicate.PredicateFailed,
		descMatchers: []string{"value is an error matching"},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:        io.ErrClosedPipe,
		result:       testpredicate.PredicateFailed,
		descMatchers: []string{"value is an error matching"},
		errMatchers:  []string{"detailed error:"},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:        io.EOF,
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value is an error matching"},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:       123,
		result:      testpredicate.PredicateInvalid,
		errMatchers: []string{"/value of type .* is not an error/"},
	})
}

func TestIsErrorWithWrappedError(t *testing.T) {
	p := pred.IsError(io.EOF)
	err := xerrors.Errorf("custom error, base error: %w", io.EOF)

	validatePredicate(t, p, &predicateExpectation{
		value:  err,
		result: testpredicate.PredicatePassed,
	})
}
