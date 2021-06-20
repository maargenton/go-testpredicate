/*

Package testpredicate -- Predicate-base test assertions library for Go

Package go-testpredicate is a test assertions library exposing a
predicate-like syntax that works with Go testing support to provide extensive
diagnostics output and reduces the need to use a debugger on every failing test.

The library contains an extensive collection of built-in predicates covering:

	- basic tests for nil, true, false
	- equality between any type of value
	- ordered comparison on numeric, string and sequence values
	- regexp match on strings
	- sub-sequences match on strings and sequences
	- set conditions on unordered collections
	- panic conditions on code fragment execution

Installation

    go get github.com/maargenton/go-testpredicate

Usage

	package examples_test

	import (
		"testing"

		"github.com/maargenton/go-testpredicate/pkg/require"
		"github.com/maargenton/go-testpredicate/pkg/verify"
	)

	func TestExample(t *testing.T) {
		t.Run("Given ", func(t *testing.T) {
			require.That(t, 123).ToString().Length().Eq(3)

			t.Run("when ", func(t *testing.T) {
				t.Run("then ", func(t *testing.T) {
					verify.That(t, "123").Eq(123)
					verify.That(t, 123).ToString().Length().Eq(4)
				})
			})
		})
	}

Output

	--- FAIL: TestExample (0.00s)
		--- FAIL: TestFoo/Given_ (0.00s)
			--- FAIL: TestFoo/Given_/when_ (0.00s)
				--- FAIL: TestFoo/Given_/when_/then_ (0.00s)
					usage_test.go:16:
						expected: value == 123
						error:    values of type 'string' and 'int' are never equal
						value:    "123"
					usage_test.go:17:
						expected: length(value.String()) == 4
						value:    123
						string:   "123"
						length:   3
*/
package testpredicate
