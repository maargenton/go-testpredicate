package predicate_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/internal/predicate"
	"github.com/maargenton/go-testpredicate/pkg/internal/predicate/impl"
)

func TestPredicateRegistration(t *testing.T) {
	var p = predicate.Predicate{}
	p.RegisterPredicate(impl.Eq(3))
	recoveredValue := capturePanic(func() {
		p.RegisterPredicate(impl.Eq(5))
	})
	if recoveredValue != "RegisterPredicate() should only be called once per predicate" {
		t.Errorf("\nunexpected recovered value:\n%v", recoveredValue)
	}
}

func capturePanic(f func()) (v interface{}) {
	defer func() {
		v = recover()
	}()
	f()
	return
}

func TestFormatDescription(t *testing.T) {
	var p = predicate.Predicate{}
	p.RegisterPredicate("{} is empty", nil)
	p.RegisterTransformation("{}.String()", nil)
	p.RegisterTransformation("ToLower({})", nil)

	desc := p.FormatDescription("x")
	if desc != "ToLower(x.String()) is empty" {
		t.Errorf("\nunexpected description: %v", desc)
	}
}

func TestEvaluateSimplePredicate(t *testing.T) {

	var p = predicate.Predicate{}
	p.RegisterPredicate(impl.Eq(3))

	if success, ctx := p.Evaluate(3); !success {
		t.Errorf("\nunexpected failure:\n%v", predicate.FormatContextValues(ctx))
	}

	if success, ctx := p.Evaluate("1234"); success {
		t.Errorf("\nunexpected success")
	} else {
		if ctx[2].Name != "error" {
			t.Errorf("\ncontext missing error:\n%v", predicate.FormatContextValues(ctx))
		}
	}
}

func TestEvaluateWithTransform(t *testing.T) {

	var p = predicate.Predicate{}
	p.RegisterTransformation(impl.Length())
	p.RegisterPredicate(impl.Eq(3))

	if success, ctx := p.Evaluate("123"); !success {
		t.Errorf("\nunexpected failure:\n%v", predicate.FormatContextValues(ctx))
	}

	if success, ctx := p.Evaluate(123); success {
		t.Errorf("\nunexpected success")
	} else {
		if ctx[2].Name != "error" {
			t.Errorf("\ncontext missing error:\n%v", predicate.FormatContextValues(ctx))
		}
	}
}
