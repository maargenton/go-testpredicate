package value

import (
	"reflect"
	"strings"
)

// Field takes a value or an array of values, navigates through the data tree
// according to a keypath, and returns the targeted values. `keypath` is a
// dot-separated list of keys, each used as either field name in a struct, a key
// in a map, or a niladic method name. Once a lookup is successful on the first
// fragment of the keypath, the evaluation continue recursively with the lookup
// result and the remainder of the keypath. If a key lookup fails, a nil value
// is returned. If a key lookup results in a method invocation that yields
// multiple values, only the first one is captured.
//
// In cases where the value is a map[string]..., and the first keypath fragment
// is not a valid key, all partial keypaths are considered as potential keys.
// For example, if the keypath is "foo.bar.baz", "foo" is considered first, then
// "foo.bar", then "foo.bar.baz". However, if the map contains both "foo" and
// "foo.bar" keys, only "foo" will be accessible with this method.
//
// When the current value is a single object and all the fields along the
// keypath are scalar types, the result is a scalar value. For each array or
// slice type along the path, the result become a slice collecting the result of
// evaluating the sub-path on each individual element. The shape of the result
// is then a N-dimensional array, where N is the number of arrays traversed
// along the path. Arrays are always returns as a type agnostic array
// (`[]interface{}`), even if all the values have a consistent type.
func Field(value interface{}, keypath string) interface{} {
	var keys = strings.Split(keypath, ".")
	var v = reflect.ValueOf(value)
	var rv = field(v, keys)
	if rv.IsValid() && rv.CanInterface() {
		return rv.Interface()
	}
	return nil
}

func isNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr,
		reflect.UnsafePointer, reflect.Interface, reflect.Slice:

		return v.IsNil()
	default:
		return false
	}
}

func field(v reflect.Value, keypath []string) reflect.Value {
	if len(keypath) == 0 || isNil(v) || !v.IsValid() {
		return v
	}

	switch v.Type().Kind() {
	case reflect.Ptr, reflect.Interface:
		return field(v.Elem(), keypath)

	case reflect.Array, reflect.Slice:
		var r = make([]interface{}, v.Len())
		for i := 0; i < v.Len(); i++ {
			var vv = v.Index(i)
			var rv = field(vv, keypath)
			if rv.IsValid() && rv.CanInterface() {
				r[i] = rv.Interface()
			}
		}
		return reflect.ValueOf(r)

	case reflect.Struct:
		var r = extractStructField(v, keypath[0])
		return field(r, keypath[1:])

	case reflect.Map:
		var r, remainingKeypath = extractMapField(v, keypath)
		return field(r, remainingKeypath)

	}
	return reflect.Value{}
}

func extractStructField(v reflect.Value, key string) reflect.Value {
	vt := v.Type()
	if field, ok := vt.FieldByName(key); ok {
		rv := v.FieldByIndex(field.Index)
		if rv.IsValid() && rv.CanInterface() {
			return rv
		}
	}
	if m, ok := vt.MethodByName(key); ok {
		return invokeMethod(v, m)
	}

	if v.CanAddr() {
		var pv = v.Addr()
		var pvt = pv.Type()
		if m, ok := pvt.MethodByName(key); ok {
			return invokeMethod(pv, m)
		}
	}

	return reflect.Value{}
}

func extractMapField(v reflect.Value, keypath []string) (r reflect.Value, remainingKeypath []string) {
	for i := range keypath {
		var key = strings.Join(keypath[:i+1], ".")
		var rv = v.MapIndex(reflect.ValueOf(key))
		if rv.IsValid() && rv.CanInterface() {
			return rv, keypath[i+1:]
		}
	}
	return reflect.Value{}, nil
}

func invokeMethod(v reflect.Value, m reflect.Method) reflect.Value {
	if m.Type.NumIn() == 1 && m.Type.NumOut() >= 1 {
		rvs := m.Func.Call([]reflect.Value{v})
		return rvs[0]
	}
	return reflect.Value{}
}
