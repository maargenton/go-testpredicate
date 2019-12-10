package p_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/p"
	"github.com/maargenton/go-testpredicate/pkg/predicate"
)

// ---------------------------------------------------------------------------
// p.All
// ---------------------------------------------------------------------------
func TestAll(t *testing.T) {
	pred := p.All(p.Lt(10))

	validatePredicate(t, pred, &predicateExpectation{
		value:        []int{1, 2, 5},
		result:       predicate.Passed,
		descMatchers: []string{"all of value < 10"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  []int{1, 12, 5, 10},
		result: predicate.Failed,
		errMatchers: []string{
			"failed for value[1]: 12",
		},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  123,
		result: predicate.Invalid,
		errMatchers: []string{
			"/value of type .* is not a container/",
		},
	})

	var p2 = p.All(p.Length(p.Le(3)))
	validatePredicate(t, p2, &predicateExpectation{
		value:  []string{"a", "bb", "ccc", "dddd", "eeeee"},
		result: predicate.Failed,
		errMatchers: []string{
			"failed for value[3]:",
			"length: 4",
		},
	})
}

// ---------------------------------------------------------------------------
// p.Any
// ---------------------------------------------------------------------------
func TestAny(t *testing.T) {
	pred := p.Any(p.Gt(10))

	validatePredicate(t, pred, &predicateExpectation{
		value:        []int{1, 12, 5},
		result:       predicate.Passed,
		descMatchers: []string{"any of value > 10"},
		errMatchers: []string{
			"passed for value[1]: 12",
		},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  []int{1, 5, 10},
		result: predicate.Failed,
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  123,
		result: predicate.Invalid,
		errMatchers: []string{
			"/value of type .* is not a container/",
		},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  []interface{}{1, "abc", 345},
		result: predicate.Invalid,
		errMatchers: []string{
			`failed for value[1]: "abc"`,
			`/values of type 'string' and 'int' are not .* comparable/`,
		},
	})
}

// ---------------------------------------------------------------------------
// p.AllKeys
// ---------------------------------------------------------------------------
func TestAllKeys(t *testing.T) {
	pred := p.AllKeys(p.Lt(10))

	validatePredicate(t, pred, &predicateExpectation{
		value:        map[int]int{1: 1, 2: 2, 5: 5},
		result:       predicate.Passed,
		descMatchers: []string{"all of value.Keys() < 10"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  map[int]int{1: 1, 12: 2, 5: 5},
		result: predicate.Failed,
		errMatchers: []string{
			"failed for key 12",
		},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  []int{1, 2, 3},
		result: predicate.Invalid,
		errMatchers: []string{
			"/value of type .* is not a map/",
		},
	})
}

// ---------------------------------------------------------------------------
// p.AnyKey
// ---------------------------------------------------------------------------
func TestAnyKey(t *testing.T) {
	pred := p.AnyKey(p.Lt(10))

	validatePredicate(t, pred, &predicateExpectation{
		value:        map[int]int{10: 1, 20: 2, 5: 5},
		result:       predicate.Passed,
		descMatchers: []string{"any of value.Keys() < 10"},
		errMatchers:  []string{`/.*/`},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  map[int]int{10: 1, 20: 2, 50: 5},
		result: predicate.Failed,
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  []int{1, 2, 3},
		result: predicate.Invalid,
		errMatchers: []string{
			"/value of type .* is not a map/",
		},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  map[string]int{"10": 1, "20": 2, "50": 5},
		result: predicate.Invalid,
		errMatchers: []string{
			`/failed for key: "\d+"/`,
			`/values of type .* and .* are not order comparable/`,
		},
	})
}

// ---------------------------------------------------------------------------
// p.AllValues
// ---------------------------------------------------------------------------
func TestAllValues(t *testing.T) {
	pred := p.AllValues(p.Lt(10))

	validatePredicate(t, pred, &predicateExpectation{
		value:        map[int]int{1: 1, 2: 2, 5: 5},
		result:       predicate.Passed,
		descMatchers: []string{"all of value.Values() < 10"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  map[int]int{1: 1, 2: 12, 5: 5},
		result: predicate.Failed,
		errMatchers: []string{
			"failed for value[2]: 12",
		},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  []int{1, 2, 3},
		result: predicate.Invalid,
		errMatchers: []string{
			"/value of type .* is not a map/",
		},
	})
}

// ---------------------------------------------------------------------------
// p.AnyValue
// ---------------------------------------------------------------------------
func TestAnyValue(t *testing.T) {
	pred := p.AnyValue(p.Lt(10))

	validatePredicate(t, pred, &predicateExpectation{
		value:        map[int]int{1: 10, 2: 20, 5: 5},
		result:       predicate.Passed,
		descMatchers: []string{"any of value.Values() < 10"},
		errMatchers:  []string{`/.*/`},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  map[int]int{1: 10, 2: 20, 5: 50},
		result: predicate.Failed,
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  []int{1, 2, 3},
		result: predicate.Invalid,
		errMatchers: []string{
			"/value of type .* is not a map/",
		},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  map[int]string{1: "10", 2: "20", 5: "50"},
		result: predicate.Invalid,
		errMatchers: []string{
			`/failed for value\[\d+\]: "\d+"/`,
			`/values of type .* and .* are not order comparable/`,
		},
	})
}

// ---------------------------------------------------------------------------
// p.MapKeys
// ---------------------------------------------------------------------------
func TestMapKeys(t *testing.T) {
	pred := p.MapKeys(p.IsSubsetOf([]string{"aaa", "bbb", "ccc"}))

	validatePredicate(t, pred, &predicateExpectation{
		value: map[string]int{
			"aaa": 1,
			"bbb": 2,
			"ccc": 3,
		},
		result:       predicate.Passed,
		descMatchers: []string{"value.Keys() is subset of "},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value: map[string]int{
			"aaa": 1,
			"bcb": 2,
			"cbc": 3,
		},
		result:      predicate.Failed,
		errMatchers: []string{`extra values:`},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:       []int{1, 4, 3},
		result:      predicate.Invalid,
		errMatchers: []string{`/value of type .* is not a map/`},
	})
}

// ---------------------------------------------------------------------------
// p.MapKeys
// ---------------------------------------------------------------------------
func TestMapValues(t *testing.T) {
	pred := p.MapValues(p.IsSubsetOf([]int{1, 2, 3}))

	validatePredicate(t, pred, &predicateExpectation{
		value: map[string]int{
			"aaa": 1,
			"bbb": 2,
			"ccc": 3,
		},
		result:       predicate.Passed,
		descMatchers: []string{"value.Values() is subset of "},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value: map[string]int{
			"aaa": 1,
			"bcb": 4,
			"cbc": 3,
		},
		result:      predicate.Failed,
		errMatchers: []string{`extra values: 4`},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:       []int{1, 4, 3},
		result:      predicate.Invalid,
		errMatchers: []string{`/value of type .* is not a map/`},
	})
}
