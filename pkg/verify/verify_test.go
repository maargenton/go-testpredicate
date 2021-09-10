package verify_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/verify"
)

func TestExample(t *testing.T) {
	v := 123
	verify.That(t, v).ToString().Length().Eq(3)
	verify.That(t, v,
		verify.Context{Name: "double", Value: v * 2},
	).ToString().Length().Eq(3)
}
