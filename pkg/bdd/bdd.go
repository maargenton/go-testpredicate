// Package bdd defines a BDD-style bifurcated evaluation context, compatible
// with the build-in testing package.
package bdd

import (
	"fmt"
	"testing"
)

// T is a testing context, similar to testing.T, passed to the testing functions
// in the BDD style bifurcated test evaluation. It forwards most of the
// testing.T API by encapsulating the current test context as a `TB` interface;
// only `Parallel()` is dropped as non-relevant.
//
// The `Run()` function has a similar interface as testing.T, but passing a
// bdd.T object for the nested testing context, and evaluating all branches in a
// BDD style bifurcated fashion, where each branch is fully and independently
// re-evaluated for every leaf.
//
// Additional functions `When()` and `Then()` are syntactic sugar on top the
// `Run()` function. `bdd.Given()` is the root function that initializes the
// bifurcated evaluation context and runs all the branches.
type T struct {
	TB
	t       *testing.T
	tracker *tracker
}

// Run defines a new fork in the current bifurcated evaluation context.
func (b *T) Run(name string, f func(b *T)) bool {
	success := true
	if b.tracker.Active() {
		success = success && b.t.Run(name, func(t *testing.T) {
			f(&T{t, t, b.tracker.SubTracker()})
		})
	}
	return success
}

// When adds syntactic sugar on top of `bdd.T.Run()` and prefixes the name
// of the section with 'when ...'.
func (b *T) When(name string, f func(b *T)) bool {
	name = fmt.Sprintf("when %v", name)
	return b.Run(name, f)
}

// Then adds syntactic sugar on top of `bdd.T.Run()` and prefixes the name
// of the section with 'then ...'.
func (b *T) Then(name string, f func(b *T)) bool {
	name = fmt.Sprintf("then %v", name)
	return b.Run(name, f)
}

// Wrap is the root function that wraps the top level testing.T context and
// starts a bifurcated bdd.T evaluation context.
func Wrap(t *testing.T, name string, f func(b *T)) bool {
	tracker := &tracker{}
	success := true
	for tracker.Next() {
		if tracker.Active() {
			s := t.Run(name, func(t *testing.T) {
				f(&T{t, t, tracker.SubTracker()})
			})
			success = success && s
		}
	}
	return success
}

// Given adds syntactic sugar on top of `bdd.Wrap()` and prefixes the name
// of the section with 'Given ...'.
func Given(t *testing.T, name string, f func(b *T)) bool {
	name = fmt.Sprintf("Given %v", name)
	return Wrap(t, name, f)
}
