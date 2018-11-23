package pred_test

import (
	"testing"

	"github.com/marcus999/go-testpredicate"
	"github.com/marcus999/go-testpredicate/pred"
)

// ---------------------------------------------------------------------------
// pred.IsNil()
// ---------------------------------------------------------------------------

func TestIsNil(t *testing.T) {
	p := pred.IsNil()

	validateredicate(t, p, &predicateExpectation{
		value:        nil,
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value is nil"},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  123,
		result: testpredicate.PredicateFailed,
	})
	validateredicate(t, p, &predicateExpectation{
		value:  "abc",
		result: testpredicate.PredicateFailed,
	})
}

// ---------------------------------------------------------------------------
// pred.IsNotNil()
// ---------------------------------------------------------------------------

func TestIsNotNil(t *testing.T) {
	p := pred.IsNotNil()

	validateredicate(t, p, &predicateExpectation{
		value:        123,
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value is not nil"},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  nil,
		result: testpredicate.PredicateFailed,
	})
}

// ---------------------------------------------------------------------------
// pred.IsEqualTo()
// ---------------------------------------------------------------------------

func TestIsEqualTo(t *testing.T) {
	p := pred.IsEqualTo("123")

	validateredicate(t, p, &predicateExpectation{
		value:        "123",
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{`value == "123"`},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  "124",
		result: testpredicate.PredicateFailed,
	})
	validateredicate(t, p, &predicateExpectation{
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
	p := pred.IsNotEqualTo("123")

	validateredicate(t, p, &predicateExpectation{
		value:        "124",
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{`value != "123"`},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  "123",
		result: testpredicate.PredicateFailed,
	})
	validateredicate(t, p, &predicateExpectation{
		value:  124,
		result: testpredicate.PredicateInvalid,
		errMatchers: []string{
			`/values of type .* and .* are never equal/`,
		},
	})
}
