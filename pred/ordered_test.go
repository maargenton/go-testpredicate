package pred_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate"
	"github.com/maargenton/go-testpredicate/pred"
)

// ---------------------------------------------------------------------------
// pred.LessThan()
// ---------------------------------------------------------------------------

func TestLessThan(t *testing.T) {

	p := pred.Lt(123)
	validatePredicate(t, p, &predicateExpectation{
		value:        120,
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value < 123"},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:  123,
		result: testpredicate.PredicateFailed,
	})
	validatePredicate(t, p, &predicateExpectation{
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
	validatePredicate(t, p, &predicateExpectation{
		value:        123,
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value <= 123"},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:  124,
		result: testpredicate.PredicateFailed,
	})
	validatePredicate(t, p, &predicateExpectation{
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
	validatePredicate(t, p, &predicateExpectation{
		value:        124,
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value > 123"},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:  123,
		result: testpredicate.PredicateFailed,
	})
	validatePredicate(t, p, &predicateExpectation{
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
	validatePredicate(t, p, &predicateExpectation{
		value:        123,
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value >= 123"},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:  122,
		result: testpredicate.PredicateFailed,
	})
	validatePredicate(t, p, &predicateExpectation{
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
	validatePredicate(t, p, &predicateExpectation{
		value:        118,
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value ≈ 123 ± 5"},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:  128,
		result: testpredicate.PredicatePassed,
	})

	validatePredicate(t, p, &predicateExpectation{
		value:  117,
		result: testpredicate.PredicateFailed,
	})
	validatePredicate(t, p, &predicateExpectation{
		value:  129,
		result: testpredicate.PredicateFailed,
	})

	validatePredicate(t, p, &predicateExpectation{
		value:  "abc",
		result: testpredicate.PredicateInvalid,
		errMatchers: []string{
			`value "abc" of type string cannot be converted to float`,
		},
	})
}

func TestCloseToEx(t *testing.T) {

	p := pred.CloseTo([3]float64{1, 2, 3}, 0.2)
	validatePredicate(t, p, &predicateExpectation{
		value:        [3]float64{1, 2.1, 3},
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value ≈ [1 2 3] ± 0.2"},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:        [3]float64{1, 2.3, 3},
		result:       testpredicate.PredicateFailed,
		descMatchers: []string{"value ≈ [1 2 3] ± 0.2"},
	})
}
