package pred

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/maargenton/go-testpredicate"
	"github.com/maargenton/go-testpredicate/utils"
)

// All tests if a sub-predicate passes for all elements of an array
// or slice
func All(p testpredicate.Predicate) testpredicate.Predicate {
	return testpredicate.MakePredicate(
		strings.Replace(p.String(), "value", "all of value", -1),

		func(value interface{}) (testpredicate.PredicateResult, error) {

			rv := reflect.ValueOf(value)
			switch rv.Kind() {
			case reflect.Array, reflect.Slice, reflect.String:

				for i, n := 0, rv.Len(); i < n; i++ {
					v := rv.Index(i).Interface()
					r, err := p.Evaluate(v)
					if err != nil || r != testpredicate.PredicatePassed {
						err = utils.WrapError(err,
							"failed for value[%v]: %v",
							i, utils.FormatValue(v))
						return r, err
					}
				}
				return testpredicate.PredicatePassed, nil

			default:
				return testpredicate.PredicateInvalid,
					fmt.Errorf(
						"value %v of type %T is not a container",
						utils.FormatValue(value), value)
			}
		})
}

// Any tests if a sub-predicate passes for any elements of an array
// or slice
func Any(p testpredicate.Predicate) testpredicate.Predicate {
	return testpredicate.MakePredicate(
		strings.Replace(p.String(), "value", "any of value", -1),

		func(value interface{}) (testpredicate.PredicateResult, error) {

			rv := reflect.ValueOf(value)
			switch rv.Kind() {
			case reflect.Array, reflect.Slice, reflect.String:

				for i, n := 0, rv.Len(); i < n; i++ {
					v := rv.Index(i).Interface()
					r, err := p.Evaluate(v)
					if err != nil || r == testpredicate.PredicateInvalid {
						err = utils.WrapError(err,
							"failed for value[%v]: %v",
							i, utils.FormatValue(v))
						return r, err
					}
					if r == testpredicate.PredicatePassed {
						err = utils.WrapError(err,
							"passed for value[%v]: %v",
							i, utils.FormatValue(v))
						return r, err
					}
				}
				return testpredicate.PredicateFailed, nil

			default:
				return testpredicate.PredicateInvalid,
					fmt.Errorf(
						"value %v of type %T is not a container",
						utils.FormatValue(value), value)
			}
		})
}

// AllKeys tests if a sub-predicate passes for all keys of a map
func AllKeys(p testpredicate.Predicate) testpredicate.Predicate {
	return testpredicate.MakePredicate(
		strings.Replace(p.String(), "value", "all of value.Keys()", -1),

		func(value interface{}) (testpredicate.PredicateResult, error) {
			rmap := reflect.ValueOf(value)
			if rmap.Kind() != reflect.Map {
				return testpredicate.PredicateInvalid,
					fmt.Errorf(
						"value %v of type %T is not a map",
						utils.FormatValue(value), value)
			}

			for _, kv := range rmap.MapKeys() {
				k := kv.Interface()
				r, err := p.Evaluate(k)
				if err != nil || r != testpredicate.PredicatePassed {
					err = utils.WrapError(err,
						"failed for key %v", utils.FormatValue(k))
					return r, err
				}
			}
			return testpredicate.PredicatePassed, nil

		})
}

// AnyKey tests if a sub-predicate passes for any key of a map
func AnyKey(p testpredicate.Predicate) testpredicate.Predicate {
	return testpredicate.MakePredicate(
		strings.Replace(p.String(), "value", "any of value.Keys()", -1),

		func(value interface{}) (testpredicate.PredicateResult, error) {
			rmap := reflect.ValueOf(value)
			if rmap.Kind() != reflect.Map {
				return testpredicate.PredicateInvalid,
					fmt.Errorf(
						"value %v of type %T is not a map",
						utils.FormatValue(value), value)
			}

			for _, kv := range rmap.MapKeys() {
				k := kv.Interface()
				r, err := p.Evaluate(k)

				if err != nil || r == testpredicate.PredicateInvalid {
					err = utils.WrapError(err,
						"failed for key: %v",
						utils.FormatValue(k))
					return r, err
				}
				if r == testpredicate.PredicatePassed {
					err = utils.WrapError(err,
						"passed for key: %v",
						utils.FormatValue(k))
					return r, err
				}
			}
			return testpredicate.PredicateFailed, nil

		})
}

