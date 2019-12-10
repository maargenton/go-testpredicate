package p_test

import (
	"io"
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/p"
	"github.com/maargenton/go-testpredicate/pkg/predicate"
)

// ---------------------------------------------------------------------------
// p.IsNoError()
// ---------------------------------------------------------------------------

func TestIsNoError(t *testing.T) {
	pred := p.IsNoError()

	validatePredicate(t, pred, &predicateExpectation{
		value:        nil,
		result:       predicate.Passed,
		descMatchers: []string{"value is not an error"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:        io.EOF,
		result:       predicate.Failed,
		descMatchers: []string{"value is not an error"},
		errMatchers:  []string{"message:"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:       123,
		result:      predicate.Invalid,
		errMatchers: []string{"/value of type .* is not an error/"},
	})
}
