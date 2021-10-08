package impl_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/utils/predicate/impl"
)

func TestLength(t *testing.T) {
	verifyTransform(t, tr(impl.Length()), expectation{
		value:  "Abc",
		result: 3,
	})
	verifyTransform(t, tr(impl.Length()), expectation{
		value:  []int{1, 2, 3, 4, 5},
		result: 5,
	})
	verifyTransform(t, tr(impl.Length()), expectation{
		value:  make([]int, 3, 5),
		result: 3,
	})
	verifyTransform(t, tr(impl.Length()), expectation{
		value:    123,
		result:   nil,
		errorMsg: "value of type 'int' does not have a length",
	})
}

func TestCapacity(t *testing.T) {
	verifyTransform(t, tr(impl.Capacity()), expectation{
		value:    "Abc",
		result:   nil,
		errorMsg: "value of type 'string' does not have a capacity",
	})
	verifyTransform(t, tr(impl.Capacity()), expectation{
		value:  make([]int, 3, 5),
		result: 5,
	})
	verifyTransform(t, tr(impl.Capacity()), expectation{
		value:    123,
		result:   nil,
		errorMsg: "value of type 'int' does not have a capacity",
	})
}

func TestIsEmpty(t *testing.T) {
	verifyPredicate(t, pr(impl.IsEmpty()), expectation{value: "", pass: true})
	verifyPredicate(t, pr(impl.IsEmpty()), expectation{value: "abc", pass: false})
	verifyPredicate(t, pr(impl.IsEmpty()), expectation{value: []int{1, 2, 3}, pass: false})
	verifyPredicate(t, pr(impl.IsEmpty()), expectation{
		value:    123,
		pass:     false,
		errorMsg: "value of type 'int' cannot be tested for emptiness",
	})
}

func TestIsNotEmpty(t *testing.T) {
	verifyPredicate(t, pr(impl.IsNotEmpty()), expectation{value: "abc", pass: true})
	verifyPredicate(t, pr(impl.IsNotEmpty()), expectation{value: []int{1, 2, 3}, pass: true})
	verifyPredicate(t, pr(impl.IsNotEmpty()), expectation{value: "", pass: false})
	verifyPredicate(t, pr(impl.IsNotEmpty()), expectation{
		value:    123,
		errorMsg: "value of type 'int' cannot be tested for emptiness",
	})
}

func TestStartsWith(t *testing.T) {
	verifyPredicate(t, pr(impl.StartsWith("abc")), expectation{value: "abcdef", pass: true})
	verifyPredicate(t, pr(impl.StartsWith("abc")), expectation{value: "azcdef", pass: false})
	verifyPredicate(t, pr(impl.StartsWith([]int{1, 2, 3})), expectation{value: []int{1, 2, 3, 4, 5}, pass: true})
	verifyPredicate(t, pr(impl.StartsWith([]int{1, 2, 3})), expectation{value: []int{1, 6, 3, 4, 5}, pass: false})
	verifyPredicate(t, pr(impl.HasPrefix("abc")), expectation{
		value:    123,
		errorMsg: "value of type 'int' is not a sequence",
	})
	verifyPredicate(t, pr(impl.StartsWith(123)), expectation{
		value:    "abc",
		errorMsg: "value of type 'int' is not a sequence",
	})
	verifyPredicate(t, pr(impl.StartsWith("abcdef")), expectation{
		value:    "abc",
		errorMsg: "sequence of length 3 is too short to contain a subsequence of length 6",
	})
}

func TestContains(t *testing.T) {
	verifyPredicate(t, pr(impl.Contains("bcd")), expectation{value: "abcdef", pass: true})
	verifyPredicate(t, pr(impl.Contains("abcd")), expectation{value: "abcdef", pass: true})
	verifyPredicate(t, pr(impl.Contains("bcdef")), expectation{value: "abcdef", pass: true})
	verifyPredicate(t, pr(impl.Contains("bcd")), expectation{value: "azcdef", pass: false})
	verifyPredicate(t, pr(impl.Contains([]int{2, 3, 4})), expectation{value: []int{1, 2, 3, 4, 5}, pass: true})
	verifyPredicate(t, pr(impl.Contains([]int{2, 3, 4})), expectation{value: []int{1, 6, 3, 4, 5}, pass: false})
	verifyPredicate(t, pr(impl.Contains("abc")), expectation{
		value:    123,
		errorMsg: "value of type 'int' is not a sequence",
	})
	verifyPredicate(t, pr(impl.Contains(123)), expectation{
		value:    "abc",
		errorMsg: "value of type 'int' is not a sequence",
	})
	verifyPredicate(t, pr(impl.Contains("abcdef")), expectation{
		value:    "bcd",
		errorMsg: "sequence of length 3 is too short to contain a subsequence of length 6",
	})
}

func TestEndsWith(t *testing.T) {
	verifyPredicate(t, pr(impl.EndsWith("def")), expectation{value: "abcdef", pass: true})
	verifyPredicate(t, pr(impl.EndsWith("def")), expectation{value: "abcdzf", pass: false})
	verifyPredicate(t, pr(impl.EndsWith([]int{3, 4, 5})), expectation{value: []int{1, 2, 3, 4, 5}, pass: true})
	verifyPredicate(t, pr(impl.EndsWith([]int{3, 4, 5})), expectation{value: []int{1, 2, 3, 6, 5}, pass: false})
	verifyPredicate(t, pr(impl.HasSuffix("abc")), expectation{
		value:    123,
		errorMsg: "value of type 'int' is not a sequence",
	})
	verifyPredicate(t, pr(impl.EndsWith(123)), expectation{
		value:    "abc",
		errorMsg: "value of type 'int' is not a sequence",
	})
	verifyPredicate(t, pr(impl.EndsWith("abcdef")), expectation{
		value:    "def",
		errorMsg: "sequence of length 3 is too short to contain a subsequence of length 6",
	})
}
