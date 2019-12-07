package assert_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/assert"
	"github.com/maargenton/go-testpredicate/pkg/predicate"
)

type testingContext struct {
	helperCount int
	errors      []string
}

func (c *testingContext) Helper() {
	c.helperCount++
}

func (c *testingContext) Errorf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	c.errors = append(c.errors, s)
}

func (c *testingContext) Fatalf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	c.errors = append(c.errors, s)
}

func TestTestingAsserter(t *testing.T) {
	ctx := &testingContext{}
	assert := assert.New(ctx)
	p := predicate.Make("description", func(v interface{}) (predicate.Result, error) {
		return predicate.Invalid, fmt.Errorf("unimplemented")
	})
	assert.That(123, p, "context: %v", 456)

	if ctx.helperCount != 1 {
		t.Errorf(
			"Expected assert.That() to call t.Helper() once, was called %v time(s)",
			ctx.helperCount)
	}

	if len(ctx.errors) != 1 {
		t.Errorf(
			"Expected assert.That() to call t.Errorf() once on failure, was called %v time(s)",
			len(ctx.errors))
	}

	var failed = false
	var err = ctx.errors[0]
	var expectedFragments = []string{
		"context: 456",
		"expected: description",
		"\nunimplemented,\n",
		"value: 123",
	}

	for _, frg := range expectedFragments {
		if !strings.Contains(err, frg) {
			failed = true
			t.Errorf("expected error to containe '%v'", frg)
		}
	}

	if failed {
		t.Errorf("error message:\n%v", err)
	}
}
