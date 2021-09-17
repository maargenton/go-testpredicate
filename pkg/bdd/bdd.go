// Package bdd defines a BDD-style bifurcated evaluation context, compatible
// with the build-in testing package.
package bdd

import (
	"fmt"
	"testing"
	"time"
)

// Ter is an interface definition that exposes all the `testing.T` functions
// that are not redefined by `bdd.T`.
type Ter interface {
	testing.TB
	Deadline() (deadline time.Time, ok bool)
}

// T is a testing context, similar to testing.T, passed to the testing functions
// in the BDD style bifurcated test evaluation. It forwards most of the
// testing.T API by encapsulating the current test context as a `Ter` interface;
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
	Ter
	t       *testing.T
	tracker *tracker
}

// Run defines a new fork in the current bifurcated evaluation context
func (b *T) Run(name string, f func(b *T)) bool {
	success := true
	if b.tracker.Active() {
		success = success && b.t.Run(name, func(t *testing.T) {
			f(&T{t, t, b.tracker.SubTracker()})
		})
	}
	return success
}

// When is a BDD-style syntactic sugar that wraps the `Run()` function and
// prefixes the name with `when ...`
func (b *T) When(name string, f func(b *T)) bool {
	name = fmt.Sprintf("when %v", name)
	return b.Run(name, f)
}

// Then is a BDD-style syntactic sugar that wraps the `Run()` function and
// prefixes the name with `then ...`
func (b *T) Then(name string, f func(b *T)) bool {
	name = fmt.Sprintf("then %v", name)
	return b.Run(name, f)
}

// Given is the root function that wraps the top level testing.T context and
// starts a bifurcated bdd.T evaluation context.
func Given(t *testing.T, name string, f func(b *T)) bool {
	name = fmt.Sprintf("Given %v", name)
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
