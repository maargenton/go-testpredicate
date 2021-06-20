package builder_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/internal/builder"
)

func TestEvaluateOnEmptyBuilderDoesNothing(t *testing.T) {
	b := builder.New(nil, nil, false)
	builder.VerifyCompletness(b) // No panic
	builder.Evaluate(b)          // No panic

}

func TestEvaluateWithFailure(t *testing.T) {
	tt := &testContext{}
	b := builder.New(tt, nil, true)
	b.Eq(123) // includes evaluation

	if !tt.Failed {
		t.Errorf("\nexpected required predicate to fail current test")
	}

	output := strings.TrimSpace(tt.Output)
	expectedOutput := "" +
		"expected: value == 123\n" +
		"value:    <nil>"
	if output != expectedOutput {
		t.Errorf("\noutput mismatch:\n%v", output)
	}
}

func TestVerifyCompletness(t *testing.T) {
	tt := &testContext{}
	b := builder.New(tt, nil, true)
	builder.CaptureCallsite(b, 0)
	builder.VerifyCompletness(b)

	output := strings.TrimSpace(tt.Output)
	expectedOutput := "predicate chain does not evaluate anything"
	if !strings.Contains(output, expectedOutput) {
		t.Errorf("\noutput mismatch:\n%v", output)
	}
}

// ---------------------------------------------------------------------------

type testContext struct {
	Output       string
	Failed       bool
	CleanupFuncs []func()
}

func (c *testContext) Helper() {}

func (c *testContext) Errorf(format string, args ...interface{}) {
	c.Output += fmt.Sprintf(format, args...)
}

func (c *testContext) FailNow() {
	c.Failed = true
}

func (c *testContext) Cleanup(f func()) {
	c.CleanupFuncs = append(c.CleanupFuncs, f)
}
