package bdd_test

import (
	"io"
	"reflect"
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/bdd"
	"github.com/maargenton/go-testpredicate/pkg/verify"
)

func TestUsed(t *testing.T) {
	var v = 0 // bdd.Used() should suppress the 'declared and not used' error
	bdd.Used(v)

	// Nothing further to test here
}

func TestTypeOf(t *testing.T) {
	verify.That(t, bdd.TypeOf[int]()).Eq(reflect.TypeOf(int(0)))
	verify.That(t, bdd.TypeOf[string]()).Eq(reflect.TypeOf(""))
	verify.That(t, bdd.TypeOf[float64]()).Eq(reflect.TypeOf(float64(0.0)))
	verify.That(t, bdd.TypeOf[[]string]()).Eq(reflect.TypeOf([]string{}))

	verify.That(t, bdd.TypeOf[io.Reader]()).ToString().Eq("io.Reader")
}

func TestFirstSecondThird(t *testing.T) {
	var f3 = func() (int, string, float64) { return 1, "two", 3.0 }

	bdd.Given(t, "a multi-value return function", func(t *bdd.T) {
		t.When("wrapped with bdd.First()", func(t *bdd.T) {
			t.Then("it returns the first value", func(t *bdd.T) {
				verify.That(t, bdd.First(f3())).Eq(1)
			})
		})
		t.When("wrapped with bdd.Second()", func(t *bdd.T) {
			t.Then("it returns the second value", func(t *bdd.T) {
				verify.That(t, bdd.Second(f3())).Eq("two")
			})
		})
		t.When("wrapped with bdd.Third()", func(t *bdd.T) {
			t.Then("it returns the third value", func(t *bdd.T) {
				verify.That(t, bdd.Third(f3())).Eq(3.0)
			})
		})
	})
}
