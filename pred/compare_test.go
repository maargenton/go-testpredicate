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

	var ptr *int
	validatePredicate(t, p, &predicateExpectation{
		value:  ptr,
		result: testpredicate.PredicatePassed,
	})
	validatePredicate(t, p, &predicateExpectation{
		value:  &ptr,
		result: testpredicate.PredicateFailed,
	})

	var slice []int
	validatePredicate(t, p, &predicateExpectation{
		value:  slice,
		result: testpredicate.PredicatePassed,
	})
	validatePredicate(t, p, &predicateExpectation{
		value:  &slice,
		result: testpredicate.PredicateFailed,
	})

	validatePredicate(t, p, &predicateExpectation{
		value:       123,
		result:      testpredicate.PredicateInvalid,
		errMatchers: []string{"value of type 'int' is never nil"},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:       "abc",
		result:      testpredicate.PredicateInvalid,
		errMatchers: []string{"value of type 'string' is never nil"},
	})
}

// ---------------------------------------------------------------------------
// pred.IsNotNil()
// ---------------------------------------------------------------------------

func TestIsNotNil(t *testing.T) {
	p := pred.IsNotNil()

	validatePredicate(t, p, &predicateExpectation{
		value:        nil,
		result:       testpredicate.PredicateFailed,
		descMatchers: []string{"value is not nil"},
	})

	var ptr *int
	validatePredicate(t, p, &predicateExpectation{
		value:  ptr,
		result: testpredicate.PredicateFailed,
	})
	validatePredicate(t, p, &predicateExpectation{
		value:  &ptr,
		result: testpredicate.PredicatePassed,
	})

	var slice []int
	validatePredicate(t, p, &predicateExpectation{
		value:  slice,
		result: testpredicate.PredicateFailed,
	})
	validatePredicate(t, p, &predicateExpectation{
		value:  &slice,
		result: testpredicate.PredicatePassed,
	})

	validatePredicate(t, p, &predicateExpectation{
		value:       123,
		result:      testpredicate.PredicateInvalid,
		errMatchers: []string{"value of type 'int' is never nil"},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:       "abc",
		result:      testpredicate.PredicateInvalid,
		errMatchers: []string{"value of type 'string' is never nil"},
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

// ---------------------------------------------------------------------------
// pred.IsTrue()
// ---------------------------------------------------------------------------

func TestIsTrue(t *testing.T) {
	p := pred.IsTrue()

	validatePredicate(t, p, &predicateExpectation{
		value:        true,
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value is true"},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:        false,
		result:       testpredicate.PredicateFailed,
		descMatchers: []string{"value is true"},
	})

	validatePredicate(t, p, &predicateExpectation{
		value:       123,
		result:      testpredicate.PredicateInvalid,
		errMatchers: []string{"value of type 'int' is never true"},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:       "abc",
		result:      testpredicate.PredicateInvalid,
		errMatchers: []string{"value of type 'string' is never true"},
	})
}

// ---------------------------------------------------------------------------
// pred.IsFalse()
// ---------------------------------------------------------------------------

func TestIsFalse(t *testing.T) {
	p := pred.IsFalse()

	validatePredicate(t, p, &predicateExpectation{
		value:        false,
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value is false"},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:        true,
		result:       testpredicate.PredicateFailed,
		descMatchers: []string{"value is false"},
	})

	validatePredicate(t, p, &predicateExpectation{
		value:       123,
		result:      testpredicate.PredicateInvalid,
		errMatchers: []string{"value of type 'int' is never false"},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:       "abc",
		result:      testpredicate.PredicateInvalid,
		errMatchers: []string{"value of type 'string' is never false"},
	})
}
