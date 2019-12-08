/*

Package testpredicate - Unit-testing predicates for Go.

Package go-testpredicate complements the built-in go testing package with a
way to define expectations in a concise and self describing way. When a test
fails, the default error message should provide enough context information
to understand the reason for the failure without having to resort to
debugging, at least most of the time.

A collection of built-in predicates are defined in the pred sub-package.
Additional predicates can easily be defined to support custom types and custom
validation needs, either locally within your test packages, or globally as
standalone predicate packages.

Usage

	package examples_test

	import (
		"testing"

		"github.com/maargenton/go-testpredicate/pkg/asserter"
		"github.com/maargenton/go-testpredicate/pkg/p"
	)

	func TestExample(t *testing.T) {
		assert := testpredicate.NewAsserter(t)
		assert.That(123, p.Lt(123))
	}

Output

	--- FAIL: TestExample (0.00s)
	/go/src/github.com/maargenton/go-testpred/examples/example_test.go:12:
		expected: value < 123,
		value: 123
*/
package testpredicate
