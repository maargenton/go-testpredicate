package predicate

import (
	"fmt"
	"io"
	"strings"

	"github.com/maargenton/go-testpredicate/pkg/utils/prettyprint"
)

// ContextValue capture one original, intermediate or final value generated
// during the evaluation of a predicate. Each value has a `Name`, a `Value`, and
// a `Pre` flag indicating if the value is a preformatted string that should be
// printed as is, or an actual value that should be printed with adequate
// formatting.
type ContextValue struct {
	Name  string
	Value interface{}
	Pre   bool
}

// FormatContextValues returns a string containing the formated print-out of the
// context values.
func FormatContextValues(values []ContextValue) string {
	width := 0
	for _, c := range values {
		if len(c.Name) > width {
			width = len(c.Name)
		}
	}

	var s strings.Builder
	var ordered = []string{"expected", "error", "value"}
	for _, name := range ordered {
		for _, c := range values {
			if c.Name != name {
				continue
			}
			formatContextValue(&s, c, width)
		}
	}

values_loop:
	for _, c := range values {
		for _, name := range ordered {
			if c.Name == name {
				continue values_loop
			}
		}
		formatContextValue(&s, c, width)
	}

	return s.String()
}

var defaultFormatter = prettyprint.Formatter{
	Width:      120,
	MinWidth:   40,
	WrapPrefix: "↩",
	WrapSuffix: "↪",
	MaxWrapped: 10,
	IndentStr:  "\t",
	NewlineStr: "\n",
}

func formatContextValue(w io.Writer, c ContextValue, width int) {
	var padding = strings.Repeat(" ", width-len(c.Name))

	fmt.Fprintf(w, "%v:%v ", c.Name, padding)
	if c.Pre {
		fmt.Fprintf(w, "%v\n", c.Value)
	} else {
		var formatter = defaultFormatter
		formatter.NewlineStr = "\n" + strings.Repeat(" ", width+2)
		fmt.Fprintf(w, "%v\n", formatter.FormatValue(c.Value))
	}
}
