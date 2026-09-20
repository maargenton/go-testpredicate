//go:build go1.25
// +build go1.25

package bdd_test

import (
	"testing"
	"time"

	"github.com/maargenton/go-testpredicate/pkg/bdd"
)

func TestCanUseSyncTestInLeafContext(t *testing.T) {
	bdd.Given(t, "something", func(t *bdd.T) {
		t.When("doing something", func(t *bdd.T) {
			t.Then("a leaf synctest context can be used", func(t *bdd.T) {
				t.SyncTest(func(t *testing.T) {
					time.Sleep(1 * time.Hour)
				})
			})
		})
	})
}
