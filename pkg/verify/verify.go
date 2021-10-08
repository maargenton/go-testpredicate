// Package verify starts a predicate chain that will let the current test
// continue even if the assertions fails.
package verify

import (
	"github.com/maargenton/go-testpredicate/pkg/utils/builder"
	"github.com/maargenton/go-testpredicate/pkg/utils/predicate"
)

// Context captures an additional context value to be displayed upon failure.
type Context = predicate.ContextValue

// That captrues the test context and a value for the purpose of building and
// evaluating a test predicate through call chaining. The current test will
// proceed even if the predicate fails.
func That(t predicate.T, v interface{}, ctx ...Context) *builder.Builder {
	var b = builder.New(t, v, false)
	builder.CaptureCallsite(b, 1)
	t.Cleanup(func() {
		builder.VerifyCompletness(b)
	})
	b.Ctx = append(b.Ctx, ctx...)
	return b
}