// AllValues tests if a sub-predicate passes for all values of a map
func AllValues(p testpredicate.Predicate) testpredicate.Predicate {
	return testpredicate.MakePredicate(
		strings.Replace(p.String(), "value", "all of value.Values()", -1),

		func(value interface{}) (testpredicate.PredicateResult, error) {
			rmap := reflect.ValueOf(value)
			if rmap.Kind() != reflect.Map {
				return testpredicate.PredicateInvalid,
					fmt.Errorf(
						"value %v of type %T is not a map",
						utils.FormatValue(value), value)
			}

			for _, kv := range rmap.MapKeys() {
				k := kv.Interface()
				v := rmap.MapIndex(kv).Interface()
				r, err := p.Evaluate(v)
				if err != nil || r != testpredicate.PredicatePassed {
					err = utils.WrapError(err,
						"failed for value[%v]: %v",
						utils.FormatValue(k), utils.FormatValue(v))
					return r, err
				}
			}
			return testpredicate.PredicatePassed, nil
		})
}

// AnyValue tests if a sub-predicate passes for any value of a map
func AnyValue(p testpredicate.Predicate) testpredicate.Predicate {
	return testpredicate.MakePredicate(
		strings.Replace(p.String(), "value", "any of value.Values()", -1),

		func(value interface{}) (testpredicate.PredicateResult, error) {
			rmap := reflect.ValueOf(value)
			if rmap.Kind() != reflect.Map {
				return testpredicate.PredicateInvalid,
					fmt.Errorf(
						"value %v of type %T is not a map",
						utils.FormatValue(value), value)
			}

			for _, kv := range rmap.MapKeys() {

				k := kv.Interface()
				v := rmap.MapIndex(kv).Interface()

				r, err := p.Evaluate(v)
				if err != nil || r == testpredicate.PredicateInvalid {
					err = utils.WrapError(err,
						"failed for value[%v]: %v",
						utils.FormatValue(k), utils.FormatValue(v))
					return r, err
				}
				if r == testpredicate.PredicatePassed {
					err = utils.WrapError(err,
						"passed for value[%v]: %v",
						utils.FormatValue(k), utils.FormatValue(v))
					return r, err
				}
			}
			return testpredicate.PredicateFailed, nil
		})
}

// MapKeys applies the sub-predicate to the keys of a map
func MapKeys(p testpredicate.Predicate) testpredicate.Predicate {
	return testpredicate.MakePredicate(
		strings.Replace(p.String(), "value", "value.Keys()", -1),

		func(value interface{}) (testpredicate.PredicateResult, error) {

			v := reflect.ValueOf(value)
			if v.Kind() != reflect.Map {
				return testpredicate.PredicateInvalid,
					fmt.Errorf(
						"value %v of type %T is not a map",
						utils.FormatValue(value), value)
			}
			keys := make([]interface{}, 0, v.Len())
			for _, k := range v.MapKeys() {
				keys = append(keys, k.Interface())
			}

			r, err := p.Evaluate(keys)
			// err = utils.WrapError(err, "length: %v", value)
			return r, err
		})
}

// MapValues applies the sub-predicate to the values of a map
func MapValues(p testpredicate.Predicate) testpredicate.Predicate {
	return testpredicate.MakePredicate(
		strings.Replace(p.String(), "value", "value.Values()", -1),

		func(value interface{}) (testpredicate.PredicateResult, error) {

			rmap := reflect.ValueOf(value)
			if rmap.Kind() != reflect.Map {
				return testpredicate.PredicateInvalid,
					fmt.Errorf(
						"value %v of type %T is not a map",
						utils.FormatValue(value), value)
			}
			values := make([]interface{}, 0, rmap.Len())
			for _, k := range rmap.MapKeys() {
				v := rmap.MapIndex(k)
				values = append(values, v.Interface())
			}

			r, err := p.Evaluate(values)
			// err = utils.WrapError(err, "length: %v", value)
			return r, err
		})
}
