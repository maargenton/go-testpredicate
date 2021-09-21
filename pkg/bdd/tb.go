//go:build go1.15
// +build go1.15

package bdd

import (
	"testing"
	"time"
)

// TB is an interface definition that exposes all the `testing.T` functions
// that are not redefined by `bdd.T`.
type TB interface {
	testing.TB
	Deadline() (deadline time.Time, ok bool)
}
