package bdd_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/bdd"
)

func TestUsed(t *testing.T) {
	var v = 0 // bdd.Used() should suppress the 'declared and not used' error
	bdd.Used(v)

	// Nothing further to test here
}
