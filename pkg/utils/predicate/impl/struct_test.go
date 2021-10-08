package impl_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/utils/predicate/impl"
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

	var desc, f = impl.Field("Value")
	if desc != "{}.Value" {
		t.Errorf("\nUnexpected description: %v", desc)
	}
	_, ctx, _ := f(v)
	if len(ctx) != 1 || ctx[0].Name != "$.Value" {
		t.Errorf("\nUnexpected ctx: %+v", ctx)

	}

}
