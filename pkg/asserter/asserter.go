package asserter

import (
	"fmt"
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/predicate"
	"github.com/maargenton/go-testpredicate/pkg/prettyprint"
)

// T is a wrapper around a testing interface that can verify predicates
// and display failure reasons on error
type T interface {

	// That verifies that a value matches a predicate, and outputs detail
	// information in case of a failure or error. Additional details can be
	// passed in as a format string and arguments, or just a list of arguments.
	That(v interface{}, p predicate.T, details ...interface{})
}

// New return an implementation of the T interface wrapping a
// testing.T context
func New(t ctx, option ...Option) T {
	var a = &assert{t: t, opts: opts{abortOnError: true}}
	for _, opt := range option {
		opt(&a.opts)
	}
	return a
}

type opts struct {
	abortOnError bool
}

// Option is passed to New() to configure assertion handling
type Option func(*opts)

// AbortOnError tells teh asserter wether or not to fail
func AbortOnError(b bool) Option {
	return func(o *opts) {
		o.abortOnError = b
	}
}

//
// ---------------------------------------------------------------------------
// Implementation of the Asserter interface
// ---------------------------------------------------------------------------

type ctx interface {
	Helper()
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

var _ ctx = &testing.T{}

type assert struct {
	opts
	t ctx
}

func (assert *assert) That(v interface{}, p predicate.T, details ...interface{}) {
	assert.t.Helper()
	r, err := p.Evaluate(v)
	if r.Success() {
		return
	}

	s := ""
	if len(details) != 0 {
		s += "\n" + formatDetails(details...)
	}
	s += fmt.Sprintf("\nexpected: %v,", p)
	if err != nil {
		s += fmt.Sprintf("\n%v,", err)
	}
	s += fmt.Sprintf("\nvalue: %v", prettyprint.FormatValue(v))

	assert.fail(s)
}

func (assert *assert) fail(s string, args ...interface{}) {
	if assert.abortOnError {
		assert.t.Fatalf(s, args...)
	} else {
		assert.t.Errorf(s, args...)
	}
}

// formatDetails formats a list of assertion details either as a format string
// with a list of arguments, or just a list of values
func formatDetails(details ...interface{}) string {
	if len(details) == 0 {
		return ""
	}
	if s, ok := details[0].(string); ok {
		return fmt.Sprintf(s, details[1:]...)
	}
	return fmt.Sprint(details...)
}
