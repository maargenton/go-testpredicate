package p_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/p"
	"github.com/maargenton/go-testpredicate/pkg/predicate"
)

// ---------------------------------------------------------------------------
// p.IsSubsetOf
// ---------------------------------------------------------------------------
func TestIsSubsetOf(t *testing.T) {

	pred := p.IsSubsetOf([]int{1, 2, 3})
	validatePredicate(t, pred, &predicateExpectation{
		value:        []int{3, 1},
		result:       predicate.Passed,
		descMatchers: []string{"value is subset of []int{ 1, 2, 3 }"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:       []int{3, 4, 1},
		result:      predicate.Failed,
		errMatchers: []string{`extra values: 4`},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:       map[int]bool{3: true, 4: false, 1: true},
		result:      predicate.Invalid,
		errMatchers: []string{`/value .* of type .* is not a indexable collection/`},
	})

	p1 := p.IsSubsetOf("abc")
	validatePredicate(t, p1, &predicateExpectation{
		value:        "ab",
		result:       predicate.Passed,
		descMatchers: []string{`value is subset of "abc"`},
	})

	p2 := p.IsSubsetOf(123)
	validatePredicate(t, p2, &predicateExpectation{
		value:        "ab",
		result:       predicate.Invalid,
		descMatchers: []string{`value is subset of 123`},
		errMatchers:  []string{`value 123 of type int is not a indexable collection`},
	})

	p3 := p.IsSubsetOf([]string{"bbb", "abc"})
	validatePredicate(t, p3, &predicateExpectation{
		value: []string{
			"bbb",
			"abc",
			"aaaaaaaaaaaaaaaaaaaaaaaa",
			"ccccccccccccccccccccccccccccccccccccccccccccc",
			"abdabdabdabdabdabdabdabdabdabdabdabd",
		},
		result:       predicate.Failed,
		descMatchers: []string{`/value is subset of .*/`},
		errMatchers: []string{
			`extra values:`,
			`/\.\.\.$/`,
		},
	})
}

// ---------------------------------------------------------------------------
// p.IsSupersetOf
// ---------------------------------------------------------------------------
func TestIsSupersetOf(t *testing.T) {

	pred := p.IsSupersetOf([]int{1, 2, 3})
	validatePredicate(t, pred, &predicateExpectation{
		value:        []int{3, 1, 4, 2},
		result:       predicate.Passed,
		descMatchers: []string{"value is superset of []int{ 1, 2, 3 }"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  []int{3},
		result: predicate.Failed,
		errMatchers: []string{
			`missing values:`,
			`/1/`,
			`/2/`,
		},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:       map[int]bool{3: true, 4: false, 1: true},
		result:      predicate.Invalid,
		errMatchers: []string{`/value .* of type .* is not a indexable collection/`},
	})

	p1 := p.IsSupersetOf("abc")
	validatePredicate(t, p1, &predicateExpectation{
		value:        "ab",
		result:       predicate.Failed,
		descMatchers: []string{`value is superset of "abc"`},
		errMatchers: []string{
			`missing values:`,
			`/0x63/`,
		},
	})

	p2 := p.IsSupersetOf(123)
	validatePredicate(t, p2, &predicateExpectation{
		value:        "ab",
		result:       predicate.Invalid,
		descMatchers: []string{`value is superset of 123`},
		errMatchers:  []string{`value 123 of type int is not a indexable collection`},
	})

	p3 := p.IsSupersetOf([]string{
		"aaaaaaaaaaaaaaaaaaaaaaaa",
		"bbb",
		"ccccccccccccccccccccccccccccccccccccccccccccc",
		"abc",
		"abdabdabdabdabdabdabdabdabdabdabdabd",
	})
	validatePredicate(t, p3, &predicateExpectation{
		value:        []string{"bbb", "abc"},
		result:       predicate.Failed,
		descMatchers: []string{`/value is superset of .*/`},
		errMatchers: []string{
			`missing values:`,
			`/\.\.\.$/`,
		},
	})
}

// ---------------------------------------------------------------------------
// p.IsDisjointSetFrom
// ---------------------------------------------------------------------------
func TestIsDisjointSetFrom(t *testing.T) {

	pred := p.IsDisjointSetFrom([]int{1, 3, 2})
	validatePredicate(t, pred, &predicateExpectation{
		value:        []int{5, 4, 6},
		result:       predicate.Passed,
		descMatchers: []string{`/value is disjoint from .*/`},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  []int{1, 6, 2},
		result: predicate.Failed,
		errMatchers: []string{
			`common values:`,
			`/1/`,
			`/2/`,
		},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:       map[int]bool{3: true, 4: false, 1: true},
		result:      predicate.Invalid,
		errMatchers: []string{`/value .* of type .* is not a indexable collection/`},
	})

	p2 := p.IsDisjointSetFrom(123)
	validatePredicate(t, p2, &predicateExpectation{
		value:        "ab",
		result:       predicate.Invalid,
		descMatchers: []string{`value is disjoint from 123`},
		errMatchers:  []string{`value 123 of type int is not a indexable collection`},
	})
}

// ---------------------------------------------------------------------------
// p.IsEqualSet
// ---------------------------------------------------------------------------
func TestIsEqualSet(t *testing.T) {

	pred := p.IsEqualSet([]int{1, 3, 2})
	validatePredicate(t, pred, &predicateExpectation{
		value:        []int{1, 2, 3},
		result:       predicate.Passed,
		descMatchers: []string{`/value is equal set as .*/`},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  []int{1, 6, 2},
		result: predicate.Failed,
		errMatchers: []string{
			`missing values: 3`,
			`extra values: 6`,
		},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:       map[int]bool{3: true, 4: false, 1: true},
		result:      predicate.Invalid,
		errMatchers: []string{`/value .* of type .* is not a indexable collection/`},
	})

	p2 := p.IsEqualSet(123)
	validatePredicate(t, p2, &predicateExpectation{
		value:        "ab",
		result:       predicate.Invalid,
		descMatchers: []string{`value is equal set as 123`},
		errMatchers:  []string{`value 123 of type int is not a indexable collection`},
	})
}
