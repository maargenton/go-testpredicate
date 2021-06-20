// package subexpr starts an unevaluated predicate chain that is captured an
// outer predicate responsible for evaluating the condition on one or more
// values of a collection.
package subexpr

import "github.com/maargenton/go-testpredicate/pkg/internal/builder"

func Value() *builder.Builder {
	var b = builder.New(nil, nil, false)
	builder.CaptureCallsite(b, 1)
	return b
}
