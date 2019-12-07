package p_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/p"
	"github.com/maargenton/go-testpredicate/pkg/predicate"
)

// ---------------------------------------------------------------------------
// p.StartsWith()
// ---------------------------------------------------------------------------
func TestMatch(t *testing.T) {
	pred := p.Matches(`aaa \d+ bbb`)

	validatePredicate(t, pred, &predicateExpectation{
		value:        "aaa 123 bbb!",
		result:       predicate.Passed,
		descMatchers: []string{"value matches /aaa \\d+ bbb/"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  123,
		result: predicate.Invalid,
		errMatchers: []string{
			"value of type int cannot be matched against a regexp",
		},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  "aaa 0x1f3 bbb!",
		result: predicate.Failed,
	})

	p1 := p.Matches(`aaa (\d+ bbb`)
	validatePredicate(t, p1, &predicateExpectation{
		value:  "aaa 123 bbb!",
		result: predicate.Invalid,
		errMatchers: []string{
			"failed to compile regexp",
		},
	})
}

// ---------------------------------------------------------------------------
// p.ToUpper()
// ---------------------------------------------------------------------------

func TestToUpper(t *testing.T) {
	pred := p.ToUpper(p.StartsWith("ABC"))

	validatePredicate(t, pred, &predicateExpectation{
		value:        "AbCdEf",
		result:       predicate.Passed,
		descMatchers: []string{"ToUpper(value) starts with"},
		errMatchers:  []string{`/.*/`},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  123,
		result: predicate.Invalid,
		errMatchers: []string{
			"value of type int cannot be transformed to uppercase",
		},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:       "aaa 0x1f3 bbb!",
		result:      predicate.Failed,
		errMatchers: []string{`/.*/`},
	})
}

// ---------------------------------------------------------------------------
// p.ToLower()
// ---------------------------------------------------------------------------

func TestToLower(t *testing.T) {
	pred := p.ToLower(p.StartsWith("abc"))

	validatePredicate(t, pred, &predicateExpectation{
		value:        "AbCdEf",
		result:       predicate.Passed,
		descMatchers: []string{"ToLower(value) starts with"},
		errMatchers:  []string{`/.*/`},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  123,
		result: predicate.Invalid,
		errMatchers: []string{
			"value of type int cannot be transformed to uppercase",
		},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:       "aaa 0x1f3 bbb!",
		result:      predicate.Failed,
		errMatchers: []string{`/.*/`},
	})
}

// ---------------------------------------------------------------------------
// p.ToString()
// ---------------------------------------------------------------------------

func TestToString(t *testing.T) {
	pred := p.ToString(p.StartsWith("123"))

	validatePredicate(t, pred, &predicateExpectation{
		value:        12345,
		result:       predicate.Passed,
		descMatchers: []string{"ToString(value) starts with"},
		errMatchers:  []string{`/.*/`},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:       54321,
		result:      predicate.Failed,
		errMatchers: []string{`/.*/`},
	})
}
