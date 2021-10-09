# go-testpredicate

Test predicate style assertions library with extensive diagnostics output.

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

Package `go-testpredicate` is an assertions library exposing a test predicate
style syntax for use with the built-in Go `testing` package, producing extensive
diagnostics output and reducing the need for debugging failing tests.

The library contains an extensive collection of built-in predicates covering:

- basic tests for nil, true, false
- equality between any type of value
- ordered comparison on numeric, string and sequence values
- regexp match on strings
- sub-sequences match on strings and sequences
- set conditions on unordered collections
- panic conditions on code fragment execution

It also includes a BDD-style bifurcated evaluation context, where each test
section is potentially evaluated multiple times in order to evaluate each branch
independently.


## Installation

```
go get github.com/maargenton/go-testpredicate
```

Optionally, you can add predefined code snippets for your text editor or IDE to
assist in writing  your test code. Snippets for VSCode are available
[here](docs/snippets.md)

## Usage

```go
package example_test

import (
    "testing"

    "github.com/maargenton/go-testpredicate/pkg/bdd"
    "github.com/maargenton/go-testpredicate/pkg/require"
    "github.com/maargenton/go-testpredicate/pkg/verify"
)

func TestExample(t *testing.T) {
    bdd.Given(t, "something", func(t *bdd.T) {
        require.That(t, 123).ToString().Length().Eq(3)

        t.When("doing something", func(t *bdd.T) {
            t.Then("something happens ", func(t *bdd.T) {
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
    --- FAIL: TestExample/Given_something (0.00s)
        --- FAIL: TestExample/Given_something/when_doing_something (0.00s)
            --- FAIL: TestExample/Given_something/when_doing_something/then_something_happens_ (0.00s)
                example_test.go:17:
                    expected: value == 123
                    error:    values of type 'string' and 'int' are never equal
                    value:    "123"
                example_test.go:18:
                    expected: length(value.String()) == 4
                    value:    123
                    string:   "123"
                    length:   3
```

## API changes and stability

Older version of this package where exposing a different API that has since been
deprecated, and has now been remove for the v1.0.0 release. The latest version
supporting the legacy API is v0.6.4.

Predicates are constructed starting with either `require.That(t, <value>)` or
`verify.That(t, <value>)`, where _require_ will abort the test on error, while
_verify_ will keep going. Both variants take the testing context `t`, and the
value to test.

Additional diagnostic context can be added to either functions with
`require.Context{}` / `verify.Context{}` passed as additional arguments.

```go
package example_test

import (
    "testing"
    "github.com/maargenton/go-testpredicate/pkg/require"
    "github.com/maargenton/go-testpredicate/pkg/verify"
)

func TestExample(t *testing.T) {
    v := 123
    require.That(t, v).ToString().Length().Eq(3)
    verify.That(t, v).ToString().Length().Eq(3)
    verify.That(t, v,
        verify.Context{Name: "double", Value: v * 2},
    ).ToString().Length().Eq(3)
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

func TestExtAPI(t *testing.T) {
    var customPredicate = func() (desc string, f predicate.PredicateFunc) {
        // ...
    }
    verify.That(t, nil).Is(customPredicate())

    var customTransform = func() (desc string, f predicate.TransformFunc) {
        // ...
    }
    verify.That(t, nil).Eval(customTransform()).Is(customPredicate())

    verify.That(t, 9).Passes(subexpr.Value().Lt(10))
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

## BDD-style bifurcated tests

### Rationale

First of all, the Go `testing` package is great and the fact that it is
standard, built in and integrated with the Go tooling infrastructure is awesome.
This is why the `go-testpredicate` packages strives to enhance it instead of
replacing it, unlike many other testing packages.

If you look at other unit-testing packages, in other languages, you will find
either traditional xUnit style packages relying on classes to define test suites
and fixtures and test cases, or more recent testing packages (like
[Catch-2](https://github.com/catchorg/Catch2) for C++) that provide, through
other means, ways to define setup and test cases than run independently. The
common pattern is that setup code, that may be shared by multiple test cases, is
usually re-evaluated for every test case so that, despite their potentially
mutating interactions with the setup, test cases don't affect each other.

Some great articles and blog posts have explained how the leverage nested
`t.Run()` calls to structure tests in way that is closer to BDD-style given /
when / then paradigm. Unfortunately, when using thees approaches, and especially
with shared setup sections, the test cases are no longer independent, as all
branches are run sequentially, going up and down each branch and into the next
branch, without resetting the setup.

The `bdd` package in `go-testpredicate` provides a way to write tests with a
BDD-style structure, using the built-in `testing.T`, but evaluating the test
cases in a bifurcated fashion, repeating the evaluation of each entire branch
for every leaf test case, so that test cases are independent from each other
again.

### Usage overview

`bdd.Wrap()` or `bdd.Given()` are the root level function that setup and iterate
through the bifurcated test evaluation context. They define blocks that receive
a `bdd.T` instead of `testing.T`, but `bdd.T` is fully compatible with
`testing.T` and can be used with any third party library that expect either the
`testing.TB` interface or a subset of it (including out own `verify.That()` /
`require.That()`).

Nested and sibling bifurcated branches are defined with `t.Run()` (on `bdd.T`)
or `t.When()` / `t.Then()` for BDD style.

> **IMPORTANT:** In a bifurcated evaluation context, as defined by `bdd.T`, test
> scenarios are run repeatedly in order to evaluate each branch (from root to
> leaf) independently of each other. When a particular branch is being
> evaluated, all the other forks and sub-branches are skipped; the other
> branches are run in separated independent iterations of the scenario.

### Usage, traditional style

```go
package example_test

import (
    "testing"
    "github.com/maargenton/go-testpredicate/pkg/bdd"
)

func TesTraditional(t *testing.T) {

    // Global immutable setup code can go here

    bdd.Wrap(t, "Given something", func(t *bdd.T) {

        // Local mutable setup code goes here

        t.Run("something happens", func(t *bdd.T) {

            // When this code runs, the code in following `t.Run()` blocks
            // will be skipped.
        })
        t.Run("something else happens", func(t *bdd.T) {

            // When this code runs, all code in preceding `t.Run()` blocks
            // has been skipped and did not affect the local setup.
        })
    })
}
```

### Usage, BDD style

```go
package bdd_test

import (
    "testing"
    "github.com/maargenton/go-testpredicate/pkg/bdd"
)

func TestBDDStyle(t *testing.T) {

    // Global immutable setup code can go here

    bdd.Given(t, "something", func(t *bdd.T) {

        // Local mutable setup code goes here

        t.When("doing something", func(t *bdd.T) {

            // or here

            t.Then("something happens", func(t *bdd.T) {

                // When this code runs, the code in the following `t.Then()`
                // blocks will be skipped.
            })
            t.Then("something else happens", func(t *bdd.T) {

                // When this code runs, all code in preceding `t.Then()`
                // blocks has been skipped and did not affect the local setup.
            })
        })
    })
}
```
