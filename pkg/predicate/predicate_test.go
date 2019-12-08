package predicate_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/predicate"
)

// ---------------------------------------------------------------------------
// Result
// ---------------------------------------------------------------------------

func TestPredicatePassed(t *testing.T) {
	r := predicate.Passed
	if !r.Valid() {
		t.Errorf("\nexprected PredicatePassed to be valid")
	}
	if !r.Success() {
		t.Errorf("\nexprected PredicatePassed to be a success")
	}
	desc := fmt.Sprintf("%v", r)
	if desc != "Passed" {
		t.Errorf("\nexprected PredicatePassed to print as 'Passed',\nactual: '%v'", desc)
	}
}

func TestPredicateFailed(t *testing.T) {
	r := predicate.Failed
	if !r.Valid() {
		t.Errorf("\nexprected PredicateFailed to be valid")
	}
	if r.Success() {
		t.Errorf("\nexprected PredicateFailed to not be a success")
	}
	desc := fmt.Sprintf("%v", r)
	if desc != "Failed" {
		t.Errorf("\nexprected PredicatePassed to print as 'Passed',\nactual: '%v'", desc)
	}
}

func TestPredicateInvalid(t *testing.T) {
	r := predicate.Invalid
	if r.Valid() {
		t.Errorf("\nexprected PredicateInvalid to not be valid")
	}
	if r.Success() {
		t.Errorf("\nexprected PredicateInvalid to not be a success")
	}
	desc := fmt.Sprintf("%v", r)
	if desc != "Invalid" {
		t.Errorf(
			"\nexprected PredicatePassed to print as 'Passed',\nactual: '%v'",
			desc)
	}
}

func TestPredicateOther(t *testing.T) {
	r := predicate.Result(123)
	if r.Valid() {
		t.Errorf("\nexprected %v to not be valid", r)
	}
	if r.Success() {
		t.Errorf("\nexprected %v to not be a success", r)
	}

	expected := "Result(123)"
	actual := fmt.Sprintf("%v", r)

	if actual != expected {
		t.Errorf("\nexprected %v to print as '%v',\nactual: '%v'",
			r, expected, actual)
	}
}

// ---------------------------------------------------------------------------
// predicate.WrapError()
// ---------------------------------------------------------------------------

func TestWrapError(t *testing.T) {
	err1 := fmt.Errorf("error1")
	err2 := predicate.WrapError(err1, "error2: %v", 123)
	if !strings.HasPrefix(err2.Error(), "error2: 123") {
		t.Errorf("unexpected error: '%v'", err2)

	}
	if !strings.HasSuffix(err2.Error(), "error1") {
		t.Errorf("unexpected error: '%v'", err2)
	}
}

func TestWrapErrorWithNoBaseError(t *testing.T) {
	err2 := predicate.WrapError(nil, "error2: %v", 123)
	if !strings.HasPrefix(err2.Error(), "error2: 123") ||
		strings.Contains(err2.Error(), "\n") {
		t.Errorf("unexpected error: '%v'", err2)
	}
}

// ---------------------------------------------------------------------------
// MakeBool
// ---------------------------------------------------------------------------

func TestMakeBoolWithFalseResult(t *testing.T) {
	p := predicate.MakeBool("description", func(v interface{}) (bool, error) {
		return false, nil
	})

	desc := fmt.Sprintf("%v", p)
	if desc != "description" {
		t.Errorf(
			"\nexpected description when printing predicate,\nactual: '%v'",
			desc)
	}

	r, err := p.Evaluate(0)
	if r != predicate.Failed {
		t.Errorf(
			"\nexpected bool predicate returning false to yield failed result\n"+
				"was: %v", r)
	}
	if err != nil {
		t.Errorf("\nexpected error to be nil\n"+
			"was %v", err)
	}
}

func TestMakeBoolWithTrueResult(t *testing.T) {
	p := predicate.MakeBool("description", func(v interface{}) (bool, error) {
		return true, nil
	})

	desc := fmt.Sprintf("%v", p)
	if desc != "description" {
		t.Errorf(
			"\nexpected description when printing predicate,\nactual: '%v'",
			desc)
	}

	r, err := p.Evaluate(0)
	if r != predicate.Passed {
		t.Errorf(
			"\nexpected bool predicate returning true to yield passed result\n"+
				"was: %v", r)
	}
	if err != nil {
		t.Errorf("\nexpected error to be nil\n"+
			"was %v", err)
	}
}

func TestMakeBoolWithErrorResult(t *testing.T) {
	e := fmt.Errorf("invalid")
	p := predicate.MakeBool("description", func(v interface{}) (bool, error) {
		return true, e
	})

	desc := fmt.Sprintf("%v", p)
	if desc != "description" {
		t.Errorf(
			"\nexpected description when printing predicate,\nactual: '%v'",
			desc)
	}

	r, err := p.Evaluate(0)
	if r != predicate.Invalid {
		t.Errorf(
			"\nexpected bool predicate returning an error to yield invalid result\n"+
				"was: %v", r)
	}
	if err != e {
		t.Errorf("\nexpected error to match error returned from predicate\n"+
			"was %v", err)
	}
}

// ---------------------------------------------------------------------------
// MakeBool
// ---------------------------------------------------------------------------

func TestMakeWithFailedResult(t *testing.T) {
	p := predicate.Make("description", func(v interface{}) (predicate.Result, error) {
		return predicate.Failed, nil
	})

	desc := fmt.Sprintf("%v", p)
	if desc != "description" {
		t.Errorf(
			"\nexpected description when printing predicate,\nactual: '%v'",
			desc)
	}

	r, err := p.Evaluate(0)
	if r != predicate.Failed {
		t.Errorf(
			"\nexpected returning failed to yield failed result\n"+
				"was: %v", r)
	}
	if err != nil {
		t.Errorf("\nexpected error to match error returned from predicate\n"+
			"was %v", err)
	}
}

func TestMakeWithPassedResult(t *testing.T) {
	p := predicate.Make("description", func(v interface{}) (predicate.Result, error) {
		return predicate.Passed, nil
	})

	desc := fmt.Sprintf("%v", p)
	if desc != "description" {
		t.Errorf(
			"\nexpected description when printing predicate,\nactual: '%v'",
			desc)
	}

	r, err := p.Evaluate(0)
	if r != predicate.Passed {
		t.Errorf(
			"\nexpected returning passed to yield passed result\n"+
				"was: %v", r)
	}
	if err != nil {
		t.Errorf("\nexpected error to match error returned from predicate\n"+
			"was %v", err)
	}
}

func TestMakeWithInvalidResult(t *testing.T) {
	e := fmt.Errorf("invalid")
	p := predicate.Make("description", func(v interface{}) (predicate.Result, error) {
		return predicate.Invalid, e
	})

	desc := fmt.Sprintf("%v", p)
	if desc != "description" {
		t.Errorf(
			"\nexpected description when printing predicate,\nactual: '%v'",
			desc)
	}

	r, err := p.Evaluate(0)
	if r != predicate.Invalid {
		t.Errorf(
			"\nexpected returning invalid to yield invalid result\n"+
				"was: %v", r)
	}
	if err != e {
		t.Errorf("\nexpected error to match error returned from predicate\n"+
			"was %v", err)
	}
}
