package testpredicate

import (
	"fmt"
	"testing"

	"github.com/maargenton/go-testpredicate/utils"
)

// Asserter is a wrapper around a testing interface that can verify predicates
// and display failure reasons on error
type Asserter interface {

	// That verifies that a value matches a predicate, and outputs detail
	// information in case of a failure or error. Additional details can be
	// passed in as a format string and ergument, or just a list of arguments.
	That(value interface{}, predicate Predicate, details ...interface{})
}

// NewAsserter return an implementation of the Asserter interface wrapping a
// testing.T context
func NewAsserter(t testingContext) Asserter {
	return &testingAsserter{t: t}
}

//
// ---------------------------------------------------------------------------
// Implementation of the Asserter interface
// ---------------------------------------------------------------------------

type testingContext interface {
	Helper()
	Errorf(format string, args ...interface{})
}

var _ testingContext = &testing.T{}

type testingAsserter struct {
	t testingContext
}

func (assert *testingAsserter) That(value interface{}, predicate Predicate, details ...interface{}) {
	assert.t.Helper()
	r, err := predicate.Evaluate(value)
	if !r.Success() {
		s := ""
		if len(details) != 0 {
			s += "\n" + utils.FormatDetails(details...)
		}
		s += fmt.Sprintf("\nexpected: %v,", predicate)
		if err != nil {
			s += fmt.Sprintf("\n%v,", err)
		}
		s += fmt.Sprintf("\nvalue: %v", utils.FormatValue(value))

		assert.t.Errorf(s)
	}
}
