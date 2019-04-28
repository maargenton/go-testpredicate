package pred_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate"
	"github.com/maargenton/go-testpredicate/pred"
)

// ---------------------------------------------------------------------------
// pred.IsEmpty()
// ---------------------------------------------------------------------------

func TestStringIsEmpty(t *testing.T) {
	p := pred.IsEmpty()

	validateredicate(t, p, &predicateExpectation{
		value:        []int{},
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value is empty"},
	})
	validateredicate(t, p, &predicateExpectation{
		value:       []int{1, 2, 3, 4, 5},
		result:      testpredicate.PredicateFailed,
		errMatchers: []string{"length: 5"},
	})
}

func TestSliceIsEmpty(t *testing.T) {
	p := pred.IsEmpty()

	validateredicate(t, p, &predicateExpectation{
		value:        "",
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value is empty"},
	})
	validateredicate(t, p, &predicateExpectation{
		value:       "Hello world!",
		result:      testpredicate.PredicateFailed,
		errMatchers: []string{"length: 12"},
	})
}

func TestMapIsEmpty(t *testing.T) {
	p := pred.IsEmpty()

	validateredicate(t, p, &predicateExpectation{
		value:        map[string]int{},
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value is empty"},
	})
	validateredicate(t, p, &predicateExpectation{
		value: map[string]int{
			"aaa":    2,
			"bbbccc": 4,
		},
		result:      testpredicate.PredicateFailed,
		errMatchers: []string{"length: 2"},
	})
}

func TestIntIsEmpty(t *testing.T) {
	p := pred.IsEmpty()

	validateredicate(t, p, &predicateExpectation{
		value:        uint64(123),
		result:       testpredicate.PredicateInvalid,
		descMatchers: []string{"value is empty"},
		errMatchers:  []string{"value of type uint64 cannot be tested for emptiness"},
	})
}

// ---------------------------------------------------------------------------
// pred.IsNotEmpty()
// ---------------------------------------------------------------------------

func TestStringIsNotEmpty(t *testing.T) {
	p := pred.IsNotEmpty()

	validateredicate(t, p, &predicateExpectation{
		value:        []int{1, 2, 3, 4, 5},
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value is not empty"},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  []int{},
		result: testpredicate.PredicateFailed,
	})
}

func TestSliceIsNotEmpty(t *testing.T) {
	p := pred.IsNotEmpty()

	validateredicate(t, p, &predicateExpectation{
		value:        "Hello world!",
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value is not empty"},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  "",
		result: testpredicate.PredicateFailed,
	})
}

func TestMapIsNotEmpty(t *testing.T) {
	p := pred.IsNotEmpty()

	validateredicate(t, p, &predicateExpectation{
		value: map[string]int{
			"aaa":    2,
			"bbbccc": 4,
		},
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value is not empty"},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  map[string]int{},
		result: testpredicate.PredicateFailed,
	})
}

func TestIntIsNotEmpty(t *testing.T) {
	p := pred.IsNotEmpty()

	validateredicate(t, p, &predicateExpectation{
		value:        uint64(123),
		result:       testpredicate.PredicateInvalid,
		descMatchers: []string{"value is not empty"},
		errMatchers:  []string{`/value .* of type .* cannot be tested for emptiness/`},
	})
}

// ---------------------------------------------------------------------------
// pred.StartsWith()
// ---------------------------------------------------------------------------
func TestStringStartsWith(t *testing.T) {
	p := pred.StartsWith("Hello")

	validateredicate(t, p, &predicateExpectation{
		value:        "Hello world!",
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value starts with \"Hello\""},
	})
	validateredicate(t, p, &predicateExpectation{
		value:       "Hey",
		result:      testpredicate.PredicateFailed,
		errMatchers: []string{`/sequence .* too small/`},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  "Hey world",
		result: testpredicate.PredicateFailed,
	})
	validateredicate(t, p, &predicateExpectation{
		value:       123,
		result:      testpredicate.PredicateInvalid,
		errMatchers: []string{`/value .* of type .* is not a sequence/`},
	})
}

// ---------------------------------------------------------------------------
// pred.Contains()
// ---------------------------------------------------------------------------

