package impl_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/internal/predicate/impl"
)

func TestField(t *testing.T) {
	var v = struct {
		Name  string
		Value string
	}{
		Name:  "name",
		Value: "value",
	}

	verifyTransform(t, tr(impl.Field("Name")), expectation{
		value:  v,
		result: "name",
	})
	verifyTransform(t, tr(impl.Field("Undefined")), expectation{
		value:  v,
		result: nil,
	})
}
