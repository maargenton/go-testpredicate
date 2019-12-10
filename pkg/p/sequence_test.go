package p_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/p"
	"github.com/maargenton/go-testpredicate/pkg/predicate"
)

// ---------------------------------------------------------------------------
// p.IsEmpty()
// ---------------------------------------------------------------------------

func TestStringIsEmpty(t *testing.T) {
	pred := p.IsEmpty()

	validatePredicate(t, pred, &predicateExpectation{
		value:        []int{},
		result:       predicate.Passed,
		descMatchers: []string{"value is empty"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:       []int{1, 2, 3, 4, 5},
		result:      predicate.Failed,
		errMatchers: []string{"length: 5"},
	})
}

func TestSliceIsEmpty(t *testing.T) {
	pred := p.IsEmpty()

	validatePredicate(t, pred, &predicateExpectation{
		value:        "",
		result:       predicate.Passed,
		descMatchers: []string{"value is empty"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:       "Hello world!",
		result:      predicate.Failed,
		errMatchers: []string{"length: 12"},
	})
}

func TestMapIsEmpty(t *testing.T) {
	pred := p.IsEmpty()

	validatePredicate(t, pred, &predicateExpectation{
		value:        map[string]int{},
		result:       predicate.Passed,
		descMatchers: []string{"value is empty"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value: map[string]int{
			"aaa":    2,
			"bbbccc": 4,
		},
		result:      predicate.Failed,
		errMatchers: []string{"length: 2"},
	})
}

func TestIntIsEmpty(t *testing.T) {
	pred := p.IsEmpty()

	validatePredicate(t, pred, &predicateExpectation{
		value:        uint64(123),
		result:       predicate.Invalid,
		descMatchers: []string{"value is empty"},
		errMatchers:  []string{"value of type 'uint64' cannot be tested for emptiness"},
	})
}

// ---------------------------------------------------------------------------
// p.IsNotEmpty()
// ---------------------------------------------------------------------------

func TestStringIsNotEmpty(t *testing.T) {
	pred := p.IsNotEmpty()

	validatePredicate(t, pred, &predicateExpectation{
		value:        []int{1, 2, 3, 4, 5},
		result:       predicate.Passed,
		descMatchers: []string{"value is not empty"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  []int{},
		result: predicate.Failed,
	})
}

func TestSliceIsNotEmpty(t *testing.T) {
	pred := p.IsNotEmpty()

	validatePredicate(t, pred, &predicateExpectation{
		value:        "Hello world!",
		result:       predicate.Passed,
		descMatchers: []string{"value is not empty"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  "",
		result: predicate.Failed,
	})
}

func TestMapIsNotEmpty(t *testing.T) {
	pred := p.IsNotEmpty()

	validatePredicate(t, pred, &predicateExpectation{
		value: map[string]int{
			"aaa":    2,
			"bbbccc": 4,
		},
		result:       predicate.Passed,
		descMatchers: []string{"value is not empty"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  map[string]int{},
		result: predicate.Failed,
	})
}

func TestIntIsNotEmpty(t *testing.T) {
	pred := p.IsNotEmpty()

	validatePredicate(t, pred, &predicateExpectation{
		value:        uint64(123),
		result:       predicate.Invalid,
		descMatchers: []string{"value is not empty"},
		errMatchers:  []string{`/value of type .* cannot be tested for emptiness/`},
	})
}

// ---------------------------------------------------------------------------
// p.StartsWith()
// ---------------------------------------------------------------------------
func TestStringStartsWith(t *testing.T) {
	pred := p.StartsWith("Hello")

	validatePredicate(t, pred, &predicateExpectation{
		value:        "Hello world!",
		result:       predicate.Passed,
		descMatchers: []string{"value starts with \"Hello\""},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:       "Hey",
		result:      predicate.Failed,
		errMatchers: []string{`/sequence .* too small/`},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  "Hey world",
		result: predicate.Failed,
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:       123,
		result:      predicate.Invalid,
		errMatchers: []string{`/value of type .* is not a sequence/`},
	})
}

// ---------------------------------------------------------------------------
// p.Contains()
// ---------------------------------------------------------------------------

func TestStringContains(t *testing.T) {
	pred := p.Contains("worl")

	validatePredicate(t, pred, &predicateExpectation{
		value:        "Hello world!",
		result:       predicate.Passed,
		descMatchers: []string{"value contains \"worl\""},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:       "Hey",
		result:      predicate.Failed,
		errMatchers: []string{`/sequence .* too small/`},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  "Hey universe",
		result: predicate.Failed,
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:       123,
		result:      predicate.Invalid,
		errMatchers: []string{`/value of type .* is not a sequence/`},
	})
}

// ---------------------------------------------------------------------------
// p.StartsWith()
// ---------------------------------------------------------------------------
func TestStringEndsWith(t *testing.T) {
	pred := p.EndsWith("world!")

	validatePredicate(t, pred, &predicateExpectation{
		value:        "Hello world!",
		result:       predicate.Passed,
		descMatchers: []string{"value ends with \"world!\""},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:       "Hey",
		result:      predicate.Failed,
		errMatchers: []string{`/sequence .* too small/`},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:  "Hey world",
		result: predicate.Failed,
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:       123,
		result:      predicate.Invalid,
		errMatchers: []string{`/value of type .* is not a sequence/`},
	})
}

// ---------------------------------------------------------------------------
// p.Length()
// ---------------------------------------------------------------------------

func TestSliceLength(t *testing.T) {
	pred := p.Length(p.Lt(3))

	validatePredicate(t, pred, &predicateExpectation{
		value:        []int{1, 2},
		result:       predicate.Passed,
		descMatchers: []string{"length(value) < 3"},
		errMatchers:  []string{"length: 2"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:        []int{1, 2, 3, 4, 5},
		result:       predicate.Failed,
		descMatchers: []string{"length(value) < 3"},
		errMatchers:  []string{"length: 5"},
	})
}

func TestStringLength(t *testing.T) {
	pred := p.Length(p.Lt(3))

	validatePredicate(t, pred, &predicateExpectation{
		value:        "Yo",
		result:       predicate.Passed,
		descMatchers: []string{"length(value) < 3"},
		errMatchers:  []string{"length: 2"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:        "Hello world!",
		result:       predicate.Failed,
		descMatchers: []string{"length(value) < 3"},
		errMatchers:  []string{"length: 12"},
	})
}

func TestIntLength(t *testing.T) {
	pred := p.Length(p.LessThan(3))
	validatePredicate(t, pred, &predicateExpectation{
		value:        uint64(123),
		result:       predicate.Invalid,
		descMatchers: []string{"length(value) < 3"},
		errMatchers:  []string{"value of type 'uint64' does not have a length"},
	})
}

// ---------------------------------------------------------------------------
// p.Capacity()
// ---------------------------------------------------------------------------

func TestSliceCapacity(t *testing.T) {
	pred := p.Capacity(p.Lt(3))

	validatePredicate(t, pred, &predicateExpectation{
		value:        []int{1, 2},
		result:       predicate.Passed,
		descMatchers: []string{"capacity(value) < 3"},
		errMatchers:  []string{"capacity: 2"},
	})
	validatePredicate(t, pred, &predicateExpectation{
		value:        []int{1, 2, 3, 4, 5},
		result:       predicate.Failed,
		descMatchers: []string{"capacity(value) < 3"},
		errMatchers:  []string{"capacity: 5"},
	})
}

func TestStringCapacity(t *testing.T) {
	pred := p.Capacity(p.Lt(3))

	validatePredicate(t, pred, &predicateExpectation{
		value:        "Yo",
		result:       predicate.Invalid,
		descMatchers: []string{"capacity(value) < 3"},
		errMatchers:  []string{"value of type 'string' does not have a capacity"},
	})
}
