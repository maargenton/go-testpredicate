package bdd

import "reflect"

// Used is a no-op function to mark variables as used, to avoid compiler errors
// while in the middle of writing tests
func Used(vv ...any) {}

// TypeOf returns the reflect.Type for a given type T, for use with IsA()
// predicate.
func TypeOf[T any]() reflect.Type {
	var p *T = nil
	return reflect.TypeOf(p).Elem()
}

// First can be used inline to wrap the call to a multi-value return function
// and extract the first value, ignoring the rest.
func First[A any](a A, _ ...any) A { return a }

// Second can be used inline to wrap the call to a multi-value return function
// and extract the second value, ignoring the rest.
func Second[A, B any](a A, b B, _ ...any) B { return b }

// Third can be used inline to wrap the call to a multi-value return function
// and extract the third value, ignoring the rest.
func Third[A, B, C any](a A, b B, c C, _ ...any) C { return c }
