package impl

import (
	"github.com/maargenton/go-testpredicate/pkg/internal/predicate"
)

// Is is an extension point allowing for the definition of a custom
// predicate function to evaluate a predicate chain
func Is(desc string, f predicate.PredicateFunc) (string, predicate.PredicateFunc) {
	return desc, f
}

// Eval is an extension point allowing for the definition of custom
// transformation functions in a predicate chain
func Eval(desc string, f predicate.TransformFunc) (string, predicate.TransformFunc) {
	return desc, f
}

// Passes evaluates a sub-expression predicate against the value.
func Passes(p *predicate.Predicate) (desc string, f predicate.PredicateFunc) {
	desc = p.FormatDescription("{}")
	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		r, ctx = p.Evaluate(v)
		return
	}
	return
}
