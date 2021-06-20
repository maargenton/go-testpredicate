package impl

import (
	"fmt"
	"reflect"

	"github.com/maargenton/go-testpredicate/pkg/internal/predicate"
)

// MapKeys is a transformation predicate that applies only to map values and
// extract its keys into an sequence for further evaluation. Note that the keys
// Will appear in no particular order.
func MapKeys() (desc string, f predicate.TransformFunc) {
	desc = "{}.Keys()"
	f = func(v interface{}) (r interface{}, ctx []predicate.ContextValue, err error) {
		r, err = extractMapKeys(v)
		if err == nil {
			ctx = []predicate.ContextValue{
				{Name: "keys", Value: r},
			}
		}
		return
	}
	return
}

// MapValues is a transformation predicate that applies only to map values and
// extract its values into an sequence for further evaluation. Note that the
// values Will appear in no particular order.
func MapValues() (desc string, f predicate.TransformFunc) {
	desc = "{}.Values()"
	f = func(v interface{}) (r interface{}, ctx []predicate.ContextValue, err error) {
		r, err = extractMapValues(v)
		if err == nil {
			ctx = []predicate.ContextValue{
				{Name: "values", Value: r},
			}
		}
		return
	}
	return
}

// ---------------------------------------------------------------------------
// Helper functions to manipulate maps through the reflect package

// extractMapKeys extracts the keys of a map into a slice of the associated key
// type (not a slice of interface{}).
func extractMapKeys(v interface{}) (r interface{}, err error) {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Map {
		err = fmt.Errorf("value of type '%T' does not have keys", v)
		return
	}

	keyType := vv.Type().Key()
	keys := reflect.MakeSlice(reflect.SliceOf(keyType), 0, vv.Len())
	for _, k := range vv.MapKeys() {
		keys = reflect.Append(keys, k)
	}
	r = keys.Interface()
	return
}

// extractMapValues extracts the values of a map into a slice of the associated
// value type (not a slice of interface{}).
func extractMapValues(v interface{}) (r interface{}, err error) {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Map {
		err = fmt.Errorf("value of type '%T' does not have values", v)
		return
	}
	valueType := vv.Type().Elem()
	values := reflect.MakeSlice(reflect.SliceOf(valueType), 0, vv.Len())
	for it := vv.MapRange(); it.Next(); {
		values = reflect.Append(values, it.Value())
	}
	r = values.Interface()
	return
}

// Helper functions to manipulate maps through the reflect package
// ---------------------------------------------------------------------------
