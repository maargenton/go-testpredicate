package pred_test

import (
	"io"
	"testing"

	"github.com/maargenton/go-testpredicate"
	"github.com/maargenton/go-testpredicate/pred"
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
