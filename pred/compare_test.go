package pred_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate"
	"github.com/maargenton/go-testpredicate/pred"
)

// ---------------------------------------------------------------------------
// pred.IsNil()
// ---------------------------------------------------------------------------

func TestIsNil(t *testing.T) {
	p := pred.IsNil()

	validatePredicate(t, p, &predicateExpectation{
		value:        nil,
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value is nil"},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:  123,
		result: testpredicate.PredicateFailed,
	})
	validatePredicate(t, p, &predicateExpectation{
		value:  "abc",
		result: testpredicate.PredicateFailed,
	})
}

// ---------------------------------------------------------------------------
// pred.IsNotNil()
// ---------------------------------------------------------------------------

func TestIsNotNil(t *testing.T) {
	p := pred.IsNotNil()

	validatePredicate(t, p, &predicateExpectation{
		value:        123,
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value is not nil"},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:  nil,
		result: testpredicate.PredicateFailed,
	})
}

// ---------------------------------------------------------------------------
// pred.IsEqualTo()
// ---------------------------------------------------------------------------

func TestIsEqualTo(t *testing.T) {
	p := pred.Eq("123")

	validatePredicate(t, p, &predicateExpectation{
		value:        "123",
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{`value == "123"`},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:  "124",
		result: testpredicate.PredicateFailed,
	})
	validatePredicate(t, p, &predicateExpectation{
		value:  124,
		result: testpredicate.PredicateInvalid,
		errMatchers: []string{
			`/values of type .* and .* are never equal/`,
		},
	})
}

// ---------------------------------------------------------------------------
// pred.IsNotEqualTo()
// ---------------------------------------------------------------------------

func TestNotIsEqualTo(t *testing.T) {
	p := pred.Ne("123")

	validatePredicate(t, p, &predicateExpectation{
		value:        "124",
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{`value != "123"`},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:  "123",
		result: testpredicate.PredicateFailed,
	})
	validatePredicate(t, p, &predicateExpectation{
		value:  124,
		result: testpredicate.PredicateInvalid,
		errMatchers: []string{
			`/values of type .* and .* are never equal/`,
		},
	})
}
