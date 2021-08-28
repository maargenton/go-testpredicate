package impl

import (
	"fmt"
	"reflect"

	"github.com/maargenton/go-testpredicate/pkg/internal/predicate"
	"github.com/maargenton/go-testpredicate/pkg/prettyprint"
	"github.com/maargenton/go-testpredicate/pkg/value"
)

// ---------------------------------------------------------------------------
// Transformation predicates on sequences

// Length is a transformation predicate that extract the length of a value for
// further evaluation. It applies to values of type String, Array, Slice, Map,
// and Channel.
func Length() (desc string, f predicate.TransformFunc) {
	desc = "length({})"
	f = func(v interface{}) (r interface{}, ctx []predicate.ContextValue, err error) {
		vv := reflect.ValueOf(v)
		switch vv.Kind() {
		case reflect.Array, reflect.Slice, reflect.Map,
			reflect.Chan, reflect.String:
			l := vv.Len()
			return l, []predicate.ContextValue{
				{Name: "length", Value: l, Pre: true},
			}, nil
		}
		return nil, nil,
			fmt.Errorf("value of type '%T' does not have a length", v)
	}
	return
}

// Capacity is a transformation predicate that extract the capacity of a value
// for further evaluation. It applies to values of type  Array, Slice and
// Channel.
func Capacity() (desc string, f predicate.TransformFunc) {
	desc = "capacity({})"
	f = func(v interface{}) (r interface{}, ctx []predicate.ContextValue, err error) {
		vv := reflect.ValueOf(v)
		switch vv.Kind() {
		case reflect.Array, reflect.Slice, reflect.Chan:
			c := vv.Cap()
			return c, []predicate.ContextValue{
				{Name: "capacity", Value: c, Pre: true},
			}, nil
		}
		return nil, nil,
			fmt.Errorf("value of type '%T' does not have a capacity", v)
	}
	return
}

// Transformation predicates on sequences
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// Evaluation predicates on sequences

// IsEmpty tests if a sequence or container is empty.
func IsEmpty() (desc string, f predicate.PredicateFunc) {
	desc = "{} is empty"
	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		vv := reflect.ValueOf(v)
		switch vv.Kind() {
		case reflect.Array, reflect.Slice, reflect.Map,
			reflect.Chan, reflect.String:
			l := vv.Len()
			return l == 0, []predicate.ContextValue{
				{Name: "length", Value: l, Pre: true},
			}, nil
		}
		return false, nil,
			fmt.Errorf("value of type '%T' cannot be tested for emptiness", v)
	}
	return
}

// IsNotEmpty tests if a sequence or container is not empty.
func IsNotEmpty() (desc string, f predicate.PredicateFunc) {
	desc = "{} is not empty"
	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		vv := reflect.ValueOf(v)
		switch vv.Kind() {
		case reflect.Array, reflect.Slice, reflect.Map,
			reflect.Chan, reflect.String:
			l := vv.Len()
			return l != 0, []predicate.ContextValue{
				{Name: "length", Value: l, Pre: true},
			}, nil
		}
		return false, nil,
			fmt.Errorf("value of type '%T' cannot be tested for emptiness", v)
	}
	return
}

// StartsWith tests if a sequence value starts with the given sequence, and can
// be applied to  strings, arrays and slices.
func StartsWith(rhs interface{}) (desc string, f predicate.PredicateFunc) {
	desc = fmt.Sprintf("{} starts with %v", prettyprint.FormatValue(rhs))
	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		v1, v2 := reflect.ValueOf(v), reflect.ValueOf(rhs)
		if err := value.PreCheckSubsequence(v1, v2); err != nil {
			return false, nil, err
		}
		l1, l2 := v1.Len(), v2.Len()
		if l1 < l2 {
			return false, nil, fmt.Errorf(
				"sequence of length %v is too short to contain a subsequence of length %v",
				l1, l2)
		}
		i := value.IndexOfSubsequence(v1.Slice(0, l2), v2)
		return i == 0, nil, nil
	}
	return
}

// Contains tests if a sequence value contains the given sequence, and can
// be applied to  strings, arrays and slices.
func Contains(rhs interface{}) (desc string, f predicate.PredicateFunc) {
	desc = fmt.Sprintf("{} contains %v", prettyprint.FormatValue(rhs))
	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		v1, v2 := reflect.ValueOf(v), reflect.ValueOf(rhs)
		if err := value.PreCheckSubsequence(v1, v2); err != nil {
			return false, nil, err
		}
		l1, l2 := v1.Len(), v2.Len()
		if l1 < l2 {
			return false, nil, fmt.Errorf(
				"sequence of length %v is too short to contain a subsequence of length %v",
				l1, l2)
		}
		i := value.IndexOfSubsequence(v1, v2)
		return i >= 0, nil, nil
	}
	return
}

// EndsWith tests if a sequence value ends with the given sequence, and can
// be applied to  strings, arrays and slices.
func EndsWith(rhs interface{}) (desc string, f predicate.PredicateFunc) {
	desc = fmt.Sprintf("{} ends with %v", prettyprint.FormatValue(rhs))
	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		v1, v2 := reflect.ValueOf(v), reflect.ValueOf(rhs)
		if err := value.PreCheckSubsequence(v1, v2); err != nil {
			return false, nil, err
		}
		l1, l2 := v1.Len(), v2.Len()
		if l1 < l2 {
			return false, nil, fmt.Errorf(
				"sequence of length %v is too short to contain a subsequence of length %v",
				l1, l2)
		}
		i := value.IndexOfSubsequence(v1.Slice(l1-l2, l1), v2)
		return i == 0, nil, nil
	}
	return
}

// Evaluation predicates on sequences
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// Aliases

// HasPrefix tests if a sequence value starts with the given sequence, and can
// be applied to  strings, arrays and slices.
func HasPrefix(rhs interface{}) (desc string, f predicate.PredicateFunc) {
	return StartsWith(rhs)
}

// HasSuffix tests if a sequence value ends with the given sequence, and can
// be applied to  strings, arrays and slices.
func HasSuffix(rhs interface{}) (desc string, f predicate.PredicateFunc) {
	return EndsWith(rhs)
}

// Aliases
// ---------------------------------------------------------------------------
