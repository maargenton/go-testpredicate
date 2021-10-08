package impl_test

import (
	"strings"
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/utils/predicate/impl"
)

func TestIsEqualSet(t *testing.T) {
	verifyPredicate(t, pr(impl.IsEqualSet([]int{1, 2, 3})), expectation{
		value: []int{3, 1, 2},
		pass:  true,
	})
	verifyPredicate(t, pr(impl.IsEqualSet([]int{1, 2, 3})), expectation{
		value: []int{3, 1, 5},
		pass:  false,
	})

	verifyPredicate(t, pr(impl.IsEqualSet(123)), expectation{
		value:    []int{3, 1, 5},
		errorMsg: "value of type 'int' is not an indexable collection",
	})
	verifyPredicate(t, pr(impl.IsEqualSet([]int{1, 2, 3})), expectation{
		value:    123,
		errorMsg: "value of type 'int' is not an indexable collection",
	})
}

func TestIsEqualSetWithLargeSet(t *testing.T) {
	var a = make([]int, 0, 50)
	for i := 0; i < 50; i++ {
		a = append(a, i)
	}
	var _, p = impl.IsEqualSet(a)
	success, ctx, _ := p([]int{0})
	if success {
		t.Errorf("\npredicate was expected to fail")
	}
	missing := ctx[1].Value.(string)
	if !strings.HasSuffix(missing, ", ...") {
		t.Errorf("\nexpected ellipsis at the en of truncated sequence\nwas: %v", missing)
	}
}

func TestIsDisjointSetFrom(t *testing.T) {
	verifyPredicate(t, pr(impl.IsDisjointSetFrom([]int{1, 2, 3})), expectation{
		value: []int{4, 5, 6},
		pass:  true,
	})
	verifyPredicate(t, pr(impl.IsDisjointSetFrom([]int{1, 2, 3})), expectation{
		value: []int{4, 5, 6, 2},
		pass:  false,
	})

	verifyPredicate(t, pr(impl.IsDisjointSetFrom(123)), expectation{
		value:    []int{3, 1, 5},
		errorMsg: "value of type 'int' is not an indexable collection",
	})
	verifyPredicate(t, pr(impl.IsDisjointSetFrom([]int{1, 2, 3})), expectation{
		value:    123,
		errorMsg: "value of type 'int' is not an indexable collection",
	})
}

func TestIsSubsetOf(t *testing.T) {
	verifyPredicate(t, pr(impl.IsSubsetOf([]int{1, 2, 3})), expectation{
		value: []int{3, 1},
		pass:  true,
	})
	verifyPredicate(t, pr(impl.IsSubsetOf([]int{1, 2, 3})), expectation{
		value: []int{3, 4},
		pass:  false,
	})

	verifyPredicate(t, pr(impl.IsSubsetOf(123)), expectation{
		value:    []int{3, 1, 5},
		errorMsg: "value of type 'int' is not an indexable collection",
	})
	verifyPredicate(t, pr(impl.IsSubsetOf([]int{1, 2, 3})), expectation{
		value:    123,
		errorMsg: "value of type 'int' is not an indexable collection",
	})
}

func TestIsSupersetOf(t *testing.T) {
	verifyPredicate(t, pr(impl.IsSupersetOf([]int{1, 2, 3})), expectation{
		value: []int{3, 1, 2, 4},
		pass:  true,
	})
	verifyPredicate(t, pr(impl.IsSupersetOf([]int{1, 2, 3})), expectation{
		value: []int{3, 5, 4, 6},
		pass:  false,
	})

	verifyPredicate(t, pr(impl.IsSupersetOf(123)), expectation{
		value:    []int{3, 1, 5},
		errorMsg: "value of type 'int' is not an indexable collection",
	})
	verifyPredicate(t, pr(impl.IsSupersetOf([]int{1, 2, 3})), expectation{
		value:    123,
		errorMsg: "value of type 'int' is not an indexable collection",
	})
}
