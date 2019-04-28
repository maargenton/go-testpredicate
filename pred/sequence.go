package pred

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/maargenton/go-testpredicate"
	"github.com/maargenton/go-testpredicate/utils"
)

// IsEmpty returns a predicate that checks if a value is nil
func IsEmpty() testpredicate.Predicate {
	return testpredicate.MakePredicate(
		"value is empty",

		func(value interface{}) (testpredicate.PredicateResult, error) {
			v := reflect.ValueOf(value)
			switch v.Kind() {
			case reflect.Array, reflect.Slice, reflect.Map,
				reflect.Chan, reflect.String:

				l := v.Len()
				if l == 0 {
					return testpredicate.PredicatePassed, nil
				}
				return testpredicate.PredicateFailed, fmt.Errorf("length: %v", l)

			default:
				return testpredicate.PredicateInvalid,
					fmt.Errorf(
						"value of type %T cannot be tested for emptiness",
						value)
			}
		})
}

// IsNotEmpty returns a predicate that checks if a value is not nil
func IsNotEmpty() testpredicate.Predicate {
	return testpredicate.MakePredicate(
		"value is not empty",

		func(value interface{}) (testpredicate.PredicateResult, error) {
			v := reflect.ValueOf(value)
			switch v.Kind() {
			case reflect.Array, reflect.Slice, reflect.Map,
				reflect.Chan, reflect.String:

				l := v.Len()
				if l != 0 {
					return testpredicate.PredicatePassed, nil
				}
				return testpredicate.PredicateFailed, nil

			default:
				return testpredicate.PredicateInvalid,
					fmt.Errorf(
						"value %v of type %T cannot be tested for emptiness",
						utils.FormatValue(value), value)
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
			utils.FormatValue(v1.Interface()), v1.Interface())
	}
	if !sequenceType(v2.Kind()) {
		return fmt.Errorf("value %v of type %T is not a sequence",
			utils.FormatValue(v2.Interface()), v2.Interface())
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

			eq, _ := utils.CompareUnordered(v1.Interface(), v2.Interface())
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
func StartsWith(rhs interface{}) testpredicate.Predicate {
	return testpredicate.MakePredicate(
		fmt.Sprintf("value starts with %v", utils.FormatValue(rhs)),
		func(value interface{}) (testpredicate.PredicateResult, error) {

			v1 := reflect.ValueOf(value)
			v2 := reflect.ValueOf(rhs)
			if err := preCheckSubsequence(v1, v2); err != nil {
				return testpredicate.PredicateInvalid, err
			}

			l1 := v1.Len()
			l2 := v2.Len()
			if l1 < l2 {
				return testpredicate.PredicateFailed, fmt.Errorf(
					"sequence of length %v is too small to contain subsequence of length %v",
					l1, l2)
			}

			if indexOfSubsequence(v1.Slice(0, l2), v2) == 0 {
				return testpredicate.PredicatePassed, nil
			}
			return testpredicate.PredicateFailed, nil
		})
}

// Contains returns a predicate that checks
func Contains(rhs interface{}) testpredicate.Predicate {
	return testpredicate.MakePredicate(
		fmt.Sprintf("value contains %v", utils.FormatValue(rhs)),
		func(value interface{}) (testpredicate.PredicateResult, error) {

			v1 := reflect.ValueOf(value)
			v2 := reflect.ValueOf(rhs)
			if err := preCheckSubsequence(v1, v2); err != nil {
				return testpredicate.PredicateInvalid, err
			}

			l1 := v1.Len()
			l2 := v2.Len()
			if l1 < l2 {
				return testpredicate.PredicateFailed, fmt.Errorf(
					"sequence of length %v is too small to contain subsequence of length %v",
					l1, l2)
			}

			if indexOfSubsequence(v1, v2) >= 0 {
				return testpredicate.PredicatePassed, nil
			}
			return testpredicate.PredicateFailed, nil
		})
}

// EndsWith returns a predicate that checks
func EndsWith(rhs interface{}) testpredicate.Predicate {
	return testpredicate.MakePredicate(
		fmt.Sprintf("value ends with %v", utils.FormatValue(rhs)),
		func(value interface{}) (testpredicate.PredicateResult, error) {

			v1 := reflect.ValueOf(value)
			v2 := reflect.ValueOf(rhs)
			if err := preCheckSubsequence(v1, v2); err != nil {
				return testpredicate.PredicateInvalid, err
			}

			l1 := v1.Len()
			l2 := v2.Len()
			if l1 < l2 {
				return testpredicate.PredicateFailed, fmt.Errorf(
					"sequence of length %v is too small to contain subsequence of length %v",
					l1, l2)
			}

			if indexOfSubsequence(v1.Slice(l1-l2, l1), v2) == 0 {
				return testpredicate.PredicatePassed, nil
			}
			return testpredicate.PredicateFailed, nil
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
func Length(p testpredicate.Predicate) testpredicate.Predicate {
	return testpredicate.MakePredicate(
		strings.Replace(p.String(), "value", "length(value)", -1),

		func(value interface{}) (testpredicate.PredicateResult, error) {
			v := reflect.ValueOf(value)
			switch v.Kind() {
			case reflect.Array, reflect.Slice, reflect.Map,
				reflect.Chan, reflect.String:

				l := v.Len()
				r, err := p.Evaluate(l)
				err = utils.WrapError(err, "length: %v", l)
				return r, err

			default:
				return testpredicate.PredicateInvalid,
					fmt.Errorf("value of type %T does not have a length", value)
			}
		})
}

// Capacity returns a predicate that checks if the capacity of a value matches
// the nested predicate. Applies to values of tpye Array, Slice or Channel.
func Capacity(p testpredicate.Predicate) testpredicate.Predicate {
	return testpredicate.MakePredicate(
		strings.Replace(p.String(), "value", "capacity(value)", -1),

		func(value interface{}) (testpredicate.PredicateResult, error) {
			v := reflect.ValueOf(value)
			switch v.Kind() {
			case reflect.Array, reflect.Slice, reflect.Chan:

				c := v.Cap()
				r, err := p.Evaluate(c)
				err = utils.WrapError(err, "capacity: %v", c)
				return r, err

			default:
				return testpredicate.PredicateInvalid,
					fmt.Errorf("value of type %T does not have a capacity", value)
			}
		})
}
