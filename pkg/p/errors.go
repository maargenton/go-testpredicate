package p

import (
	"fmt"

	"github.com/maargenton/go-testpredicate/pkg/predicate"
)

// IsNoError tests if an error value is nil
func IsNoError() predicate.T {
	return predicate.Make(
		"value is not an error",
		func(value interface{}) (predicate.Result, error) {
			if value == nil {
				return predicate.Passed, nil
			} else if err, ok := value.(error); ok {
				return predicate.Failed,
					fmt.Errorf("detailed error: %v", err)
			} else {
				return predicate.Invalid,
					fmt.Errorf("value of type '%T' is not an error", value)
			}
		})
}