func TestStringContains(t *testing.T) {
	p := pred.Contains("worl")

	validateredicate(t, p, &predicateExpectation{
		value:        "Hello world!",
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value contains \"worl\""},
	})
	validateredicate(t, p, &predicateExpectation{
		value:       "Hey",
		result:      testpredicate.PredicateFailed,
		errMatchers: []string{`/sequence .* too small/`},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  "Hey universe",
		result: testpredicate.PredicateFailed,
	})
	validateredicate(t, p, &predicateExpectation{
		value:       123,
		result:      testpredicate.PredicateInvalid,
		errMatchers: []string{`/value .* of type .* is not a sequence/`},
	})
}

// ---------------------------------------------------------------------------
// pred.StartsWith()
// ---------------------------------------------------------------------------
func TestStringEndsWith(t *testing.T) {
	p := pred.EndsWith("world!")

	validateredicate(t, p, &predicateExpectation{
		value:        "Hello world!",
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"value ends with \"world!\""},
	})
	validateredicate(t, p, &predicateExpectation{
		value:       "Hey",
		result:      testpredicate.PredicateFailed,
		errMatchers: []string{`/sequence .* too small/`},
	})
	validateredicate(t, p, &predicateExpectation{
		value:  "Hey world",
		result: testpredicate.PredicateFailed,
	})
	validateredicate(t, p, &predicateExpectation{
		value:       123,
		result:      testpredicate.PredicateInvalid,
		errMatchers: []string{`/value .* of type .* is not a sequence/`},
	})
}

// ---------------------------------------------------------------------------
// pred.Length()
// ---------------------------------------------------------------------------

func TestSliceLength(t *testing.T) {
	p := pred.Length(pred.Lt(3))

	validateredicate(t, p, &predicateExpectation{
		value:        []int{1, 2},
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"length(value) < 3"},
		errMatchers:  []string{"length: 2"},
	})
	validateredicate(t, p, &predicateExpectation{
		value:        []int{1, 2, 3, 4, 5},
		result:       testpredicate.PredicateFailed,
		descMatchers: []string{"length(value) < 3"},
		errMatchers:  []string{"length: 5"},
	})
}

func TestStringLength(t *testing.T) {
	p := pred.Length(pred.Lt(3))

	validateredicate(t, p, &predicateExpectation{
		value:        "Yo",
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"length(value) < 3"},
		errMatchers:  []string{"length: 2"},
	})
	validateredicate(t, p, &predicateExpectation{
		value:        "Hello world!",
		result:       testpredicate.PredicateFailed,
		descMatchers: []string{"length(value) < 3"},
		errMatchers:  []string{"length: 12"},
	})
}

func TestIntLength(t *testing.T) {
	p := pred.Length(pred.LessThan(3))
	validateredicate(t, p, &predicateExpectation{
		value:        uint64(123),
		result:       testpredicate.PredicateInvalid,
		descMatchers: []string{"length(value) < 3"},
		errMatchers:  []string{"value of type uint64 does not have a length"},
	})
}

// ---------------------------------------------------------------------------
// pred.Capacity()
// ---------------------------------------------------------------------------

func TestSliceCapacity(t *testing.T) {
	p := pred.Capacity(pred.Lt(3))

	validateredicate(t, p, &predicateExpectation{
		value:        []int{1, 2},
		result:       testpredicate.PredicatePassed,
		descMatchers: []string{"capacity(value) < 3"},
		errMatchers:  []string{"capacity: 2"},
	})
	validateredicate(t, p, &predicateExpectation{
		value:        []int{1, 2, 3, 4, 5},
		result:       testpredicate.PredicateFailed,
		descMatchers: []string{"capacity(value) < 3"},
		errMatchers:  []string{"capacity: 5"},
	})
}

func TestStringCapacity(t *testing.T) {
	p := pred.Capacity(pred.Lt(3))

	validateredicate(t, p, &predicateExpectation{
		value:        "Yo",
		result:       testpredicate.PredicateInvalid,
		descMatchers: []string{"capacity(value) < 3"},
		errMatchers:  []string{"value of type string does not have a capacity"},
	})
}
