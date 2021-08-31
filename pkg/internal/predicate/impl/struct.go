package impl

import (
	"fmt"

	"github.com/maargenton/go-testpredicate/pkg/internal/predicate"
	"github.com/maargenton/go-testpredicate/pkg/value"
)

// Field is a transformation predicate that extract a field from a struct or a
// value from a map, identified by the given `keypath`. See value.Field() for
// more details.
func Field(keypath string) (desc string, f predicate.TransformFunc) {
	desc = fmt.Sprintf("{}.%v", keypath)
	f = func(v interface{}) (r interface{}, ctx []predicate.ContextValue, err error) {
		r = value.Field(v, keypath)
		ctx = []predicate.ContextValue{
			{Name: fmt.Sprintf("$.%v", keypath), Value: r},
		}
		return
	}
	return
}
