package p_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/p"
	"github.com/maargenton/go-testpredicate/pkg/predicate"
)

// ---------------------------------------------------------------------------
// p.LessThan()
// ---------------------------------------------------------------------------

func TestLessThan(t *testing.T) {

	pred := p.Lt(123)
	validatePredicate(t, pred, &predicateExpectation{
		value:        120,
		result:       predicate.Passed,
		descMatchers: []string{"value < 123"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  123,
		result: predicate.Failed,
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  "abc",
		result: predicate.Invalid,
		errMatchers: []string{
			`/values of type .* and .* are not order comparable/`,
		},
	})
}

// ---------------------------------------------------------------------------
// p.LessOrEqual()
// ---------------------------------------------------------------------------

func TestLessOrEqual(t *testing.T) {

	pred := p.Le(123)
	validatePredicate(t, pred, &predicateExpectation{
		value:        123,
		result:       predicate.Passed,
		descMatchers: []string{"value <= 123"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  124,
		result: predicate.Failed,
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  "abc",
		result: predicate.Invalid,
		errMatchers: []string{
			`/values of type .* and .* are not order comparable/`,
		},
	})
}

// ---------------------------------------------------------------------------
// p.GreaterThan()
// ---------------------------------------------------------------------------

func TestGreaterThan(t *testing.T) {

	pred := p.Gt(123)
	validatePredicate(t, pred, &predicateExpectation{
		value:        124,
		result:       predicate.Passed,
		descMatchers: []string{"value > 123"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  123,
		result: predicate.Failed,
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  "abc",
		result: predicate.Invalid,
		errMatchers: []string{
			`/values of type .* and .* are not order comparable/`,
		},
	})
}

// ---------------------------------------------------------------------------
// p.GreaterOrEqual()
// ---------------------------------------------------------------------------

func TestGreaterOrEqual(t *testing.T) {

	pred := p.Ge(123)
	validatePredicate(t, pred, &predicateExpectation{
		value:        123,
		result:       predicate.Passed,
		descMatchers: []string{"value >= 123"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  122,
		result: predicate.Failed,
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  "abc",
		result: predicate.Invalid,
		errMatchers: []string{
			`/values of type .* and .* are not order comparable/`},
	})
}

// ---------------------------------------------------------------------------
// p.CloseTo()
// ---------------------------------------------------------------------------

func TestCloseTo(t *testing.T) {

	pred := p.CloseTo(123, 5)
	validatePredicate(t, pred, &predicateExpectation{
		value:        118,
		result:       predicate.Passed,
		descMatchers: []string{"value ≈ 123 ± 5"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  128,
		result: predicate.Passed,
	})

	validatePredicate(t, pred, &predicateExpectation{
		value:  117,
		result: predicate.Failed,
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  129,
		result: predicate.Failed,
	})

	validatePredicate(t, pred, &predicateExpectation{
		value:  "abc",
		result: predicate.Invalid,
		errMatchers: []string{
			`value of type 'string' cannot be converted to float`,
		},
	})
}

func TestCloseToEx(t *testing.T) {

	pred := p.CloseTo([3]float64{1, 2, 3}, 0.2)
	validatePredicate(t, pred, &predicateExpectation{
		value:        [3]float64{1, 2.1, 3},
		result:       predicate.Passed,
		descMatchers: []string{"value ≈ [1 2 3] ± 0.2"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:        [3]float64{1, 2.3, 3},
		result:       predicate.Failed,
		descMatchers: []string{"value ≈ [1 2 3] ± 0.2"},
	})
}
