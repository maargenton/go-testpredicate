package impl_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/utils/predicate/impl"
)

func TestIsTrue(t *testing.T) {
	verifyPredicate(t, pr(impl.IsTrue()), expectation{value: true, pass: true})
	verifyPredicate(t, pr(impl.IsTrue()), expectation{value: false, pass: false})
	verifyPredicate(t, pr(impl.IsTrue()), expectation{
		value:    123,
		pass:     false,
		errorMsg: "value of type 'int' is never true",
	})
}

func TestIsFalse(t *testing.T) {
	verifyPredicate(t, pr(impl.IsFalse()), expectation{value: false, pass: true})
	verifyPredicate(t, pr(impl.IsFalse()), expectation{value: true, pass: false})
	verifyPredicate(t, pr(impl.IsFalse()), expectation{
		value:    123,
		pass:     false,
		errorMsg: "value of type 'int' is never false",
	})
}

func TestIsNil(t *testing.T) {
	var p *int
	verifyPredicate(t, pr(impl.IsNil()), expectation{value: nil, pass: true})
	verifyPredicate(t, pr(impl.IsNil()), expectation{value: p, pass: true})
	verifyPredicate(t, pr(impl.IsNil()), expectation{value: &struct{}{}, pass: false})
	verifyPredicate(t, pr(impl.IsNil()), expectation{
		value:    123,
		pass:     false,
		errorMsg: "value of type 'int' is never nil",
	})
}

func TestIsNotNil(t *testing.T) {
	var p *int
	verifyPredicate(t, pr(impl.IsNotNil()), expectation{value: &struct{}{}, pass: true})
	verifyPredicate(t, pr(impl.IsNotNil()), expectation{value: nil, pass: false})
	verifyPredicate(t, pr(impl.IsNotNil()), expectation{value: p, pass: false})
	verifyPredicate(t, pr(impl.IsNotNil()), expectation{
		value:    123,
		pass:     false,
		errorMsg: "value of type 'int' can never be nil",
	})
}

func TestIsEqualTo(t *testing.T) {
	verifyPredicate(t, pr(impl.IsEqualTo(123)), expectation{value: 123, pass: true})
	verifyPredicate(t, pr(impl.IsEqualTo(123)), expectation{value: 124, pass: false})
	verifyPredicate(t, pr(impl.Eq("123")), expectation{
		value:    123,
		pass:     false,
		errorMsg: "values of type 'int' and 'string' are never equal",
	})
}

func TestIsNotEqualTo(t *testing.T) {
	verifyPredicate(t, pr(impl.IsNotEqualTo(123)), expectation{value: 124, pass: true})
	verifyPredicate(t, pr(impl.IsNotEqualTo(123)), expectation{value: 123, pass: false})
	verifyPredicate(t, pr(impl.Ne("123")), expectation{
		value:    123,
		pass:     false,
		errorMsg: "values of type 'int' and 'string' are never equal",
	})
}
