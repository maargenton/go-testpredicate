package impl_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/internal/predicate/impl"
)

func TestMapKeys(t *testing.T) {
	verifyTransform(t, tr(impl.MapKeys()), expectation{
		value: map[string]string{
			"aaa": "bbb",
		},
		result: []string{"aaa"},
	})
	verifyTransform(t, tr(impl.MapKeys()), expectation{
		value:    123,
		result:   nil,
		errorMsg: "value of type 'int' does not have keys",
	})
}

func TestMapValues(t *testing.T) {
	verifyTransform(t, tr(impl.MapValues()), expectation{
		value: map[string]string{
			"aaa": "bbb",
		},
		result: []string{"bbb"},
	})
	verifyTransform(t, tr(impl.MapValues()), expectation{
		value:    123,
		result:   nil,
		errorMsg: "value of type 'int' does not have values",
	})
}
