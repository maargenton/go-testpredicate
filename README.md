# go-testpredicate

Unit-testing predicates for Go.

[![GoDoc](https://godoc.org/github.com/marcus999/go-testpredicate?status.svg)](https://godoc.org/github.com/marcus999/go-testpredicate)
[![Build Status](https://travis-ci.org/marcus999/go-testpredicate.svg?branch=master)](https://travis-ci.org/marcus999/go-testpredicate)
[![Go Report Card](https://goreportcard.com/badge/github.com/marcus999/go-testpredicate)](https://goreportcard.com/report/github.com/marcus999/go-testpredicate)

Package `go-testpredicate` complements the built-in go testing package with a
way to define expectations in a concise and self describing way. When a test
fails, the default error message should provide enought context information
to understand the reason for the failure without having to resort to
debugging, at least most of the time.

A collection of built-in prediactes are defined in the pred sub-package.
Additional predicates can easily be defined to support custom types and custom
validation needs, either locally within your test packages, or globally as
standalone predicate packages.

## Installation

    go get github.com/marcus999/go-testpredicate

## Usage

```go
package examples_test

import (
    "testing"

    "github.com/marcus999/go-testpredredicate"
    "github.com/marcus999/go-testpredredicate/pred"
)

func TestExample(t *testing.T) {
    assert := testpredredicate.NewAsserter(t)
    assert.That(123, pred.Lt(123))
}
```

Output:
```go
--- FAIL: TestExample (0.00s)
/go/src/github.com/marcus999/go-testpred/examples/example_test.go:12:
    expected: value < 123,
    value: 123
```

## Built-in predicates

