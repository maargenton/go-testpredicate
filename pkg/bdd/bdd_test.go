package bdd_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/bdd"
	"github.com/maargenton/go-testpredicate/pkg/verify"
)

func TestVariablesWithinGivenBlockAreResetForEveryBranch(t *testing.T) {
	i := 0
	bdd.Given(t, "something", func(t *bdd.T) {
		i++
		j := 0
		t.When("doing something", func(t *bdd.T) {
			j++
			t.Then("something happens", func(t *bdd.T) {
				verify.That(t, i).Eq(1)
				verify.That(t, j).Eq(1)
			})
			t.Then("something else happens", func(t *bdd.T) {
				verify.That(t, i).Eq(2)
				verify.That(t, j).Eq(1)
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
