# go-testpredicate

Test assertions library using predicate-like syntax, producing extensive
diagnostics output

[![Latest](
  https://img.shields.io/github/v/tag/maargenton/go-testpredicate?color=blue&label=latest&logo=go&logoColor=white&sort=semver)](
  https://pkg.go.dev/github.com/maargenton/go-testpredicate)
[![Build](
  https://img.shields.io/github/workflow/status/maargenton/go-testpredicate/build?label=build&logo=github&logoColor=aaaaaa)](
  https://github.com/maargenton/go-testpredicate/actions?query=branch%3Amaster)
[![Codecov](
  https://img.shields.io/codecov/c/github/maargenton/go-testpredicate?label=codecov&logo=codecov&logoColor=aaaaaa&token=fVZ3ZMAgfo)](
  https://codecov.io/gh/maargenton/go-testpredicate)
[![Go Report Card](
  https://goreportcard.com/badge/github.com/maargenton/go-testpredicate)](
  https://goreportcard.com/report/github.com/maargenton/go-testpredicate)


---------------------------

Package `go-testpredicate` is a test assertions library exposing a
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


## Installation

    go get github.com/maargenton/go-testpredicate

## Usage

```go
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
```

Output:
```
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
```

## API changes and stability

As of v0.5.0, a new and improved API is now available, intended to replace the
original API in the long run. The original API remains available but should now
be considered deprecated, and will be remove by v1.0.0.

### New API

In the new API, either `require.That()` or `verify.That()` is used to capture
both the testing context and the value under test, and the test assertion is
built through call chaining of optional transformations and a final condition.
Both `require` and `verify` can be mixed and matched as needed.

There is no more need to declare of an asserter object in every block, making
DBB-style more streamlined. There is no more package `p` sitting on the global
namespace, and the call-chaining syntax makes complex predicates more readable.

```go
package example_test

import (
    "testing"
    "github.com/maargenton/go-testpredicate/pkg/require"
    "github.com/maargenton/go-testpredicate/pkg/verify"
)

func TestExample(t *testing.T) {
    require.That(t, 123).ToString().Length().Eq(3)
    verify.That(t, 123).ToString().Length().Eq(4)
}
```

### Original API

The original API was based on an `assert` object, capturing the test context
`t`, with an option to either abort on any error or not. Complex predicate,
applying transformation to the value under test were hard to read due to the
layout of calling parentheses. BDD-style given/when/then structure was possible
but required the redefinition of an `assert` object in every nested block. The
reliance on package `p` as a short-hand for predicate was questionable at best.

```go
package example_test

import (
    "testing"
    "github.com/maargenton/go-testpredicate/pkg/asserter"
    "github.com/maargenton/go-testpredicate/pkg/p"
)

func TestExample(t *testing.T) {
    assert := asserter.New(t, asserter.AbortOnError())
    assert.That(123, p.ToString(p.Length(p.Eq(3))))
}
```


## Built-in predicates

All predicates are built through call chaining on the builder object returned by
`require.That()` or `verify.That()`. For an up-to-date full list of supported
predicates and their use, take a look at
`pkg/internal/builder/builder_api_test.go`

```go
func TestCollectionAPI(t *testing.T) {
    verify.That(t, []string{"a", "bb", "ccc"}).All(
        subexpr.Value().Length().Lt(5))
    verify.That(t, []string{"a", "bb", "ccc"}).Any(
        subexpr.Value().Length().Ge(3))
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

func TestMapAPI(t *testing.T) {
    var m = map[string]string{ "aaa": "bbb", "ccc": "ddd" }

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
```
