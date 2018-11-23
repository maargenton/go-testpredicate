package pred_test

import (
	"testing"

	"github.com/marcus999/go-testpredicate"
	"github.com/marcus999/go-testpredicate/pred"
)

// ---------------------------------------------------------------------------
// pred.IsSubsetOf
// ---------------------------------------------------------------------------
func TestIsSubsetOf(t *testing.T) {

	p := pred.IsSubsetOf([]int{1, 2, 3})
	validateredicate(t, p, &predicateExpectation{
		value:        []int{3, 1},
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value is subset of []int{1, 2, 3}"},
	})
	validateredicate(t, p, &predicateExpectation{
		value:       []int{3, 4, 1},
		result:      testpredicate.PredicateFailed,
		errMatchers: []string{`extra values: 4`},
	})
	validateredicate(t, p, &predicateExpectation{
		value:       map[int]bool{3: true, 4: false, 1: true},
		result:      testpredicate.PredicateInvalid,
		errMatchers: []string{`/value .* of type .* is not a indexable collection/`},
	})

	p1 := pred.IsSubsetOf("abc")
	validateredicate(t, p1, &predicateExpectation{
		value:        "ab",
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{`value is subset of "abc"`},
	})

	p2 := pred.IsSubsetOf(123)
	validateredicate(t, p2, &predicateExpectation{
		value:        "ab",
		result:       testpredicate.PredicateInvalid,
		descMatchers: []string{`value is subset of 123`},
		errMatchers:  []string{`value 123 of type int is not a indexable collection`},
	})

	p3 := pred.IsSubsetOf([]string{"bbb", "abc"})
	validateredicate(t, p3, &predicateExpectation{
		value: []string{
			"bbb",
			"abc",
			"aaaaaaaaaaaaaaaaaaaaaaaa",
			"ccccccccccccccccccccccccccccccccccccccccccccc",
			"abdabdabdabdabdabdabdabdabdabdabdabd",
		},
		result:       testpredicate.PredicateFailed,
		descMatchers: []string{`/value is subset of .*/`},
		errMatchers: []string{
			`extra values:`,
			`/\.\.\.$/`,
		},
	})
}

// ---------------------------------------------------------------------------
// pred.IsSupersetOf
// ---------------------------------------------------------------------------
func TestIsSupersetOf(t *testing.T) {

	p := pred.IsSupersetOf([]int{1, 2, 3})
	validateredicate(t, p, &predicateExpectation{
		value:        []int{3, 1, 4, 2},
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value is superset of []int{1, 2, 3}"},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  []int{3},
		result: testpredicate.PredicateFailed,
		errMatchers: []string{
			`missing values:`,
			`/1/`,
			`/2/`,
		},
	})
	validateredicate(t, p, &predicateExpectation{
		value:       map[int]bool{3: true, 4: false, 1: true},
		result:      testpredicate.PredicateInvalid,
		errMatchers: []string{`/value .* of type .* is not a indexable collection/`},
	})

	p1 := pred.IsSupersetOf("abc")
	validateredicate(t, p1, &predicateExpectation{
		value:        "ab",
		result:       testpredicate.PredicateFailed,
		descMatchers: []string{`value is superset of "abc"`},
		errMatchers: []string{
			`missing values:`,
			`/0x63/`,
		},
	})

	p2 := pred.IsSupersetOf(123)
	validateredicate(t, p2, &predicateExpectation{
		value:        "ab",
		result:       testpredicate.PredicateInvalid,
		descMatchers: []string{`value is superset of 123`},
		errMatchers:  []string{`value 123 of type int is not a indexable collection`},
	})

	p3 := pred.IsSupersetOf([]string{
		"aaaaaaaaaaaaaaaaaaaaaaaa",
		"bbb",
		"ccccccccccccccccccccccccccccccccccccccccccccc",
		"abc",
		"abdabdabdabdabdabdabdabdabdabdabdabd",
	})
	validateredicate(t, p3, &predicateExpectation{
		value:        []string{"bbb", "abc"},
		result:       testpredicate.PredicateFailed,
		descMatchers: []string{`/value is superset of .*/`},
		errMatchers: []string{
			`missing values:`,
			`/\.\.\.$/`,
		},
	})
}

// ---------------------------------------------------------------------------
// pred.IsDisjointSetFrom
// ---------------------------------------------------------------------------
func TestIsDisjointSetFrom(t *testing.T) {

	p := pred.IsDisjointSetFrom([]int{1, 3, 2})
	validateredicate(t, p, &predicateExpectation{
		value:        []int{5, 4, 6},
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{`/value is disjoint from .*/`},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  []int{1, 6, 2},
		result: testpredicate.PredicateFailed,
		errMatchers: []string{
			`common values:`,
			`/1/`,
			`/2/`,
		},
	})
	validateredicate(t, p, &predicateExpectation{
		value:       map[int]bool{3: true, 4: false, 1: true},
		result:      testpredicate.PredicateInvalid,
		errMatchers: []string{`/value .* of type .* is not a indexable collection/`},
	})

	p2 := pred.IsDisjointSetFrom(123)
	validateredicate(t, p2, &predicateExpectation{
		value:        "ab",
		result:       testpredicate.PredicateInvalid,
		descMatchers: []string{`value is disjoint from 123`},
		errMatchers:  []string{`value 123 of type int is not a indexable collection`},
	})
}

// ---------------------------------------------------------------------------
// pred.IsEqualSet
// ---------------------------------------------------------------------------
func TestIsEqualSet(t *testing.T) {

	p := pred.IsEqualSet([]int{1, 3, 2})
	validateredicate(t, p, &predicateExpectation{
		value:        []int{1, 2, 3},
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{`/value is equal set as .*/`},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  []int{1, 6, 2},
		result: testpredicate.PredicateFailed,
		errMatchers: []string{
			`missing values: 3`,
			`extra values: 6`,
		},
	})
	validateredicate(t, p, &predicateExpectation{
		value:       map[int]bool{3: true, 4: false, 1: true},
		result:      testpredicate.PredicateInvalid,
		errMatchers: []string{`/value .* of type .* is not a indexable collection/`},
	})

	p2 := pred.IsEqualSet(123)
	validateredicate(t, p2, &predicateExpectation{
		value:        "ab",
		result:       testpredicate.PredicateInvalid,
		descMatchers: []string{`value is equal set as 123`},
		errMatchers:  []string{`value 123 of type int is not a indexable collection`},
	})
}
