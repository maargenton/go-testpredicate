package value

import (
	"reflect"
	"strings"
)

// Field takes a root value or an array of root values, navigates through the
// data tree according to a keypath, and returns the targetted values. `keypath`
// is a dot-separated list of keys, each used as either field name in a struct,
// a key in a map, or a niladic method name. Any error during key evalation
// results in a nil value. If a method invocation yields multiple return values,
// only the first one is captured.
//
// When the root is a single object and all the fields along the keypath are
// scalar types, the result is a scalar value. For each array or slice type
// along the path, the result become a slice collecting the result of
// evaluating the sub-path on each individual element. The shape of the result
// is then a N-dimensional array, where N is the number of arrays traversed
// along the path.
//
// Because the implementation is using the reflect package and is mostly type
// agnostic, the resulting arrays are always of type []interface{}, even if the
// field types are consistent accross values.
func Field(value interface{}, keypath string) interface{} {
	var keys = strings.Split(keypath, ".")
	return field(value, keys)
}

func field(value interface{}, keypath []string) interface{} {
	if len(keypath) == 0 || value == nil {
		return value
	}

	v := reflect.ValueOf(value)
	switch v.Type().Kind() {
	case reflect.Ptr:
		return field(v.Elem().Interface(), keypath)

	case reflect.Array, reflect.Slice:
		r := make([]interface{}, v.Len())
		for i := 0; i < v.Len(); i++ {
			vv := v.Index(i)
			r[i] = field(vv.Interface(), keypath)
		}
		return r

	case reflect.Struct:
		r := extractStructField(v, keypath[0])
		return field(r, keypath[1:])

	case reflect.Map:
		r := extractMapField(v, keypath[0])
		return field(r, keypath[1:])

	}
	return nil
}

func extractStructField(v reflect.Value, key string) (r interface{}) {
	vt := v.Type()
	if field, ok := vt.FieldByName(key); ok {
		rv := v.FieldByIndex(field.Index)
		if rv.IsValid() && rv.CanInterface() {
			return rv.Interface()
		}
	}
	if m, ok := vt.MethodByName(key); ok {
		return invokeMethod(v, m)
	}
	return nil
}

func extractMapField(v reflect.Value, key string) (r interface{}) {
	rv := v.MapIndex(reflect.ValueOf(key))
	if rv.IsValid() && rv.CanInterface() {
		return rv.Interface()
	}
	return nil
}

func invokeMethod(v reflect.Value, m reflect.Method) (r interface{}) {
	if m.Type.NumIn() == 1 && m.Type.NumOut() >= 1 {
		rvs := m.Func.Call([]reflect.Value{v})
		return rvs[0].Interface()
	}
	return nil
}
