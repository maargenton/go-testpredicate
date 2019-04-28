package pred_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate"
	"github.com/maargenton/go-testpredicate/pred"
)

// ---------------------------------------------------------------------------
// pred.StartsWith()
// ---------------------------------------------------------------------------
func TestMatch(t *testing.T) {
	p := pred.Matches(`aaa \d+ bbb`)

	validateredicate(t, p, &predicateExpectation{
		value:        "aaa 123 bbb!",
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value matches /aaa \\d+ bbb/"},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  123,
		result: testpredicate.PredicateInvalid,
		errMatchers: []string{
			"value of type int cannot be matched against a regexp",
		},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  "aaa 0x1f3 bbb!",
		result: testpredicate.PredicateFailed,
	})

	p1 := pred.Matches(`aaa (\d+ bbb`)
	validateredicate(t, p1, &predicateExpectation{
		value:  "aaa 123 bbb!",
		result: testpredicate.PredicateInvalid,
		errMatchers: []string{
			"failed to compile regexp",
		},
	})
}

// ---------------------------------------------------------------------------
// pred.ToUpper()
// ---------------------------------------------------------------------------

func TestToUpper(t *testing.T) {
	p := pred.ToUpper(pred.StartsWith("ABC"))

	validateredicate(t, p, &predicateExpectation{
		value:        "AbCdEf",
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"ToUpper(value) starts with"},
		errMatchers:  []string{`/.*/`},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  123,
		result: testpredicate.PredicateInvalid,
		errMatchers: []string{
			"value of type int cannot be transformed to uppercase",
		},
	})
	validateredicate(t, p, &predicateExpectation{
		value:       "aaa 0x1f3 bbb!",
		result:      testpredicate.PredicateFailed,
		errMatchers: []string{`/.*/`},
	})
}

// ---------------------------------------------------------------------------
// pred.ToLower()
// ---------------------------------------------------------------------------

func TestToLower(t *testing.T) {
	p := pred.ToLower(pred.StartsWith("abc"))

	validateredicate(t, p, &predicateExpectation{
		value:        "AbCdEf",
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"ToLower(value) starts with"},
		errMatchers:  []string{`/.*/`},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  123,
		result: testpredicate.PredicateInvalid,
		errMatchers: []string{
			"value of type int cannot be transformed to uppercase",
		},
	})
	validateredicate(t, p, &predicateExpectation{
		value:       "aaa 0x1f3 bbb!",
		result:      testpredicate.PredicateFailed,
		errMatchers: []string{`/.*/`},
	})
}

// ---------------------------------------------------------------------------
// pred.ToString()
// ---------------------------------------------------------------------------

func TestToString(t *testing.T) {
	p := pred.ToString(pred.StartsWith("123"))

	validateredicate(t, p, &predicateExpectation{
		value:        12345,
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"ToString(value) starts with"},
		errMatchers:  []string{`/.*/`},
	})
	validateredicate(t, p, &predicateExpectation{
		value:       54321,
		result:      testpredicate.PredicateFailed,
		errMatchers: []string{`/.*/`},
	})
}
