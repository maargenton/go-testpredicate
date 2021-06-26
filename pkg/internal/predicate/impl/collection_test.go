package impl_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/internal/predicate"
	"github.com/maargenton/go-testpredicate/pkg/internal/predicate/impl"
)

func TestAll(t *testing.T) {
	var p = &predicate.Predicate{}
	p.RegisterPredicate(impl.Lt(3))

	verifyPredicate(t, pr(impl.All(p)), expectation{value: []int{1, 2}, pass: true})
	verifyPredicate(t, pr(impl.All(p)), expectation{value: []int{1, 2, 3}, pass: false})
	verifyPredicate(t, pr(impl.All(p)), expectation{
		value:    2,
		errorMsg: "value of type 'int' is not a collection",
	})
}

func TestAny(t *testing.T) {
	var p = &predicate.Predicate{}
	p.RegisterPredicate(impl.Lt(3))

	verifyPredicate(t, pr(impl.Any(p)), expectation{value: []int{3, 4, 2, 5}, pass: true})
	verifyPredicate(t, pr(impl.Any(p)), expectation{value: []int{3, 4, 5}, pass: false})
	verifyPredicate(t, pr(impl.Any(p)), expectation{
		value:    2,
		errorMsg: "value of type 'int' is not a collection",
	})
}

func TestAnyNested(t *testing.T) {
	var p0 = &predicate.Predicate{}
	p0.RegisterPredicate(impl.Lt(3))

	var p1 = &predicate.Predicate{}
	p1.RegisterPredicate(impl.Any(p0))

	verifyPredicate(t, pr(impl.Any(p1)), expectation{value: [][]int{{3, 4, 5}}, pass: false})
}
