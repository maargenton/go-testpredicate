package impl

import (
	"errors"
	"fmt"

	"github.com/maargenton/go-testpredicate/pkg/internal/predicate"
)

// IsError tests a value is an error matching or wrapping the expected error
// (according to go 1.13 error.Is()).
func IsError(expected error) (desc string, f predicate.PredicateFunc) {
	desc = fmt.Sprintf("{} is error '%v'", expected)
	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		if v == nil {
			r = expected == nil
		} else if errValue, ok := v.(error); ok {
			r = errors.Is(errValue, expected)
			ctx = []predicate.ContextValue{
				{Name: "error", Value: errValue.Error(), Pre: true},
			}

		} else {
			err = fmt.Errorf("value of type '%T' is not an error", v)
		}
		return
	}
	return
}
