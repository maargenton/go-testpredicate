// Package verify starts a predicate chain that will let the current test
// continue even if the assertions fails.
package verify

import (
	"github.com/maargenton/go-testpredicate/pkg/internal/builder"
	"github.com/maargenton/go-testpredicate/pkg/internal/predicate"
)

// That captrues the test context and a value for the purpose of building and
// evaluating a test predicate through call chaining. The current test will
// proceed even if the predicate fails.
func That(t predicate.T, v interface{}) *builder.Builder {
	var b = builder.New(t, v, false)
	builder.CaptureCallsite(b, 1)
	t.Cleanup(func() {
		builder.VerifyCompletness(b)
	})
	return b
}
