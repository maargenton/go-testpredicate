//go:build go1.25
// +build go1.25

package bdd_test

import (
	"testing"
	"time"

	"github.com/maargenton/go-testpredicate/pkg/bdd"
	"github.com/maargenton/go-testpredicate/pkg/verify"
)

func TestCanUseSyncTestInLeafContext(t *testing.T) {
	bdd.Given(t, "something", func(t *bdd.T) {
		t.When("doing something", func(t *bdd.T) {
			t.Then("a leaf synctest context can be used", func(t *bdd.T) {

				var t0 = time.Now()
				t.SyncTest(func(t *testing.T) {
					time.Sleep(5 * time.Second)
				})
				var elapsedMs = float64(time.Since(t0)) / float64(time.Millisecond)

				verify.That(t, elapsedMs).Le(float64(300))
			})
		})
	})
}
