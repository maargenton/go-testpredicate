package p_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/p"
	"github.com/maargenton/go-testpredicate/pkg/predicate"
)

func TestPanics(t *testing.T) {

	pred := p.Panics()
	validatePredicate(t, pred, &predicateExpectation{
		value:        func() { panic(123) },
		result:       predicate.Passed,
		descMatchers: []string{"fct() panics"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:       func() {},
		result:      predicate.Failed,
		errMatchers: []string{"did not panic"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:       func(i int) { panic(i) },
		result:      predicate.Invalid,
		errMatchers: []string{"is not callable"},
	})
}

func TestPanicsAndResult(t *testing.T) {

	pred := p.PanicsAndResult(p.Contains("aaa"))

	validatePredicate(t, pred, &predicateExpectation{
		value:        func() { panic("something aaa something") },
		result:       predicate.Passed,
		descMatchers: []string{"fct() panics"},
		errMatchers:  []string{"panic: "},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:       func() {},
		result:      predicate.Failed,
		errMatchers: []string{"did not panic"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:       func(i int) { panic(i) },
		result:      predicate.Invalid,
		errMatchers: []string{"is not callable"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:        func() { panic("something bbb something") },
		result:       predicate.Failed,
		descMatchers: []string{"fct() panics"},
		errMatchers:  []string{"panic: "},
	})
}
