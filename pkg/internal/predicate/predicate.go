package predicate

import "strings"

// Transformation captures one transformation step in the predicate evaluation
// chain, with a `Description` and an actual transformation function `Func`.
type Transformation struct {
	Description string
	Func        TransformFunc
}

// TransformFunc is the function type for use in a `Transformation`.
type TransformFunc func(
	value interface{}) (
	result interface{}, ctx []ContextValue, err error)

// Predicate captures a complete predicate chain with `Transformations`, a
// `Description` and an actual evaluation function `Func`.
type Predicate struct {
	Transformations []Transformation
	Description     string
	Func            PredicateFunc
}

// PredicateFunc is the function type for use in a `Predicate`.
type PredicateFunc func(
	value interface{}) (
	success bool, ctx []ContextValue, err error)

// FormatDescription return a formatted description of the full predicate chain,
// using the `value` string to represent the input value.
func (p *Predicate) FormatDescription(value string) string {
	var s = p.Description
	for i := len(p.Transformations) - 1; i >= 0; i-- {
		tr := p.Transformations[i]
		s = strings.Replace(s, "{}", tr.Description, -1)
	}
	return strings.Replace(s, "{}", value, -1)
}

// Evaluate evaluates the full predicate chain on the given `value`, and returns
// a `success` flag and, upon failure, a `context` containing all the relevant
// values captured during evaluation.
func (p *Predicate) Evaluate(value interface{}) (success bool, context []ContextValue) {

	context = []ContextValue{
		{"expected", p.FormatDescription("value"), true},
		{"value", value, false},
	}

	for _, tr := range p.Transformations {
		r, ctx, err := tr.Func(value)
		context = append(context, ctx...)

		if err != nil {
			context = append(context, ContextValue{"error", err, true})
			return
		}
		value = r
	}

	success, ctx, err := p.Func(value)
	context = append(context, ctx...)
	if err != nil {
		context = append(context, ContextValue{"error", err, true})
		success = false
	}
	return
}

// RegisterTransformation appends the given transformation to the list of
// transformations attached to the predicate.
func (p *Predicate) RegisterTransformation(desc string, f TransformFunc) {
	p.Transformations = append(p.Transformations, Transformation{
		Description: desc,
		Func:        f,
	})
}

// RegisterPredicate sets the predicate evaluation function and description for
// the current predicate.
func (p *Predicate) RegisterPredicate(desc string, f PredicateFunc) {
	if p.Func != nil {
		panic("RegisterPredicate() should only be called once per predicate")
	}
	p.Description = desc
	p.Func = f
}
