package impl_test

import (
	"io"
	"strings"
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/bdd"
	"github.com/maargenton/go-testpredicate/pkg/utils/predicate/impl"
)

func TestIsA(t *testing.T) {
	verifyPredicate(t, pr(impl.IsA(bdd.TypeOf[int]())), expectation{value: nil, pass: false})

	verifyPredicate(t, pr(impl.IsA(bdd.TypeOf[int]())), expectation{value: 123, pass: true})
	verifyPredicate(t, pr(impl.IsA(bdd.TypeOf[string]())), expectation{value: "test", pass: true})

	verifyPredicate(t, pr(impl.IsA(bdd.TypeOf[float64]())), expectation{value: 123.45, pass: true})
	verifyPredicate(t, pr(impl.IsA(bdd.TypeOf[int]())), expectation{value: 123.45, pass: false})

	var r = strings.NewReader("test")
	verifyPredicate(t, pr(impl.IsA(bdd.TypeOf[io.Reader]())), expectation{value: r, pass: true})
	verifyPredicate(t, pr(impl.IsA(bdd.TypeOf[io.Writer]())), expectation{value: r, pass: false})

}
