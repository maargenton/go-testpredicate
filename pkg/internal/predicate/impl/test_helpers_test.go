package impl_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/internal/predicate"
)

type predicateRecord struct {
	desc string
	f    predicate.PredicateFunc
}

type transformRecord struct {
	desc string
	f    predicate.TransformFunc
}

func pr(desc string, f predicate.PredicateFunc) predicateRecord {
	return predicateRecord{
		desc: desc,
		f:    f,
	}
}

func tr(desc string, f predicate.TransformFunc) transformRecord {
	return transformRecord{
		desc: desc,
		f:    f,
	}
}

type expectation struct {
	value    interface{}
	pass     bool
	result   interface{}
	errorMsg string
}

func verifyPredicate(t *testing.T, rec predicateRecord, e expectation) {
	t.Helper()
	t.Run(fmt.Sprintf("with value %v", e.value), func(t *testing.T) {
		t.Helper()

		passed, ctx, err := rec.f(e.value)
		if e.pass && !passed {
			t.Errorf("\npredicate expected to pass")
		}
		if !e.pass && passed {
			t.Errorf("\npredicate expected to fail")
		}

		if err != nil && e.errorMsg == "" {
			t.Errorf("\npredicates returned an unexpected error: %v", err)
		}
		if e.errorMsg != "" {
			if err == nil {
				t.Errorf("\npredicates did not return expected error")
			} else {
				if !compareErrorMessage(err, e.errorMsg) {
					t.Errorf("\npredicates returned a different error than expected: %v", err)
				}
			}
		}
		_ = ctx
		_ = err
	})
}

func verifyTransform(t *testing.T, rec transformRecord, e expectation) {
	t.Helper()
	t.Run(fmt.Sprintf("with value %v", e.value), func(t *testing.T) {
		t.Helper()

		r, ctx, err := rec.f(e.value)
		if !reflect.DeepEqual(r, e.result) {
			t.Errorf("\nresult mismatch\n  expected: %v\n  actual: %v",
				e.result, r)

		}
		if err != nil && e.errorMsg == "" {
			t.Errorf("\npredicates returned an unexpected error: %v", err)
		}
		if e.errorMsg != "" {
			if err == nil {
				t.Errorf("\npredicates did not return expected error")
			} else {
				if !compareErrorMessage(err, e.errorMsg) {
					t.Errorf("\npredicates returned a different error than expected: %v", err)
				}
			}
		}
		_ = ctx
		_ = err
	})
}

func compareErrorMessage(err error, msg string) bool {
	if err == nil {
		return false
	}
	e := normalizeErrorMessage(err.Error())
	msg = normalizeErrorMessage(msg)
	return strings.Contains(e, msg)
}

func normalizeErrorMessage(s string) string {
	s = strings.ToLower(s)
	return strings.Join(
		strings.FieldsFunc(s, func(r rune) bool {
			return !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9'))
		}),
		" ",
	)
}
