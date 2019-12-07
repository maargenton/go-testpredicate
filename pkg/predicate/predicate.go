package predicate

import "fmt"

// ---------------------------------------------------------------------------
// Result
// ---------------------------------------------------------------------------

// Result that can be either failed, passed or invalid
type Result int8

const (
	// Invalid is the result of a predicate that could not be evaluated
	// and whom result is undetermined
	Invalid Result = iota

	// Failed is the result of a predicate evaluated successfully
	// yielding a negative result
	Failed

	// Passed is the result of a predicate evaluated successfully
	// yielding a positive result
	Passed
)

// Success is true for valid and positive results
func (r Result) Success() bool {
	return r == Passed
}

// Valid is true for successful evaluation yielding either positive or negative result
func (r Result) Valid() bool {
	return r == Passed || r == Failed
}

func (r Result) String() string {
	switch r {
	case Invalid:
		return "Invalid"
	case Failed:
		return "Failed"
	case Passed:
		return "Passed"
	}
	return fmt.Sprintf("Result(%v)", int8(r))
}

// WrapError appends the nested error after the formated message, on a new line
func WrapError(nestedErr error, format string, a ...interface{}) error {
	if nestedErr != nil {
		msg := fmt.Sprintf(format, a...)
		return fmt.Errorf("%v\n%v", msg, nestedErr)
	}

	return fmt.Errorf(format, a...)
}

// ---------------------------------------------------------------------------
//
// ---------------------------------------------------------------------------

// Func is a the funcPredicate type that implements a predicate
type Func func(value interface{}) (Result, error)

// T defines the minimum interface that predicates must implement
type T interface {
	fmt.Stringer
	Evaluate(v interface{}) (Result, error)
}

//
// ---------------------------------------------------------------------------
// Make for funcPredicate based predicates
// ---------------------------------------------------------------------------

// Make wraps a predicate funcPredicate into a predicate interface
func Make(description string, fn Func) T {
	return funcPredicate{fn: fn, description: description}
}

// MakeBool wraps a predicate funcPredicate returning bool into a predicate
// interface. Any error returned from the funcPredicate is interpreted as an invalid
// evaluation.
func MakeBool(
	description string,
	fn func(value interface{}) (bool, error)) T {

	return Make(description,
		func(value interface{}) (Result, error) {
			success, err := fn(value)
			if err != nil {
				return Invalid, err
			}
			if success {
				return Passed, nil
			}
			return Failed, nil
		},
	)
}

//
// ---------------------------------------------------------------------------
// Implementation of the  interface
// ---------------------------------------------------------------------------

type funcPredicate struct {
	fn          Func
	description string
}

func (p funcPredicate) String() string {
	return p.description
}

func (p funcPredicate) Evaluate(v interface{}) (Result, error) {
	return p.fn(v)
}
