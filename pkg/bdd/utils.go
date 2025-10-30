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
