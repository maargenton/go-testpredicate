package testpredicate

import "fmt"

// ---------------------------------------------------------------------------
// PredicateResult
// ---------------------------------------------------------------------------

// PredicateResult that can be either failed, passed or invalid
type PredicateResult int8

const (
	// PredicateInvalid is the result of a predicate that could not be evaluated
	// and whom result is undetermined
	PredicateInvalid PredicateResult = iota

	// PredicateFailed is the result of a predicate evaluated successfully
	// yielding a negative result
	PredicateFailed

	// PredicatePassed is the result of a predicate evaluated successfully
	// yielding a positive result
	PredicatePassed
)

// Success is true for valid and positive results
func (r PredicateResult) Success() bool {
	return r == PredicatePassed
}

// Valid is true for successful evaluation yielding either positive or negative result
func (r PredicateResult) Valid() bool {
	return r == PredicatePassed || r == PredicateFailed
}

func (r PredicateResult) String() string {
	switch r {
	case PredicateInvalid:
		return "Invalid"
	case PredicateFailed:
		return "Failed"
	case PredicatePassed:
		return "Passed"
	}
	return fmt.Sprintf("PredicateResult(%v)", int8(r))
}

//
// ---------------------------------------------------------------------------
// PredicateFunc
// ---------------------------------------------------------------------------

// PredicateFunc is a the function type that implements a predicate
type PredicateFunc func(value interface{}) (PredicateResult, error)

// PredicateBoolFunc is a the function type that implements a predicate,
// returning a boolean and an error. Any error returned is interpreted as an
// invalid result due to a failure during evaluation. If a valid negative
// result needs an explanation for its failure, use PredicateFunc instead.
type PredicateBoolFunc func(value interface{}) (bool, error)

// Predicate defines the minimum interface that predicates must implement
type Predicate interface {
	fmt.Stringer
	Evaluate(value interface{}) (PredicateResult, error)
}

// SpecialValueFormatingPredicate defines an optional interface that predicates
// can implement to provide a custom value formater
type SpecialValueFormatingPredicate interface {
	FormatPredicateValue(value interface{}) string
}

//
// ---------------------------------------------------------------------------
// MakePredicate for function based predicates
// ---------------------------------------------------------------------------

// MakePredicate wraps a predicate function into a predicate itnerface
func MakePredicate(description string, fn PredicateFunc) Predicate {
	return functionPredicate{fn: fn, description: description}
}

// MakeBoolPredicate wraps a predicate function returning bool into a predicate
// interface. Any error returned from the function is interpreted as an invalid
// evaluation.
func MakeBoolPredicate(description string, fn PredicateBoolFunc) Predicate {
	return MakePredicate(description,
		func(value interface{}) (PredicateResult, error) {
			success, err := fn(value)
			if err != nil {
				return PredicateInvalid, err
			}
			if success {
				return PredicatePassed, nil
			}
			return PredicateFailed, nil
		})
}

type functionPredicate struct {
	fn          PredicateFunc
	description string
}

func (p functionPredicate) String() string {
	return p.description
}

func (p functionPredicate) Evaluate(value interface{}) (PredicateResult, error) {
	return p.fn(value)
}
