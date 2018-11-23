package pred_test

import (
	"testing"

	"github.com/marcus999/go-testpredicate"
	"github.com/marcus999/go-testpredicate/pred"
)

// ---------------------------------------------------------------------------
// pred.LessThan()
// ---------------------------------------------------------------------------

func TestLessThan(t *testing.T) {

	p := pred.Lt(123)
	validateredicate(t, p, &predicateExpectation{
		value:        120,
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value < 123"},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  123,
		result: testpredicate.PredicateFailed,
	})
	validateredicate(t, p, &predicateExpectation{
		value:  "abc",
		result: testpredicate.PredicateInvalid,
		errMatchers: []string{
			`/values of type .* and .* are not order comparable/`,
		},
	})
}

// ---------------------------------------------------------------------------
// pred.LessOrEqual()
// ---------------------------------------------------------------------------

func TestLessOrEqual(t *testing.T) {

	p := pred.Le(123)
	validateredicate(t, p, &predicateExpectation{
		value:        123,
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value <= 123"},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  124,
		result: testpredicate.PredicateFailed,
	})
	validateredicate(t, p, &predicateExpectation{
		value:  "abc",
		result: testpredicate.PredicateInvalid,
		errMatchers: []string{
			`/values of type .* and .* are not order comparable/`,
		},
	})
}

// ---------------------------------------------------------------------------
// pred.GreaterThan()
// ---------------------------------------------------------------------------

func TestGreaterThan(t *testing.T) {

	p := pred.Gt(123)
	validateredicate(t, p, &predicateExpectation{
		value:        124,
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value > 123"},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  123,
		result: testpredicate.PredicateFailed,
	})
	validateredicate(t, p, &predicateExpectation{
		value:  "abc",
		result: testpredicate.PredicateInvalid,
		errMatchers: []string{
			`/values of type .* and .* are not order comparable/`,
		},
	})
}

// ---------------------------------------------------------------------------
// pred.GreaterOrEqual()
// ---------------------------------------------------------------------------

func TestGreaterOrEqual(t *testing.T) {

	p := pred.Ge(123)
	validateredicate(t, p, &predicateExpectation{
		value:        123,
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value >= 123"},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  122,
		result: testpredicate.PredicateFailed,
	})
	validateredicate(t, p, &predicateExpectation{
		value:  "abc",
		result: testpredicate.PredicateInvalid,
		errMatchers: []string{
			`/values of type .* and .* are not order comparable/`},
	})
}

// ---------------------------------------------------------------------------
// pred.CloseTo()
// ---------------------------------------------------------------------------

func TestCloseTo(t *testing.T) {

	p := pred.CloseTo(123, 5)
	validateredicate(t, p, &predicateExpectation{
		value:        118,
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value ≈ 123 ± 5"},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  128,
		result: testpredicate.PredicatePassed,
	})

	validateredicate(t, p, &predicateExpectation{
		value:  117,
		result: testpredicate.PredicateFailed,
	})
	validateredicate(t, p, &predicateExpectation{
		value:  129,
		result: testpredicate.PredicateFailed,
	})

	validateredicate(t, p, &predicateExpectation{
		value:  "abc",
		result: testpredicate.PredicateInvalid,
		errMatchers: []string{
			`value "abc" of type string cannot be converted to float`,
		},
	})
}
