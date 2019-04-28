package pred_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate"
	"github.com/maargenton/go-testpredicate/pred"
)

// ---------------------------------------------------------------------------
// pred.All
// ---------------------------------------------------------------------------
func TestAll(t *testing.T) {
	p := pred.All(pred.Lt(10))

	validateredicate(t, p, &predicateExpectation{
		value:        []int{1, 2, 5},
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"all of value < 10"},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  []int{1, 12, 5, 10},
		result: testpredicate.PredicateFailed,
		errMatchers: []string{
			"failed for value[1]: 12",
		},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  123,
		result: testpredicate.PredicateInvalid,
		errMatchers: []string{
			"/value .* of type .* is not a container/",
		},
	})
}

// ---------------------------------------------------------------------------
// pred.Any
// ---------------------------------------------------------------------------
func TestAny(t *testing.T) {
	p := pred.Any(pred.Gt(10))

	validateredicate(t, p, &predicateExpectation{
		value:        []int{1, 12, 5},
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"any of value > 10"},
		errMatchers: []string{
			"passed for value[1]: 12",
		},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  []int{1, 5, 10},
		result: testpredicate.PredicateFailed,
	})
	validateredicate(t, p, &predicateExpectation{
		value:  123,
		result: testpredicate.PredicateInvalid,
		errMatchers: []string{
			"/value .* of type .* is not a container/",
		},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  []interface{}{1, "abc", 345},
		result: testpredicate.PredicateInvalid,
		errMatchers: []string{
			`failed for value[1]: "abc"`,
			`/values of type string and int are not .* comparable/`,
		},
	})
}

// ---------------------------------------------------------------------------
// pred.AllKeys
// ---------------------------------------------------------------------------
func TestAllKeys(t *testing.T) {
	p := pred.AllKeys(pred.Lt(10))

	validateredicate(t, p, &predicateExpectation{
		value:        map[int]int{1: 1, 2: 2, 5: 5},
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"all of value.Keys() < 10"},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  map[int]int{1: 1, 12: 2, 5: 5},
		result: testpredicate.PredicateFailed,
		errMatchers: []string{
			"failed for key 12",
		},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  []int{1, 2, 3},
		result: testpredicate.PredicateInvalid,
		errMatchers: []string{
			"/value .* of type .* is not a map/",
		},
	})
}

// ---------------------------------------------------------------------------
// pred.AnyKey
// ---------------------------------------------------------------------------
func TestAnyKey(t *testing.T) {
	p := pred.AnyKey(pred.Lt(10))

	validateredicate(t, p, &predicateExpectation{
		value:        map[int]int{10: 1, 20: 2, 5: 5},
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"any of value.Keys() < 10"},
		errMatchers:  []string{`/.*/`},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  map[int]int{10: 1, 20: 2, 50: 5},
		result: testpredicate.PredicateFailed,
	})
	validateredicate(t, p, &predicateExpectation{
		value:  []int{1, 2, 3},
		result: testpredicate.PredicateInvalid,
		errMatchers: []string{
			"/value .* of type .* is not a map/",
		},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  map[string]int{"10": 1, "20": 2, "50": 5},
		result: testpredicate.PredicateInvalid,
		errMatchers: []string{
			`/failed for key: "\d+"/`,
			`/values of type .* and .* are not order comparable/`,
		},
	})
}

// ---------------------------------------------------------------------------
// pred.AllValues
// ---------------------------------------------------------------------------
func TestAllValues(t *testing.T) {
	p := pred.AllValues(pred.Lt(10))

	validateredicate(t, p, &predicateExpectation{
		value:        map[int]int{1: 1, 2: 2, 5: 5},
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"all of value.Values() < 10"},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  map[int]int{1: 1, 2: 12, 5: 5},
		result: testpredicate.PredicateFailed,
		errMatchers: []string{
			"failed for value[2]: 12",
		},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  []int{1, 2, 3},
		result: testpredicate.PredicateInvalid,
		errMatchers: []string{
			"/value .* of type .* is not a map/",
		},
	})
}

// ---------------------------------------------------------------------------
// pred.AnyValue
// ---------------------------------------------------------------------------
func TestAnyValue(t *testing.T) {
	p := pred.AnyValue(pred.Lt(10))

	validateredicate(t, p, &predicateExpectation{
		value:        map[int]int{1: 10, 2: 20, 5: 5},
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"any of value.Values() < 10"},
		errMatchers:  []string{`/.*/`},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  map[int]int{1: 10, 2: 20, 5: 50},
		result: testpredicate.PredicateFailed,
	})
	validateredicate(t, p, &predicateExpectation{
		value:  []int{1, 2, 3},
		result: testpredicate.PredicateInvalid,
		errMatchers: []string{
			"/value .* of type .* is not a map/",
		},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  map[int]string{1: "10", 2: "20", 5: "50"},
		result: testpredicate.PredicateInvalid,
		errMatchers: []string{
			`/failed for value\[\d+\]: "\d+"/`,
			`/values of type .* and .* are not order comparable/`,
		},
	})
}

// ---------------------------------------------------------------------------
// pred.MapKeys
// ---------------------------------------------------------------------------
func TestMapKeys(t *testing.T) {
	p := pred.MapKeys(pred.IsSubsetOf([]string{"aaa", "bbb", "ccc"}))

	validateredicate(t, p, &predicateExpectation{
		value: map[string]int{
			"aaa": 1,
			"bbb": 2,
			"ccc": 3,
		},
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value.Keys() is subset of "},
	})
	validateredicate(t, p, &predicateExpectation{
		value: map[string]int{
			"aaa": 1,
			"bcb": 2,
			"cbc": 3,
		},
		result:      testpredicate.PredicateFailed,
		errMatchers: []string{`extra values:`},
	})
	validateredicate(t, p, &predicateExpectation{
		value:       []int{1, 4, 3},
		result:      testpredicate.PredicateInvalid,
		errMatchers: []string{`/value .* of type .* is not a map/`},
	})
}

// ---------------------------------------------------------------------------
// pred.MapKeys
// ---------------------------------------------------------------------------
func TestMapValues(t *testing.T) {
	p := pred.MapValues(pred.IsSubsetOf([]int{1, 2, 3}))

	validateredicate(t, p, &predicateExpectation{
		value: map[string]int{
			"aaa": 1,
			"bbb": 2,
			"ccc": 3,
		},
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value.Values() is subset of "},
	})
	validateredicate(t, p, &predicateExpectation{
		value: map[string]int{
			"aaa": 1,
			"bcb": 4,
			"cbc": 3,
		},
		result:      testpredicate.PredicateFailed,
		errMatchers: []string{`extra values: 4`},
	})
	validateredicate(t, p, &predicateExpectation{
		value:       []int{1, 4, 3},
		result:      testpredicate.PredicateInvalid,
		errMatchers: []string{`/value .* of type .* is not a map/`},
	})
}
