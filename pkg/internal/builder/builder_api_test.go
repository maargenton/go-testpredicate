package builder_test

import (
	"fmt"
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/internal/predicate"
	"github.com/maargenton/go-testpredicate/pkg/subexpr"
	"github.com/maargenton/go-testpredicate/pkg/verify"
)

// NOTE: This file contains calls to each of the generate Builder API functions,
// but only verifies the passing case. Tests for predicates failures and errors
// are expected to be handled in the `predicate/impl` package.

func TestCollectionAPI(t *testing.T) {
	verify.That(t, []string{"a", "bb", "ccc"}).All(
		subexpr.Value().Length().Lt(5))
	verify.That(t, []string{"a", "bb", "ccc"}).Any(
		subexpr.Value().Length().Ge(3))

	verify.That(t, [][]string{{"a", "bb", "cc"}, {"a", "bb", "ccc"}}).All(
		subexpr.Value().All(
			subexpr.Value().Length().Lt(5)))
}
func TestCompareAPI(t *testing.T) {
	verify.That(t, true).IsTrue()
	verify.That(t, false).IsFalse()
	verify.That(t, nil).IsNil()
	verify.That(t, &struct{}{}).IsNotNil()
	verify.That(t, 123).IsEqualTo(123)
	verify.That(t, 123).IsNotEqualTo(124)
	verify.That(t, 123).Eq(123)
	verify.That(t, 123).Ne(124)
}

func TestErrorAPI(t *testing.T) {
	var sentinel = fmt.Errorf("sentinel")
	var err = fmt.Errorf("error: %w", sentinel)
	verify.That(t, err).IsError(sentinel)
}

func TestExtAPI(t *testing.T) {
	var customPredicate = func() (desc string, f predicate.PredicateFunc) {
		desc = "{} is custom"
		f = func(
			value interface{}) (
			success bool, ctx []predicate.ContextValue, err error) {

			return value == nil, nil, nil
		}
		return
	}
	verify.That(t, nil).Is(customPredicate())

	var customTransform = func() (desc string, f predicate.TransformFunc) {
		desc = "custom({})"
		f = func(
			value interface{}) (
			r interface{}, ctx []predicate.ContextValue, err error) {

			return value, nil, nil
		}
		return
	}
	verify.That(t, nil).Eval(customTransform()).Is(customPredicate())

	verify.That(t, 9).Passes(subexpr.Value().Lt(10))
}

func TestMapAPI(t *testing.T) {
	var m = map[string]string{
		"aaa": "bbb",
		"ccc": "ddd",
	}
	verify.That(t, m).MapKeys().IsEqualSet([]string{"aaa", "ccc"})
	verify.That(t, m).MapValues().IsEqualSet([]string{"bbb", "ddd"})
}

func TestOrderedAPI(t *testing.T) {
	verify.That(t, 123).IsLessThan(124)
	verify.That(t, 123).IsLessOrEqualTo(123)
	verify.That(t, 123).IsGreaterThan(122)
	verify.That(t, 123).IsGreaterOrEqualTo(123)
	verify.That(t, 123).IsCloseTo(133, 10)
	verify.That(t, 123).Lt(124)
	verify.That(t, 123).Le(123)
	verify.That(t, 123).Gt(122)
	verify.That(t, 123).Ge(123)
}

func TestPanicAPI(t *testing.T) {
	verify.That(t, func() {
		panic(123)
	}).Panics()

	verify.That(t, func() {
		panic(123)
	}).PanicsAndRecoveredValue().Eq(123)
}

func TestSequenceAPI(t *testing.T) {
	verify.That(t, make([]int, 3, 5)).Length().Eq(3)
	verify.That(t, make([]int, 3, 5)).Capacity().Eq(5)

	verify.That(t, []int{}).IsEmpty()
	verify.That(t, []int{1, 2, 3, 4, 5}).IsNotEmpty()
	verify.That(t, []int{1, 2, 3, 4, 5}).StartsWith([]int{1, 2})
	verify.That(t, []int{1, 2, 3, 4, 5}).Contains([]int{2, 3, 4})
	verify.That(t, []int{1, 2, 3, 4, 5}).EndsWith([]int{4, 5})

	verify.That(t, []int{1, 2, 3, 4, 5}).HasPrefix([]int{1, 2})
	verify.That(t, []int{1, 2, 3, 4, 5}).HasSuffix([]int{4, 5})
}

func TestSetAPI(t *testing.T) {
	verify.That(t, []int{1, 2, 3, 4, 5}).IsEqualSet([]int{1, 4, 3, 2, 5})
	verify.That(t, []int{1, 2, 3, 4, 5}).IsDisjointSetFrom([]int{6, 9, 8, 7})
	verify.That(t, []int{1, 2, 3, 4, 5}).IsSubsetOf([]int{1, 4, 3, 2, 5, 6})
	verify.That(t, []int{1, 2, 3, 4, 5}).IsSupersetOf([]int{1, 4, 5})
}

func TestStringAPI(t *testing.T) {
	verify.That(t, "123").Matches(`\d+`)
	verify.That(t, 123).ToString().Eq("123")
	verify.That(t, "aBc").ToLower().Eq("abc")
	verify.That(t, "aBc").ToUpper().Eq("ABC")
}

func TestStructAPI(t *testing.T) {
	var v = struct {
		Name  string
		Value string
	}{
		Name:  "name",
		Value: "value",
	}
	verify.That(t, v).Field("Name").Eq("name")
}
