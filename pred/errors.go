package pred

import (
	"fmt"

	"github.com/maargenton/go-testpredicate"
)

// IsNoError tests if an error value is nil
func IsNoError() testpredicate.Predicate {
	return testpredicate.MakePredicate(
		"value is not an error",
		func(value interface{}) (testpredicate.PredicateResult, error) {
			if value == nil {
				return testpredicate.PredicatePassed, nil
			} else if err, ok := value.(error); ok {
				return testpredicate.PredicateFailed,
					fmt.Errorf("detailed error: %v", err)
			} else {
				return testpredicate.PredicateInvalid,
					fmt.Errorf("value of type '%T' is not an error", value)
			}
		})
}
