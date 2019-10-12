package pred_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate"
	"github.com/maargenton/go-testpredicate/pred"
)

func TestPanics(t *testing.T) {

	p := pred.Panics()
	validatePredicate(t, p, &predicateExpectation{
		value:        func() { panic(123) },
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"fct() panics"},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:       func() {},
		result:      testpredicate.PredicateFailed,
		errMatchers: []string{"did not panic"},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:       func(i int) { panic(i) },
		result:      testpredicate.PredicateInvalid,
		errMatchers: []string{"is not callable"},
	})
}

func TestPanicsAndResult(t *testing.T) {

	p := pred.PanicsAndResult(pred.Contains("aaa"))

	validatePredicate(t, p, &predicateExpectation{
		value:        func() { panic("something aaa something") },
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"fct() panics"},
		errMatchers:  []string{"panic: "},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:       func() {},
		result:      testpredicate.PredicateFailed,
		errMatchers: []string{"did not panic"},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:       func(i int) { panic(i) },
		result:      testpredicate.PredicateInvalid,
		errMatchers: []string{"is not callable"},
	})
	validatePredicate(t, p, &predicateExpectation{
		value:        func() { panic("something bbb something") },
		result:       testpredicate.PredicateFailed,
		descMatchers: []string{"fct() panics"},
		errMatchers:  []string{"panic: "},
	})
}
