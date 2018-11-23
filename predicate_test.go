package testpredicate_test

import (
	"fmt"
	"testing"

	testpredicate "github.com/marcus999/go-testpredicate"
)

// ---------------------------------------------------------------------------
// PredicateResult
// ---------------------------------------------------------------------------

func TestPredicatePassed(t *testing.T) {
	r := testpredicate.PredicatePassed
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
	r := testpredicate.PredicateFailed
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
	r := testpredicate.PredicateInvalid
	if r.Valid() {
		t.Errorf("\nexprected PredicateInvalid to not be valid")
	}
	if r.Success() {
		t.Errorf("\nexprected PredicateInvalid to not be a success")
	}
	desc := fmt.Sprintf("%v", r)
	if desc != "Invalid" {
		t.Errorf("\nexprected PredicatePassed to print as 'Passed',\nactual: '%v'", desc)
	}
}

func TestPredicateOther(t *testing.T) {
	r := testpredicate.PredicateResult(123)
	if r.Valid() {
		t.Errorf("\nexprected %v to not be valid", r)
	}
	if r.Success() {
		t.Errorf("\nexprected %v to not be a success", r)
	}

	expected := "PredicateResult(123)"
	actual := fmt.Sprintf("%v", r)

	if actual != expected {
		t.Errorf("\nexprected %v to print as '%v',\nactual: '%v'",
			r, expected, actual)
	}
}

// ---------------------------------------------------------------------------
// MakeBoolPredicate
// ---------------------------------------------------------------------------

func TestMakeBoolPredicateFalse(t *testing.T) {
	p := testpredicate.MakeBoolPredicate("description", func(value interface{}) (bool, error) {
		return false, nil
	})

	desc := fmt.Sprintf("%v", p)
	if desc != "description" {
		t.Errorf("\nexpected description when printing predicate,\nactual: '%v'", desc)
	}

	r, _ := p.Evaluate(0)
	if !r.Valid() {
		t.Errorf("\nexpected bool predicate returning false to be valid")
	}
	if r.Success() {
		t.Errorf("\nexpected bool predicate returning false to fail")
	}
}

func TestMakeBoolPredicateTrue(t *testing.T) {
	p := testpredicate.MakeBoolPredicate("description", func(value interface{}) (bool, error) {
		return true, nil
	})

	desc := fmt.Sprintf("%v", p)
	if desc != "description" {
		t.Errorf("\nexpected description when printing predicate,\nactual: '%v'", desc)
	}

	r, _ := p.Evaluate(0)
	if !r.Valid() {
		t.Errorf("\nexpected bool predicate returning true to be valid")
	}
	if !r.Success() {
		t.Errorf("\nexpected bool predicate returning true to succeed")
	}
}

func TestMakeBoolPredicateInvalid(t *testing.T) {
	e := fmt.Errorf("invalid")
	p := testpredicate.MakeBoolPredicate("description", func(value interface{}) (bool, error) {
		return true, e
	})

	desc := fmt.Sprintf("%v", p)
	if desc != "description" {
		t.Errorf("\nexpected description when printing predicate,\nactual: '%v'", desc)
	}

	r, err := p.Evaluate(0)
	if r.Valid() {
		t.Errorf("\nexpected bool predicate returning and error to be invalid")
	}
	if r.Success() {
		t.Errorf("\nexpected bool predicate returning and error to fail")
	}
	if err != e {
		t.Errorf("\nexpected error to match error returned from predicate")
	}
}

// ---------------------------------------------------------------------------
// MakeUnimplemented
// ---------------------------------------------------------------------------

// func TestMakeUnimplemented(t *testing.T) {
// 	e := fmt.Errorf("unimplemented")
// 	p := testpredicate.MakeUnimplemented()

// 	desc := fmt.Sprintf("%v", p)
// 	if desc != "unimplemented" {
// 		t.Errorf(
// 			"\nexpected unimplemented predicate to print as 'unimplemented',\n"+
// 				"actual: '%v'", desc)
// 	}

// 	r, err := p.Evaluate(0)
// 	if r.Valid() {
// 		t.Errorf("\nexpected unimplemented predicate to be invalid")
// 	}
// 	if r.Success() {
// 		t.Errorf("\nexpected unimplemented predicate to fail")
// 	}
// 	if err.Error() != e.Error() {
// 		t.Errorf("\nexpected error from unimplemented predicateto to be '%v',\nactual: %v",
// 			e, err)
// 	}

// }
