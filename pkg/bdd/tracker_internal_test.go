package bdd

import (
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/require"
)

func TestTrackerInitiallyInvalid(t *testing.T) {
	tracker := &tracker{}
	require.That(t, func() {
		tracker.Active()
	}).PanicsAndRecoveredValue().Eq(
		"tracker.Active() called before interation started with tracker.Next()")

	require.That(t, func() {
		tracker.SubTracker()
	}).PanicsAndRecoveredValue().Eq(
		"tracker.SubTracker() called before interation started with tracker.Next()")
}

func TestTrackerRepeatsRootForEachLeaf(t *testing.T) {
	var branches []string
	var tracker = &tracker{}

	for tracker.Next() {
		if tracker.Active() {
			branches = append(branches, "a")
			tracker := tracker.SubTracker()
			if tracker.Active() {
				branches = append(branches, "aa")
			}
			if tracker.Active() {
				branches = append(branches, "ab")
			}
		}
	}

	require.That(t, branches).Eq([]string{"a", "aa", "a", "ab"})
}

func TestTrackerEavluatesAllNestedBirfurcationsBeforeNextSibling(t *testing.T) {
	var branches []string
	var tracker = &tracker{}

	for tracker.Next() {
		if tracker.Active() {
			branches = append(branches, "a")
			tracker := tracker.SubTracker()
			if tracker.Active() {
				branches = append(branches, "aa")
				tracker := tracker.SubTracker()
				if tracker.Active() {
					branches = append(branches, "aaa")
				}
				if tracker.Active() {
					branches = append(branches, "aab")
				}
			}
			if tracker.Active() {
				branches = append(branches, "ab")
			}
		}
	}

	require.That(t, branches).Eq([]string{
		"a", "aa", "aaa",
		"a", "aa", "aab",
		"a", "ab"})
}
