package impl_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/internal/predicate/impl"
)

func TestIsLessThan(t *testing.T) {
	verifyPredicate(t, pr(impl.IsLessThan(123)), expectation{value: 122, pass: true})
	verifyPredicate(t, pr(impl.IsLessThan(123)), expectation{value: 123, pass: false})
	verifyPredicate(t, pr(impl.IsLessThan(123)), expectation{value: 124, pass: false})
	verifyPredicate(t, pr(impl.Lt(123)), expectation{
		value:    "123",
		pass:     false,
		errorMsg: "values of type 'string' and 'int' are not order comparable",
	})

	verifyPredicate(t, pr(impl.IsLessThan("Abc")), expectation{value: "Ab", pass: true})
	verifyPredicate(t, pr(impl.IsLessThan("Abc")), expectation{value: "Abc", pass: false})
	verifyPredicate(t, pr(impl.IsLessThan("Abc")), expectation{value: "Abd", pass: false})
	verifyPredicate(t, pr(impl.Lt("Abc")), expectation{
		value:    123,
		pass:     false,
		errorMsg: "values of type 'int' and 'string' are not order comparable",
	})
}

func TestIsLessOrEqualTo(t *testing.T) {
	verifyPredicate(t, pr(impl.IsLessOrEqualTo(123)), expectation{value: 122, pass: true})
	verifyPredicate(t, pr(impl.IsLessOrEqualTo(123)), expectation{value: 123, pass: true})
	verifyPredicate(t, pr(impl.IsLessOrEqualTo(123)), expectation{value: 124, pass: false})
	verifyPredicate(t, pr(impl.Le(123)), expectation{
		value:    "123",
		pass:     false,
		errorMsg: "values of type 'string' and 'int' are not order comparable",
	})

	verifyPredicate(t, pr(impl.IsLessOrEqualTo("Abc")), expectation{value: "Ab", pass: true})
	verifyPredicate(t, pr(impl.IsLessOrEqualTo("Abc")), expectation{value: "Abc", pass: true})
	verifyPredicate(t, pr(impl.IsLessOrEqualTo("Abc")), expectation{value: "Abd", pass: false})
	verifyPredicate(t, pr(impl.Le("Abc")), expectation{
		value:    123,
		pass:     false,
		errorMsg: "values of type 'int' and 'string' are not order comparable",
	})
}

func TestIsGreaterThan(t *testing.T) {
	verifyPredicate(t, pr(impl.IsGreaterThan(123)), expectation{value: 122, pass: false})
	verifyPredicate(t, pr(impl.IsGreaterThan(123)), expectation{value: 123, pass: false})
	verifyPredicate(t, pr(impl.IsGreaterThan(123)), expectation{value: 124, pass: true})
	verifyPredicate(t, pr(impl.Gt(123)), expectation{
		value:    "123",
		pass:     false,
		errorMsg: "values of type 'string' and 'int' are not order comparable",
	})

	verifyPredicate(t, pr(impl.IsGreaterThan("Abc")), expectation{value: "Ab", pass: false})
	verifyPredicate(t, pr(impl.IsGreaterThan("Abc")), expectation{value: "Abc", pass: false})
	verifyPredicate(t, pr(impl.IsGreaterThan("Abc")), expectation{value: "Abd", pass: true})
	verifyPredicate(t, pr(impl.Gt("Abc")), expectation{
		value:    123,
		pass:     false,
		errorMsg: "values of type 'int' and 'string' are not order comparable",
	})
}

func TestIsGreaterOrEqualTo(t *testing.T) {
	verifyPredicate(t, pr(impl.IsGreaterOrEqualTo(123)), expectation{value: 122, pass: false})
	verifyPredicate(t, pr(impl.IsGreaterOrEqualTo(123)), expectation{value: 123, pass: true})
	verifyPredicate(t, pr(impl.IsGreaterOrEqualTo(123)), expectation{value: 124, pass: true})
	verifyPredicate(t, pr(impl.Ge(123)), expectation{
		value:    "123",
		pass:     false,
		errorMsg: "values of type 'string' and 'int' are not order comparable",
	})

	verifyPredicate(t, pr(impl.IsGreaterOrEqualTo("Abc")), expectation{value: "Ab", pass: false})
	verifyPredicate(t, pr(impl.IsGreaterOrEqualTo("Abc")), expectation{value: "Abc", pass: true})
	verifyPredicate(t, pr(impl.IsGreaterOrEqualTo("Abc")), expectation{value: "Abd", pass: true})
	verifyPredicate(t, pr(impl.Ge("Abc")), expectation{
		value:    123,
		pass:     false,
		errorMsg: "values of type 'int' and 'string' are not order comparable",
	})
}

func TestIsCloseTo(t *testing.T) {
	verifyPredicate(t, pr(impl.IsCloseTo(123, 10)), expectation{value: 113, pass: true})
	verifyPredicate(t, pr(impl.IsCloseTo(123, 10)), expectation{value: 133, pass: true})
	verifyPredicate(t, pr(impl.IsCloseTo(123, 10)), expectation{value: 112, pass: false})
	verifyPredicate(t, pr(impl.IsCloseTo(123, 10)), expectation{value: 134, pass: false})
	verifyPredicate(t, pr(impl.IsCloseTo(123, 10)), expectation{
		value:    "123",
		pass:     false,
		errorMsg: " value of type 'string' cannot be converted to float",
	})
}
