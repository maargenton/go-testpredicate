// GENERATE CODE -- DO NOT EDIT

package builder

import (
	"reflect"

	"github.com/maargenton/go-testpredicate/pkg/utils/predicate"
	"github.com/maargenton/go-testpredicate/pkg/utils/predicate/impl"
)

// ---------------------------------------------------------------------------
// From pkg/utils/predicate/impl/collection.go

// All tests if all values of a collection match the given predicate
func (b *Builder) All(p *predicate.Predicate) *predicate.Predicate {
	b.p.RegisterPredicate(impl.All(p))
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// Any tests if at least one values of a collection match the given predicate
func (b *Builder) Any(p *predicate.Predicate) *predicate.Predicate {
	b.p.RegisterPredicate(impl.Any(p))
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// From pkg/utils/predicate/impl/collection.go
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// From pkg/utils/predicate/impl/compare.go

// IsTrue tests if a value is true
func (b *Builder) IsTrue() *predicate.Predicate {
	b.p.RegisterPredicate(impl.IsTrue())
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// IsFalse tests if a value is false
func (b *Builder) IsFalse() *predicate.Predicate {
	b.p.RegisterPredicate(impl.IsFalse())
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// IsNil tests if a value is either a nil literal or a nillable type set to nil
func (b *Builder) IsNil() *predicate.Predicate {
	b.p.RegisterPredicate(impl.IsNil())
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// IsNotNil tests if a value is neither a nil literal nor a nillable type set to
// nil; any value of a non-nillable type is considered not nil.
func (b *Builder) IsNotNil() *predicate.Predicate {
	b.p.RegisterPredicate(impl.IsNotNil())
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// IsEqualTo tests if a value is equatable and equal to the specified value.
func (b *Builder) IsEqualTo(rhs interface{}) *predicate.Predicate {
	b.p.RegisterPredicate(impl.IsEqualTo(rhs))
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// IsNotEqualTo tests if a value is equatable but different from the specified
// value.
func (b *Builder) IsNotEqualTo(rhs interface{}) *predicate.Predicate {
	b.p.RegisterPredicate(impl.IsNotEqualTo(rhs))
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// Eq tests if a value is equatable and equal to the specified value.
func (b *Builder) Eq(rhs interface{}) *predicate.Predicate {
	b.p.RegisterPredicate(impl.Eq(rhs))
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// Ne tests if a value is equatable but different from the specified
// value.
func (b *Builder) Ne(rhs interface{}) *predicate.Predicate {
	b.p.RegisterPredicate(impl.Ne(rhs))
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// From pkg/utils/predicate/impl/compare.go
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// From pkg/utils/predicate/impl/error.go

// IsError tests an error value to be either nil, a specific error according to
// `errors.Is()`, or an error whose message contains a specified string or
// matches a regexp. `.IsError("")` matches any error whose message contains an
// empty string, which is any non-nil error.
func (b *Builder) IsError(expected any) *predicate.Predicate {
	b.p.RegisterPredicate(impl.IsError(expected))
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// AsError tests if a value is an error matching or wrapping the expected error
// (according to go 1.13 error.As()) and returns the unwrapped error for further
// evaluation.
func (b *Builder) AsError(target interface{}) *Builder {
	b.p.RegisterTransformation(impl.AsError(target))
	return b
}

// From pkg/utils/predicate/impl/error.go
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// From pkg/utils/predicate/impl/ext.go

// Is is an extension point allowing for the definition of a custom
// predicate function to evaluate a predicate chain
func (b *Builder) Is(desc string, f predicate.PredicateFunc) *predicate.Predicate {
	b.p.RegisterPredicate(impl.Is(desc, f))
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// Eval is an extension point allowing for the definition of custom
// transformation functions in a predicate chain
func (b *Builder) Eval(desc string, f predicate.TransformFunc) *Builder {
	b.p.RegisterTransformation(impl.Eval(desc, f))
	return b
}

// Passes evaluates a sub-expression predicate against the value.
func (b *Builder) Passes(p *predicate.Predicate) *predicate.Predicate {
	b.p.RegisterPredicate(impl.Passes(p))
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// From pkg/utils/predicate/impl/ext.go
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// From pkg/utils/predicate/impl/impl.go

// From pkg/utils/predicate/impl/impl.go
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// From pkg/utils/predicate/impl/map.go

// MapKeys is a transformation predicate that applies only to map values and
// extract its keys into an sequence for further evaluation. Note that the keys
// Will appear in no particular order.
func (b *Builder) MapKeys() *Builder {
	b.p.RegisterTransformation(impl.MapKeys())
	return b
}

// MapValues is a transformation predicate that applies only to map values and
// extract its values into an sequence for further evaluation. Note that the
// values Will appear in no particular order.
func (b *Builder) MapValues() *Builder {
	b.p.RegisterTransformation(impl.MapValues())
	return b
}

// From pkg/utils/predicate/impl/map.go
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// From pkg/utils/predicate/impl/ordered.go

// IsLessThan tests if a value is strictly less than a reference value
func (b *Builder) IsLessThan(rhs interface{}) *predicate.Predicate {
	b.p.RegisterPredicate(impl.IsLessThan(rhs))
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// IsLessOrEqualTo tests if a value is less than or equal to a reference value
func (b *Builder) IsLessOrEqualTo(rhs interface{}) *predicate.Predicate {
	b.p.RegisterPredicate(impl.IsLessOrEqualTo(rhs))
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// IsGreaterThan tests if a value is strictly greater than a reference value
func (b *Builder) IsGreaterThan(rhs interface{}) *predicate.Predicate {
	b.p.RegisterPredicate(impl.IsGreaterThan(rhs))
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// IsGreaterOrEqualTo tests if a value is greater than or equal to a reference value
func (b *Builder) IsGreaterOrEqualTo(rhs interface{}) *predicate.Predicate {
	b.p.RegisterPredicate(impl.IsGreaterOrEqualTo(rhs))
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// IsCloseTo tests if a value is within tolerance of a reference value
func (b *Builder) IsCloseTo(rhs interface{}, tolerance float64) *predicate.Predicate {
	b.p.RegisterPredicate(impl.IsCloseTo(rhs, tolerance))
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// Lt tests if a value is strictly less than a reference value
func (b *Builder) Lt(rhs interface{}) *predicate.Predicate {
	b.p.RegisterPredicate(impl.Lt(rhs))
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// Le tests if a value is less than or equal to a reference value
func (b *Builder) Le(rhs interface{}) *predicate.Predicate {
	b.p.RegisterPredicate(impl.Le(rhs))
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// Gt tests if a value is strictly greater than a reference value
func (b *Builder) Gt(rhs interface{}) *predicate.Predicate {
	b.p.RegisterPredicate(impl.Gt(rhs))
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// Ge tests if a value is greater than or equal to a reference value
func (b *Builder) Ge(rhs interface{}) *predicate.Predicate {
	b.p.RegisterPredicate(impl.Ge(rhs))
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// From pkg/utils/predicate/impl/ordered.go
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// From pkg/utils/predicate/impl/panic.go

// Panics verifies that the value under test is a callable function that panics.
// Since version > 1.5.0, `panic(nil)` is no longer a special case, reflecting
// the change of behavior in Go 1.21.
func (b *Builder) Panics() *predicate.Predicate {
	b.p.RegisterPredicate(impl.Panics())
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// PanicsAndRecoveredValue verifies that the value under test is a callable
// function that panics, and captures the recovered value for further
// evaluation.
func (b *Builder) PanicsAndRecoveredValue() *Builder {
	b.p.RegisterTransformation(impl.PanicsAndRecoveredValue())
	return b
}

// From pkg/utils/predicate/impl/panic.go
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// From pkg/utils/predicate/impl/sequence.go

// Length is a transformation predicate that extract the length of a value for
// further evaluation. It applies to values of type String, Array, Slice, Map,
// and Channel.
func (b *Builder) Length() *Builder {
	b.p.RegisterTransformation(impl.Length())
	return b
}

// Capacity is a transformation predicate that extract the capacity of a value
// for further evaluation. It applies to values of type  Array, Slice and
// Channel.
func (b *Builder) Capacity() *Builder {
	b.p.RegisterTransformation(impl.Capacity())
	return b
}

// IsEmpty tests if a sequence or container is empty.
func (b *Builder) IsEmpty() *predicate.Predicate {
	b.p.RegisterPredicate(impl.IsEmpty())
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// IsNotEmpty tests if a sequence or container is not empty.
func (b *Builder) IsNotEmpty() *predicate.Predicate {
	b.p.RegisterPredicate(impl.IsNotEmpty())
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// StartsWith tests if a sequence value starts with the given sequence, and can
// be applied to  strings, arrays and slices.
func (b *Builder) StartsWith(rhs interface{}) *predicate.Predicate {
	b.p.RegisterPredicate(impl.StartsWith(rhs))
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// Contains tests if a sequence value contains the given sequence, and can
// be applied to  strings, arrays and slices.
func (b *Builder) Contains(rhs interface{}) *predicate.Predicate {
	b.p.RegisterPredicate(impl.Contains(rhs))
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// EndsWith tests if a sequence value ends with the given sequence, and can
// be applied to  strings, arrays and slices.
func (b *Builder) EndsWith(rhs interface{}) *predicate.Predicate {
	b.p.RegisterPredicate(impl.EndsWith(rhs))
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// HasPrefix tests if a sequence value starts with the given sequence, and can
// be applied to  strings, arrays and slices.
func (b *Builder) HasPrefix(rhs interface{}) *predicate.Predicate {
	b.p.RegisterPredicate(impl.HasPrefix(rhs))
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// HasSuffix tests if a sequence value ends with the given sequence, and can
// be applied to  strings, arrays and slices.
func (b *Builder) HasSuffix(rhs interface{}) *predicate.Predicate {
	b.p.RegisterPredicate(impl.HasSuffix(rhs))
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// From pkg/utils/predicate/impl/sequence.go
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// From pkg/utils/predicate/impl/set.go

// IsEqualSet tests if two containers contain the same set of values,
// independently of order.
func (b *Builder) IsEqualSet(rhs interface{}) *predicate.Predicate {
	b.p.RegisterPredicate(impl.IsEqualSet(rhs))
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// IsDisjointSetFrom tests if two containers contain no common values
func (b *Builder) IsDisjointSetFrom(rhs interface{}) *predicate.Predicate {
	b.p.RegisterPredicate(impl.IsDisjointSetFrom(rhs))
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// IsSubsetOf tests if the value under test is a subset of the reference value.
// Both values must be containers and are treated as unordered sets.
func (b *Builder) IsSubsetOf(rhs interface{}) *predicate.Predicate {
	b.p.RegisterPredicate(impl.IsSubsetOf(rhs))
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// IsSupersetOf tests if the value under test is a superset of the reference
// value. Both values must be containers and are treated as unordered sets.
func (b *Builder) IsSupersetOf(rhs interface{}) *predicate.Predicate {
	b.p.RegisterPredicate(impl.IsSupersetOf(rhs))
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// From pkg/utils/predicate/impl/set.go
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// From pkg/utils/predicate/impl/string.go

// Matches tests if a string matches a regular expression
func (b *Builder) Matches(re string) *predicate.Predicate {
	b.p.RegisterPredicate(impl.Matches(re))
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// ToString is a transformation predicate that converts any value to a string
// representation using `%v` formatting option.
func (b *Builder) ToString() *Builder {
	b.p.RegisterTransformation(impl.ToString())
	return b
}

// ToLower is a transformation predicate that converts any value to a string
// representation using `%v` formatting option.
func (b *Builder) ToLower() *Builder {
	b.p.RegisterTransformation(impl.ToLower())
	return b
}

// ToUpper is a transformation predicate that converts any value to a string
// representation using `%v` formatting option.
func (b *Builder) ToUpper() *Builder {
	b.p.RegisterTransformation(impl.ToUpper())
	return b
}

// From pkg/utils/predicate/impl/string.go
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// From pkg/utils/predicate/impl/struct.go

// Field is a transformation predicate that extract a field from a struct or a
// value from a map, identified by the given `keypath`. See value.Field() for
// more details.
func (b *Builder) Field(keypath string) *Builder {
	b.p.RegisterTransformation(impl.Field(keypath))
	return b
}

// From pkg/utils/predicate/impl/struct.go
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// From pkg/utils/predicate/impl/type.go

// IsA tests if a value is of a given type.
func (b *Builder) IsA(typ reflect.Type) *predicate.Predicate {
	b.p.RegisterPredicate(impl.IsA(typ))
	if b.t != nil {
		b.t.Helper()
		Evaluate(b)
	}
	return &b.p
}

// From pkg/utils/predicate/impl/type.go
// ---------------------------------------------------------------------------
