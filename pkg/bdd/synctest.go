//go:build go1.25
// +build go1.25

package bdd

import (
	"testing"
	"testing/synctest"
)

// SyncTest calls `synctest.Test` in the current BDD test context, passing out a
// regular `*testing.T` instance because synctest does not allow any further
// nesting within the synctest bubble.
func (t *T) SyncTest(f func(t *testing.T)) {
	synctest.Test(t.UnderlyingT(), func(bubbleT *testing.T) {
		f(bubbleT)
	})
}
