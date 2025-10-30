package impl

import (
	"fmt"
	"reflect"

	"github.com/maargenton/go-testpredicate/pkg/utils/predicate"
)

// IsA tests if a value is of a given type.
func IsA(typ reflect.Type) (desc string, f predicate.PredicateFunc) {
	desc = fmt.Sprintf("{} is a %v", typ)

	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		var t = reflect.TypeOf(v)

		if t != nil {
			if typ.Kind() == reflect.Interface {
				if t.Implements(typ) {
					return true, nil, nil
				}
			} else if t == typ {
				return true, nil, nil
			}
		}

		ctx = append(ctx, predicate.ContextValue{
			Name:  "type",
			Value: fmt.Sprintf("%T", v),
		})
		return false, ctx, nil
	}
	return
}
