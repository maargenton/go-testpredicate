package asserter_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/asserter"
	"github.com/maargenton/go-testpredicate/pkg/p"
	"github.com/maargenton/go-testpredicate/pkg/predicate"
)

type testingContext struct {
	helperCount int
	errors      []string
	fatalCount  int
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
	c.fatalCount++
}

func TestAssertThatWithSuccessfulPredicateGenerateNoError(t *testing.T) {
	ctx := &testingContext{}
	assert := asserter.New(ctx)
	assert.That(true, p.IsTrue())
	if l := len(ctx.errors); l != 0 {
		t.Errorf("\nexpected no error got %v", l)
	}
}

func TestAsserterIsConfigurable(t *testing.T) {
	ctx := &testingContext{}
	assert := asserter.New(ctx, asserter.AbortOnError())
	assert.That(true, p.IsFalse())
	if ctx.fatalCount != 1 {
		t.Errorf(
			"\nexpected asserter to use Fatalf(), fatalCount: %v",
			ctx.fatalCount)
	}
}

func TestTestingAsserter(t *testing.T) {
	ctx := &testingContext{}
	assert := asserter.New(ctx)
	p := predicate.Make("description", func(v interface{}) (predicate.Result, error) {
		return predicate.Invalid, fmt.Errorf("unimplemented")
	})
	assert.That(123, p, "context", 456, "aaa", "bbb", "ccc", 123, "last detail")

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
		"aaa: \"bbb\"",
		"ccc: 123",
		"last detail",
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
