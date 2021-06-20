package subexpr

import "github.com/maargenton/go-testpredicate/pkg/internal/builder"

func Value() *builder.Builder {
	var b = builder.New(nil, nil, false)
	builder.CaptureCallsite(b, 1)
	return b
}
