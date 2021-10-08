// Package subexpr starts an unevaluated predicate chain that is captured an
// outer predicate responsible for evaluating the condition on one or more
// values of a collection.
package subexpr

import "github.com/maargenton/go-testpredicate/pkg/utils/builder"

// Value starts a sub-expression predicate for use as part of collection
// predicate (e.g. .All() or .Any()) that would evaluate it on multiple values
// and aggregate the results.
func Value() *builder.Builder {
	var b = builder.New(nil, nil, false)
	builder.CaptureCallsite(b, 1)
	return b
}
