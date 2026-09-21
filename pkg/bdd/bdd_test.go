package bdd_test

import (
	"fmt"
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/bdd"
	"github.com/maargenton/go-testpredicate/pkg/subexpr"
	"github.com/maargenton/go-testpredicate/pkg/utils/predicate"
	"github.com/maargenton/go-testpredicate/pkg/verify"
)

func TestVariablesWithinGivenBlockAreResetForEveryBranch(t *testing.T) {
	i := 0
	bdd.Given(t, "something", func(t *bdd.T) {
		i++
		j := 0
		t.When("doing something", func(t *bdd.T) {
			j++
			t.With("something", func(t *bdd.T) {
				t.Then("something happens", func(t *bdd.T) {
					verify.That(t, i).Eq(1)
					verify.That(t, j).Eq(1)
				})
				t.Then("something else happens", func(t *bdd.T) {
					verify.That(t, i).Eq(2)
					verify.That(t, j).Eq(1)
				})
			})
		})
		t.When("doing something else", func(t *bdd.T) {
			t.Then("something happens", func(t *bdd.T) {
				j++
				verify.That(t, i).Eq(3)
				verify.That(t, j).Eq(1)
			})
			t.Then("something else happens", func(t *bdd.T) {
				verify.That(t, i).Eq(4)
				verify.That(t, j).Eq(0)
			})
		})
	})
}

func TestUnderlyingT(t *testing.T) {
	var testStack []*testing.T

	testStack = append(testStack, t)

	bdd.Given(t, "something", func(t *bdd.T) {
		testStack = append(testStack, t.UnderlyingT())
		t.When("doing something", func(t *bdd.T) {
			testStack = append(testStack, t.UnderlyingT())
			t.With("something", func(t *bdd.T) {
				testStack = append(testStack, t.UnderlyingT())
				t.Then("each UnderlyingT is unique", func(t *bdd.T) {
					testStack = append(testStack, t.UnderlyingT())

					verify.That(t, testStack).All(
						subexpr.Value().Eval(countIn(testStack)).Eq(1),
					)

				})
			})
		})
	})
}

func countOf[T comparable](slice []T, item T) int {
	var count = 0
	for _, v := range slice {
		if v == item {
			count++
		}
	}
	return count
}

func countIn[T comparable](slice []T) (desc string, f predicate.TransformFunc) {
	return "slice.countOf( {} )", func(value any) (result any, ctx []predicate.ContextValue, err error) {
		var v, ok = value.(T)
		if !ok {
			err = fmt.Errorf("value of type '%T' is not countable in '%T'",
				value, slice,
			)
			return
		}
		result = countOf(slice, v)
		ctx = append(ctx, predicate.ContextValue{Name: "slice", Value: slice})
		ctx = append(ctx, predicate.ContextValue{Name: "count", Value: result})
		return
	}
}
