package impl

import (
	"github.com/maargenton/go-testpredicate/pkg/internal/predicate"
	"github.com/maargenton/go-testpredicate/pkg/value"
)

// Field is a transformation predicate that extract a field from a struct or a
// value from a map, identified by the given `keypath`. See value.Field() for
// more details.
func Field(keypath string) (desc string, f predicate.TransformFunc) {
	desc = "{}.Keys()"
	f = func(v interface{}) (r interface{}, ctx []predicate.ContextValue, err error) {
		r = value.Field(v, keypath)
		ctx = []predicate.ContextValue{
			{Name: "field", Value: r},
		}
		return
	}
	return
}
