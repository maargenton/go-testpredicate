package p

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/maargenton/go-testpredicate/pkg/predicate"
	"github.com/maargenton/go-testpredicate/pkg/prettyprint"
	"github.com/maargenton/go-testpredicate/pkg/value"
)

// IsEmpty returns a predicate that checks if a value is nil
func IsEmpty() predicate.T {
	return predicate.Make(
		"value is empty",

		func(value interface{}) (predicate.Result, error) {
			v := reflect.ValueOf(value)
			switch v.Kind() {
			case reflect.Array, reflect.Slice, reflect.Map,
				reflect.Chan, reflect.String:

				l := v.Len()
				if l == 0 {
					return predicate.Passed, nil
				}
				return predicate.Failed, fmt.Errorf("length: %v", l)

			default:
				return predicate.Invalid,
					fmt.Errorf(
						"value of type %T cannot be tested for emptiness",
						value)
			}
		})
}

// IsNotEmpty returns a predicate that checks if a value is not nil
func IsNotEmpty() predicate.T {
	return predicate.Make(
		"value is not empty",

		func(value interface{}) (predicate.Result, error) {
			v := reflect.ValueOf(value)
			switch v.Kind() {
			case reflect.Array, reflect.Slice, reflect.Map,
				reflect.Chan, reflect.String:

				l := v.Len()
				if l != 0 {
					return predicate.Passed, nil
				}
				return predicate.Failed, nil

			default:
				return predicate.Invalid,
					fmt.Errorf(
						"value %v of type %T cannot be tested for emptiness",
						prettyprint.FormatValue(value), value)
			}
		})
}

//
//
// ---------------------------------------------------------------------------
// Sub-sequence predicates
// ---------------------------------------------------------------------------

func sequenceType(k reflect.Kind) bool {
	return k == reflect.Array || k == reflect.Slice || k == reflect.String
}

func preCheckSubsequence(v1, v2 reflect.Value) error {

	if !sequenceType(v1.Kind()) {
		return fmt.Errorf("value %v of type %T is not a sequence",
			prettyprint.FormatValue(v1.Interface()), v1.Interface())
	}
	if !sequenceType(v2.Kind()) {
		return fmt.Errorf("value %v of type %T is not a sequence",
			prettyprint.FormatValue(v2.Interface()), v2.Interface())
	}
	return nil
}

func indexOfSubsequence(seq, sub reflect.Value) int {
	l1 := seq.Len()
	l2 := sub.Len()
	if l1 < l2 {
		return -1
	}

	for i := 0; i <= l1-l2; i++ {
		allEq := true
		for j := 0; j < l2; j++ {
			v1 := seq.Index(i + j)
			v2 := sub.Index(j)

			eq, _ := value.CompareUnordered(v1.Interface(), v2.Interface())
			if !eq {
				allEq = false
				break
			}
		}
		if allEq {
			return i
		}
	}

	return -1
}

// StartsWith returns a predicate that checks
func StartsWith(rhs interface{}) predicate.T {
	return predicate.Make(
		fmt.Sprintf("value starts with %v", prettyprint.FormatValue(rhs)),
		func(value interface{}) (predicate.Result, error) {

			v1 := reflect.ValueOf(value)
			v2 := reflect.ValueOf(rhs)
			if err := preCheckSubsequence(v1, v2); err != nil {
				return predicate.Invalid, err
			}

			l1 := v1.Len()
			l2 := v2.Len()
			if l1 < l2 {
				return predicate.Failed, fmt.Errorf(
					"sequence of length %v is too small to contain subsequence of length %v",
					l1, l2)
			}

			if indexOfSubsequence(v1.Slice(0, l2), v2) == 0 {
				return predicate.Passed, nil
			}
			return predicate.Failed, nil
		})
}

// Contains returns a predicate that checks
func Contains(rhs interface{}) predicate.T {
	return predicate.Make(
		fmt.Sprintf("value contains %v", prettyprint.FormatValue(rhs)),
		func(value interface{}) (predicate.Result, error) {

			v1 := reflect.ValueOf(value)
			v2 := reflect.ValueOf(rhs)
			if err := preCheckSubsequence(v1, v2); err != nil {
				return predicate.Invalid, err
			}

			l1 := v1.Len()
			l2 := v2.Len()
			if l1 < l2 {
				return predicate.Failed, fmt.Errorf(
					"sequence of length %v is too small to contain subsequence of length %v",
					l1, l2)
			}

			if indexOfSubsequence(v1, v2) >= 0 {
				return predicate.Passed, nil
			}
			return predicate.Failed, nil
		})
}

// EndsWith returns a predicate that checks
func EndsWith(rhs interface{}) predicate.T {
	return predicate.Make(
		fmt.Sprintf("value ends with %v", prettyprint.FormatValue(rhs)),
		func(value interface{}) (predicate.Result, error) {

			v1 := reflect.ValueOf(value)
			v2 := reflect.ValueOf(rhs)
			if err := preCheckSubsequence(v1, v2); err != nil {
				return predicate.Invalid, err
			}

			l1 := v1.Len()
			l2 := v2.Len()
			if l1 < l2 {
				return predicate.Failed, fmt.Errorf(
					"sequence of length %v is too small to contain subsequence of length %v",
					l1, l2)
			}

			if indexOfSubsequence(v1.Slice(l1-l2, l1), v2) == 0 {
				return predicate.Passed, nil
			}
			return predicate.Failed, nil
		})
}

//
//
// ---------------------------------------------------------------------------
// Length and capacity predicate modifiers
// ---------------------------------------------------------------------------

// Length returns a predicate that checks if the length of a value matches
// the nested predicate. Applies to values of tpye String, Array, Slice, Map,
// or Channel.
func Length(p predicate.T) predicate.T {
	return predicate.Make(
		strings.Replace(p.String(), "value", "length(value)", -1),

		func(value interface{}) (predicate.Result, error) {
			v := reflect.ValueOf(value)
			switch v.Kind() {
			case reflect.Array, reflect.Slice, reflect.Map,
				reflect.Chan, reflect.String:

				l := v.Len()
				r, err := p.Evaluate(l)
				err = predicate.WrapError(err, "length: %v", l)
				return r, err

			default:
				return predicate.Invalid,
					fmt.Errorf("value of type %T does not have a length", value)
			}
		})
}

// Capacity returns a predicate that checks if the capacity of a value matches
// the nested predicate. Applies to values of tpye Array, Slice or Channel.
func Capacity(p predicate.T) predicate.T {
	return predicate.Make(
		strings.Replace(p.String(), "value", "capacity(value)", -1),

		func(value interface{}) (predicate.Result, error) {
			v := reflect.ValueOf(value)
			switch v.Kind() {
			case reflect.Array, reflect.Slice, reflect.Chan:

				c := v.Cap()
				r, err := p.Evaluate(c)
				err = predicate.WrapError(err, "capacity: %v", c)
				return r, err

			default:
				return predicate.Invalid,
					fmt.Errorf("value of type %T does not have a capacity", value)
			}
		})
}
