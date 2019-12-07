package p

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/maargenton/go-testpredicate/pkg/predicate"
	"github.com/maargenton/go-testpredicate/pkg/prettyprint"
)

// All tests if a sub-predicate passes for all elements of an array
// or slice
func All(p predicate.T) predicate.T {
	return predicate.Make(
		strings.Replace(p.String(), "value", "all of value", -1),

		func(value interface{}) (predicate.Result, error) {

			rv := reflect.ValueOf(value)
			switch rv.Kind() {
			case reflect.Array, reflect.Slice, reflect.String:

				for i, n := 0, rv.Len(); i < n; i++ {
					v := rv.Index(i).Interface()
					r, err := p.Evaluate(v)
					if r != predicate.Passed {
						err = predicate.WrapError(err,
							"failed for value[%v]: %v",
							i, prettyprint.FormatValue(v))
						return r, err
					}
				}
				return predicate.Passed, nil

			default:
				return predicate.Invalid,
					fmt.Errorf(
						"value %v of type %T is not a container",
						prettyprint.FormatValue(value), value)
			}
		})
}

// Any tests if a sub-predicate passes for any elements of an array
// or slice
func Any(p predicate.T) predicate.T {
	return predicate.Make(
		strings.Replace(p.String(), "value", "any of value", -1),

		func(value interface{}) (predicate.Result, error) {

			rv := reflect.ValueOf(value)
			switch rv.Kind() {
			case reflect.Array, reflect.Slice, reflect.String:

				for i, n := 0, rv.Len(); i < n; i++ {
					v := rv.Index(i).Interface()
					r, err := p.Evaluate(v)
					if r == predicate.Invalid {
						err = predicate.WrapError(err,
							"failed for value[%v]: %v",
							i, prettyprint.FormatValue(v))
						return r, err
					}
					if r == predicate.Passed {
						err = predicate.WrapError(err,
							"passed for value[%v]: %v",
							i, prettyprint.FormatValue(v))
						return r, err
					}
				}
				return predicate.Failed, nil

			default:
				return predicate.Invalid,
					fmt.Errorf(
						"value %v of type %T is not a container",
						prettyprint.FormatValue(value), value)
			}
		})
}

// AllKeys tests if a sub-predicate passes for all keys of a map
func AllKeys(p predicate.T) predicate.T {
	return predicate.Make(
		strings.Replace(p.String(), "value", "all of value.Keys()", -1),

		func(value interface{}) (predicate.Result, error) {
			rmap := reflect.ValueOf(value)
			if rmap.Kind() != reflect.Map {
				return predicate.Invalid,
					fmt.Errorf(
						"value %v of type %T is not a map",
						prettyprint.FormatValue(value), value)
			}

			for _, kv := range rmap.MapKeys() {
				k := kv.Interface()
				r, err := p.Evaluate(k)
				if r != predicate.Passed {
					err = predicate.WrapError(err,
						"failed for key %v", prettyprint.FormatValue(k))
					return r, err
				}
			}
			return predicate.Passed, nil

		})
}

// AnyKey tests if a sub-predicate passes for any key of a map
func AnyKey(p predicate.T) predicate.T {
	return predicate.Make(
		strings.Replace(p.String(), "value", "any of value.Keys()", -1),

		func(value interface{}) (predicate.Result, error) {
			rmap := reflect.ValueOf(value)
			if rmap.Kind() != reflect.Map {
				return predicate.Invalid,
					fmt.Errorf(
						"value %v of type %T is not a map",
						prettyprint.FormatValue(value), value)
			}

			for _, kv := range rmap.MapKeys() {
				k := kv.Interface()
				r, err := p.Evaluate(k)

				if r == predicate.Invalid {
					err = predicate.WrapError(err,
						"failed for key: %v",
						prettyprint.FormatValue(k))
					return r, err
				}
				if r == predicate.Passed {
					err = predicate.WrapError(err,
						"passed for key: %v",
						prettyprint.FormatValue(k))
					return r, err
				}
			}
			return predicate.Failed, nil

		})
}

// AllValues tests if a sub-predicate passes for all values of a map
func AllValues(p predicate.T) predicate.T {
	return predicate.Make(
		strings.Replace(p.String(), "value", "all of value.Values()", -1),

		func(value interface{}) (predicate.Result, error) {
			rmap := reflect.ValueOf(value)
			if rmap.Kind() != reflect.Map {
				return predicate.Invalid,
					fmt.Errorf(
						"value %v of type %T is not a map",
						prettyprint.FormatValue(value), value)
			}

			for _, kv := range rmap.MapKeys() {
				k := kv.Interface()
				v := rmap.MapIndex(kv).Interface()
				r, err := p.Evaluate(v)
				if r != predicate.Passed {
					err = predicate.WrapError(err,
						"failed for value[%v]: %v",
						prettyprint.FormatValue(k), prettyprint.FormatValue(v))
					return r, err
				}
			}
			return predicate.Passed, nil
		})
}

// AnyValue tests if a sub-predicate passes for any value of a map
func AnyValue(p predicate.T) predicate.T {
	return predicate.Make(
		strings.Replace(p.String(), "value", "any of value.Values()", -1),

		func(value interface{}) (predicate.Result, error) {
			rmap := reflect.ValueOf(value)
			if rmap.Kind() != reflect.Map {
				return predicate.Invalid,
					fmt.Errorf(
						"value %v of type %T is not a map",
						prettyprint.FormatValue(value), value)
			}

			for _, kv := range rmap.MapKeys() {

				k := kv.Interface()
				v := rmap.MapIndex(kv).Interface()

				r, err := p.Evaluate(v)
				if r == predicate.Invalid {
					err = predicate.WrapError(err,
						"failed for value[%v]: %v",
						prettyprint.FormatValue(k), prettyprint.FormatValue(v))
					return r, err
				}
				if r == predicate.Passed {
					err = predicate.WrapError(err,
						"passed for value[%v]: %v",
						prettyprint.FormatValue(k), prettyprint.FormatValue(v))
					return r, err
				}
			}
			return predicate.Failed, nil
		})
}

// MapKeys applies the sub-predicate to the keys of a map
func MapKeys(p predicate.T) predicate.T {
	return predicate.Make(
		strings.Replace(p.String(), "value", "value.Keys()", -1),

		func(value interface{}) (predicate.Result, error) {

			v := reflect.ValueOf(value)
			if v.Kind() != reflect.Map {
				return predicate.Invalid,
					fmt.Errorf(
						"value %v of type %T is not a map",
						prettyprint.FormatValue(value), value)
			}
			keys := make([]interface{}, 0, v.Len())
			for _, k := range v.MapKeys() {
				keys = append(keys, k.Interface())
			}

			r, err := p.Evaluate(keys)
			// err = predicate.WrapError(err, "length: %v", value)
			return r, err
		})
}

// MapValues applies the sub-predicate to the values of a map
func MapValues(p predicate.T) predicate.T {
	return predicate.Make(
		strings.Replace(p.String(), "value", "value.Values()", -1),

		func(value interface{}) (predicate.Result, error) {

			rmap := reflect.ValueOf(value)
			if rmap.Kind() != reflect.Map {
				return predicate.Invalid,
					fmt.Errorf(
						"value %v of type %T is not a map",
						prettyprint.FormatValue(value), value)
			}
			values := make([]interface{}, 0, rmap.Len())
			for _, k := range rmap.MapKeys() {
				v := rmap.MapIndex(k)
				values = append(values, v.Interface())
			}

			r, err := p.Evaluate(values)
			// err = predicate.WrapError(err, "length: %v", value)
			return r, err
		})
}
