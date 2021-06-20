package impl_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/internal/predicate/impl"
)

func TestMatches(t *testing.T) {
	verifyPredicate(t, pr(impl.Matches("\\d+")), expectation{value: "123", pass: true})
	verifyPredicate(t, pr(impl.Matches("\\d+")), expectation{value: "abc", pass: false})
	verifyPredicate(t, pr(impl.Matches("\\d+")), expectation{
		value:    123,
		errorMsg: "value of type 'int' cannot be matched against a regexp",
	})
	verifyPredicate(t, pr(impl.Matches("\\d++")), expectation{
		value:    "123",
		errorMsg: "failed to compile regexp:",
	})
}

func TestToString(t *testing.T) {
	verifyTransform(t, tr(impl.ToString()), expectation{
		value:  123,
		result: "123",
	})
}

func TestToLower(t *testing.T) {
	verifyTransform(t, tr(impl.ToLower()), expectation{
		value:  "Abc",
		result: "abc",
	})
	verifyTransform(t, tr(impl.ToLower()), expectation{
		value:    123,
		result:   nil,
		errorMsg: "value of type 'int' cannot be transformed to lowercase",
	})
}

func TestToUpper(t *testing.T) {
	verifyTransform(t, tr(impl.ToUpper()), expectation{
		value:  "Abc",
		result: "ABC",
	})
	verifyTransform(t, tr(impl.ToUpper()), expectation{
		value:    123,
		result:   nil,
		errorMsg: "value of type 'int' cannot be transformed to uppercase",
	})
}
