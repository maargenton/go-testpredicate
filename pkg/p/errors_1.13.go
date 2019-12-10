// +build go1.13

package p

import (
	"errors"
	"fmt"

	"github.com/maargenton/go-testpredicate/pkg/predicate"
	"github.com/maargenton/go-testpredicate/pkg/prettyprint"
)

// IsError tests if an error value matches a specific error,
// using errros.Is() to support go v1.13 error wrapping
func IsError(expectedErr error) predicate.T {
	return predicate.Make(
		fmt.Sprintf("value is an error matching %v", expectedErr),
		func(value interface{}) (predicate.Result, error) {
			if value == nil {
				if expectedErr == nil {
					return predicate.Passed, nil
				}
				return predicate.Failed, nil
			} else if err, ok := value.(error); ok {
				if errors.Is(err, expectedErr) {
					return predicate.Passed, nil
				}
				return predicate.Failed,
					fmt.Errorf("message: %v", prettyprint.FormatValue(err.Error()))

			} else {
				return predicate.Invalid,
					fmt.Errorf("value of type '%T' is not an error", value)
			}
		})
}
