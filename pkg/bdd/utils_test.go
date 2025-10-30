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
