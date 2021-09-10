// Package require starts a predicate chain that will fail the current test
// if the assertions fails.
package require

import (
	"github.com/maargenton/go-testpredicate/pkg/internal/builder"
	"github.com/maargenton/go-testpredicate/pkg/internal/predicate"
)

// Context captures an additional context value to be displayed upon failure.
type Context = predicate.ContextValue

// That captrues the test context and a value for the purpose of building and
// evaluating a test predicate through call chaining. The current test will
// fail immediately if the predicate fails.
func That(t predicate.T, v interface{}, ctx ...Context) *builder.Builder {
	var b = builder.New(t, v, true)
	builder.CaptureCallsite(b, 1)
	t.Cleanup(func() {
		builder.VerifyCompletness(b)
	})
	b.Ctx = append(b.Ctx, ctx...)
	return b
}
