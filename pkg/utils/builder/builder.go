package builder

//go:generate go run github.com/maargenton/go-testpredicate/cmd/codegen/forward_api ../predicate/impl builder_api.tmpl

import (
	"fmt"
	"runtime"

	"github.com/maargenton/go-testpredicate/pkg/utils/predicate"
)

// Builder is the main type of the package, used to build predicates by chaining
// calls to transformation builder functions followed by one final predicate
// builder function.
type Builder struct {
	t        predicate.T
	value    interface{}
	file     string
	line     int
	p        predicate.Predicate
	required bool

	Ctx []predicate.ContextValue
}

// New returns a new predicate builder capture the given test context, value and
// the required flag indicating that a failed evaluation should fail the test.
func New(t predicate.T, value interface{}, required bool) *Builder {
	return &Builder{
		t:        t,
		value:    value,
		required: required,
	}
}

// CaptureCallsite is intended to be called right after creation of a new
// builder, to capture the callsite location and  allow better reporting.
func CaptureCallsite(b *Builder, skip int) {
	if _, file, line, ok := runtime.Caller(skip + 1); ok {
		b.file = file
		b.line = line
	}
}

// Evaluate is used to evaluate the predicate on the value and test context
// captured by the builder. Note that if no test context is set, no evaluation
// is performed.
func Evaluate(b *Builder) {
	if b.t == nil {
		return
	}
	b.t.Helper()

	success, ctx := b.p.Evaluate(b.value)
	if !success {
		ctx = append(ctx, b.Ctx...)
		b.t.Errorf("\n%v", predicate.FormatContextValues(ctx))
		if b.required {
			b.t.FailNow()
		}
	}
}

// VerifyCompletness is used during test cleanup to report improperly
// constructed predicates that do not evaluate any condition. Note that if no
// test context is set, no verification is performed.
func VerifyCompletness(b *Builder) {
	if b.t == nil {
		return
	}
	b.t.Helper()

	if b.p.Func == nil {
		prefix := ""
		if b.file != "" {
			prefix = fmt.Sprintf("%v:%v:\n", b.file, b.line)
		}
		b.t.Errorf("\n%vpredicate chain does not evaluate anything", prefix)
	}
}
