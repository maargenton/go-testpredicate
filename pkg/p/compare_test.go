package p_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/p"
	"github.com/maargenton/go-testpredicate/pkg/predicate"
)

// ---------------------------------------------------------------------------
// p.IsNil()
// ---------------------------------------------------------------------------

func TestIsNil(t *testing.T) {
	p := p.IsNil()

	validatePredicate(t, p, &predicateExpectation{
		value:        nil,
		result:       predicate.Passed,
		descMatchers: []string{"value is nil"},
	})

	var ptr *int
	validatePredicate(t, p, &predicateExpectation{
		value:  ptr,
		result: predicate.Passed,
	})
	validatePredicate(t, p, &predicateExpectation{
		value:  &ptr,
		result: predicate.Failed,
	})

	var slice []int
	validatePredicate(t, p, &predicateExpectation{
		value:  slice,
		result: predicate.Passed,
	})
	validatePredicate(t, p, &predicateExpectation{
		value:  &slice,
		result: predicate.Failed,
	})

	validatePredicate(t, p, &predicateExpectation{
		value:       123,
		result:      predicate.Invalid,
		errMatchers: []string{"value of type 'int' is never nil"},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:       "abc",
		result:      predicate.Invalid,
		errMatchers: []string{"value of type 'string' is never nil"},
	})
}

// ---------------------------------------------------------------------------
// p.IsNotNil()
// ---------------------------------------------------------------------------

func TestIsNotNil(t *testing.T) {
	p := p.IsNotNil()

	validatePredicate(t, p, &predicateExpectation{
		value:        nil,
		result:       predicate.Failed,
		descMatchers: []string{"value is not nil"},
	})

	var ptr *int
	validatePredicate(t, p, &predicateExpectation{
		value:  ptr,
		result: predicate.Failed,
	})
	validatePredicate(t, p, &predicateExpectation{
		value:  &ptr,
		result: predicate.Passed,
	})

	var slice []int
	validatePredicate(t, p, &predicateExpectation{
		value:  slice,
		result: predicate.Failed,
	})
	validatePredicate(t, p, &predicateExpectation{
		value:  &slice,
		result: predicate.Passed,
	})

	validatePredicate(t, p, &predicateExpectation{
		value:       123,
		result:      predicate.Invalid,
		errMatchers: []string{"value of type 'int' is never nil"},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:       "abc",
		result:      predicate.Invalid,
		errMatchers: []string{"value of type 'string' is never nil"},
	})
}

// ---------------------------------------------------------------------------
// p.IsEqualTo()
// ---------------------------------------------------------------------------

func TestIsEqualTo(t *testing.T) {
	p := p.Eq("123")

	validatePredicate(t, p, &predicateExpectation{
		value:        "123",
		result:       predicate.Passed,
		descMatchers: []string{`value == "123"`},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:  "124",
		result: predicate.Failed,
	})
	validatePredicate(t, p, &predicateExpectation{
		value:  124,
		result: predicate.Invalid,
		errMatchers: []string{
			`/values of type .* and .* are never equal/`,
		},
	})
}

// ---------------------------------------------------------------------------
// p.IsNotEqualTo()
// ---------------------------------------------------------------------------

func TestNotIsEqualTo(t *testing.T) {
	p := p.Ne("123")

	validatePredicate(t, p, &predicateExpectation{
		value:        "124",
		result:       predicate.Passed,
		descMatchers: []string{`value != "123"`},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:  "123",
		result: predicate.Failed,
	})
	validatePredicate(t, p, &predicateExpectation{
		value:  124,
		result: predicate.Invalid,
		errMatchers: []string{
			`/values of type .* and .* are never equal/`,
		},
	})
}

// ---------------------------------------------------------------------------
// p.IsTrue()
// ---------------------------------------------------------------------------

func TestIsTrue(t *testing.T) {
	p := p.IsTrue()

	validatePredicate(t, p, &predicateExpectation{
		value:        true,
		result:       predicate.Passed,
		descMatchers: []string{"value is true"},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:        false,
		result:       predicate.Failed,
		descMatchers: []string{"value is true"},
	})

	validatePredicate(t, p, &predicateExpectation{
		value:       123,
		result:      predicate.Invalid,
		errMatchers: []string{"value of type 'int' is never true"},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:       "abc",
		result:      predicate.Invalid,
		errMatchers: []string{"value of type 'string' is never true"},
	})
}

// ---------------------------------------------------------------------------
// p.IsFalse()
// ---------------------------------------------------------------------------

func TestIsFalse(t *testing.T) {
	p := p.IsFalse()

	validatePredicate(t, p, &predicateExpectation{
		value:        false,
		result:       predicate.Passed,
		descMatchers: []string{"value is false"},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:        true,
		result:       predicate.Failed,
		descMatchers: []string{"value is false"},
	})

	validatePredicate(t, p, &predicateExpectation{
		value:       123,
		result:      predicate.Invalid,
		errMatchers: []string{"value of type 'int' is never false"},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:       "abc",
		result:      predicate.Invalid,
		errMatchers: []string{"value of type 'string' is never false"},
	})
}
