package predicate_test

import (
	"fmt"
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/internal/predicate"
)

func TestFormatContextValue(t *testing.T) {

	var ctx = []predicate.ContextValue{
		{Name: "expected", Value: "value is nil", Pre: true},
		{Name: "actual", Value: nil, Pre: false},
		{Name: "error", Value: fmt.Errorf("custom error"), Pre: false},
	}

	var s = predicate.FormatContextValues(ctx)
	var expected = "" +
		"expected: value is nil\n" +
		"error:    &errors.errorString{\n" +
		"          \ts: \"custom error\",\n" +
		"          }\n" +
		"actual:   <nil>\n"

	if s != expected {
		t.Errorf("\noutput mismatch\n%v", s)
	}
}
