# go-testpredicate

Unit-testing predicates for Go.

[![GoDoc](https://godoc.org/github.com/maargenton/go-testpredicate?status.svg)](https://godoc.org/github.com/maargenton/go-testpredicate)
[![Build Status](https://travis-ci.org/maargenton/go-testpredicate.svg?branch=master)](https://travis-ci.org/maargenton/go-testpredicate)
[![codecov](https://codecov.io/gh/maargenton/go-testpredicate/branch/master/graph/badge.svg)](https://codecov.io/gh/maargenton/go-testpredicate)
[![Go Report Card](https://goreportcard.com/badge/github.com/maargenton/go-testpredicate)](https://goreportcard.com/report/github.com/maargenton/go-testpredicate)

Package `go-testpredicate` complements the built-in go testing package with a
way to define expectations in a concise and self describing way. When a test
fails, the default error message should provide enough context information
to understand the reason for the failure without having to resort to
debugging, at least most of the time.

A collection of built-in predicates are defined in the pred sub-package.
Additional predicates can easily be defined to support custom types and custom
validation needs, either locally within your test packages, or globally as
standalone predicate packages.

## Installation

    go get github.com/maargenton/go-testpredicate

## Usage

```go
package examples_test


import (
    "testing"

    "github.com/maargenton/go-testpredicate/pkg/assert"
    "github.com/maargenton/go-testpredicate/pkg/p"
)

func TestExample(t *testing.T) {
    assert := testpredicate.NewAsserter(t)
    assert.That(123, p.Lt(123))
}
```

Output:
```go
--- FAIL: TestExample (0.00s)
/go/src/github.com/maargenton/go-testpred/examples/example_test.go:12:
    expected: value < 123,
    value: 123
```

## Built-in predicates

### Nil and equality

- `IsNil()`
- `IsNotNil()`
- `IsTrue()`
- `IsFalse()`
- `IsEqualTo( value )` / `Eq( value )`
- `IsNotEqualTo( value )` / `Ne( value )`
- `IsNoError()`, preferred to IsNil() for testing errors
- `IsError( error )`, preferred to IsEqualTo() for testing errors

### Order comparable values

- `LessThan( value )` / `Lt( value )`
- `LessOrEqualTo( value )` / `Le( value )`
- `GreaterThan( value )` / `Gt( value )`
- `GreaterOrEqualTo( value )` / `Ge( value )`
- `CloseTo( value, tolerance )`, for floating point values

### String and sequence values

- `IsEmpty()`
- `IsNotEmpty()`
- `StartsWith( value )`
- `Contains( value )`
- `EndsWith( value )`
- `Matches( regexp )`, for string values

### Sets

- `IsSubsetOf( collection )`
- `IsSupersetOf( collection )`
- `IsDisjointSetFrom( collection )`
- `IsEqualSet( collection )`

## Composable predicates

Some predicates are not directly testing against a specific value, but instead define how to transform the value before applying a nested predicate, or how to apply the nested predicate to the elements of a collection

### Strings

- `ToUpper( predicate )`: apply predicate to the uppercase version of a string
- `ToLower( predicate )`: apply predicate to the lowercase version of a string
- `ToString( predicate )`: apply predicate to the stringified version of a value

### Collections attributes

- `Length( predicate )`: apply predicate to the length of a collection or string
- `Capacity( predicate )`: apply predicate to the capacity of a collection
- `MapKeys( predicate )`: apply predicate to a collection made of the keys of a map
- `MapValues( predicate )`: apply predicate to a collection made of the values of a map

### Collections elements

- `All( predicate )`: all elements must match the predicate
- `Any( predicate )`: at least one element must match the predicate
- `AllKeys( predicate )`: all keys of a map must match the predicate
- `AnyKey( predicate )`: at least one key of a map must match the predicate
- `AllValues( predicate )`: all values of a map must match the predicate
- `AnyValue( predicate )`: at least one value of a map must match the predicate

## Special predicates

### Panic

- `Panics()`: evaluates value as a callable function and expect it to panic
- `PanicsAndResult( predicate )`: evaluates value as a callable function,
  expects it to panic, and evaluates the panic value against the nested
  predicate
