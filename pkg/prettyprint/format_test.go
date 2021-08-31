package prettyprint_test

import (
	"strings"
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/prettyprint"
)

func makeInts(n int) []int {
	var r []int
	for i := 0; i < n; i++ {
		r = append(r, i)
	}
	return r
}

type formatExpectation struct {
	lineCount int
	maxWidth  int
}

func validateFormat(t *testing.T, s string, exp *formatExpectation) {
	t.Helper()

	lines := strings.Split(s, "\n")
	if l := len(lines); l != exp.lineCount {
		t.Fatalf("\nexpected %v line(s)\n"+
			"actual:  %v line(s)\n"+
			"value:  |%v|\n", exp.lineCount, l, s)
	}

	for i, line := range lines {
		line = strings.Replace(line, "\t", "    ", -1)
		l := len([]rune(line))
		if l > exp.maxWidth {
			t.Fatalf("\nexpected all lines to be no more than %v long\n"+
				"length:  %v for line %v\n"+
				"ruler:  |0----v----1----v----2----v----3----v----4----v----5----v----6----v----7----v----8----v----| \n"+
				"line:   |%v|\n", exp.maxWidth, l, i, line)
		}
	}
}

func TestFormatValueWithShortListCollapsesValuesIntoASingleLine(t *testing.T) {
	v := makeInts(3)
	s := prettyprint.FormatValue(v)

	validateFormat(t, s, &formatExpectation{
		lineCount: 1,
		maxWidth:  81,
	})
}

func TestFormatValueCollapsesNestedLists(t *testing.T) {
	v := struct{ i []int }{i: makeInts(3)}
	s := prettyprint.FormatValue(v)

	validateFormat(t, s, &formatExpectation{
		lineCount: 3,
		maxWidth:  81,
	})
}

func TestFormatValueWithLongListCollapsesValuesIntoMultipleLines(t *testing.T) {
	v := makeInts(30)
	s := prettyprint.FormatValue(v)

	validateFormat(t, s, &formatExpectation{
		lineCount: 5,
		maxWidth:  80,
	})
}

func TestFormatValueWithWeryLongListTruncatesCollapsedLines(t *testing.T) {
	v := makeInts(200)
	s := prettyprint.FormatValue(v)

	validateFormat(t, s, &formatExpectation{
		lineCount: 14,
		maxWidth:  80,
	})
}

func TestFormatValueWithWeryLongListTruncatesCollapsedLinesToMaxWrapped(t *testing.T) {
	pp := prettyprint.New()
	pp.MaxWrapped = 5
	v := makeInts(200)
	s := pp.FormatValue(v)

	validateFormat(t, s, &formatExpectation{
		lineCount: 9,
		maxWidth:  80,
	})
}

// ---------------------------------------------------------------------------
// Test alignment of values in key-value tokens

func TestFormatValueAlignsValues(t *testing.T) {
	v := struct {
		a          int
		bbbbbbbbbb int
	}{
		a:          1,
		bbbbbbbbbb: 2,
	}
	s := prettyprint.FormatValue(v)

	lines := strings.Split(s, "\n")
	if l := lines[1]; l != "\ta:          1," {
		t.Fatalf("\nunexpected line 1: |%v|", l)
	}
	if l := lines[2]; l != "\tbbbbbbbbbb: 2," {
		t.Fatalf("\nunexpected line 2: |%v|", l)
	}
}

func TestFormatValueAlignsValuesInChunks(t *testing.T) {
	v := struct {
		a          int
		bbbbbbbbbb int
		c          []int
		d          int
		eee        int
	}{
		a:          1,
		bbbbbbbbbb: 2,
		c:          makeInts(20),
		d:          4,
		eee:        5,
	}
	s := prettyprint.FormatValue(v)

	lines := strings.Split(s, "\n")

	if l := lines[1]; l != "\ta:          1," {
		t.Fatalf("\nunexpected line 1: |%v|", l)
	}
	if l := lines[2]; l != "\tbbbbbbbbbb: 2," {
		t.Fatalf("\nunexpected line 2: |%v|", l)
	}

	if l := lines[7]; l != "\td:   4," {
		t.Fatalf("\nunexpected line 7: |%v|", l)
	}
	if l := lines[8]; l != "\teee: 5," {
		t.Fatalf("\nunexpected line 8: |%v|", l)
	}
}

// ---------------------------------------------------------------------------
// Test long tokens are wrapped onto multiple lines

func TestFormatValueWrapsLongStringTokens(t *testing.T) {
	v := strings.Repeat("abcdefghi ", 20)
	s := prettyprint.FormatValue(v)

	validateFormat(t, s, &formatExpectation{
		lineCount: 3,
		maxWidth:  80,
	})
}

func TestFormatValueForceWrapsLongStringTokens(t *testing.T) {
	v := strings.Repeat("abcdefghi ", 20) +
		strings.Repeat("_aaabbb_", 20) +
		strings.Repeat("abcdefghi ", 20)
	s := prettyprint.FormatValue(v)

	validateFormat(t, s, &formatExpectation{
		lineCount: 8,
		maxWidth:  81,
	})
}

func TestFormatValueCanBreakOnEscapedCharacters(t *testing.T) {
	v := strings.Repeat("abcdefghi\n", 20)
	s := prettyprint.FormatValue(v)

	validateFormat(t, s, &formatExpectation{
		lineCount: 4,
		maxWidth:  81,
	})

	lines := strings.Split(s, "\n")
	for i, line := range lines[:3] {
		if !strings.HasSuffix(line, "\\n↩") {
			t.Fatalf(
				"\nexpected line breaks on '\\n' with wrap marker\n"+
					"line:    %v\n"+
					"actual: |...%v|",
				i, line[len(line)-10:],
			)
		}
	}
}

func TestFormatValueBreakAlignment(t *testing.T) {
	v := strings.Repeat("a bb ccc dddd eeeee ffffff ggggg hhhh iii jj k ", 80)
	s := prettyprint.FormatValue(v)

	validateFormat(t, s, &formatExpectation{
		lineCount: 51,
		maxWidth:  81,
	})
}

// ---------------------------------------------------------------------------
// Test formatting of interface{} type

func TestFormatInterfaceArray(t *testing.T) {
	v := []interface{}{
		"hello", "world", 42,
	}
	s := prettyprint.FormatValue(v)

	if s != `{ "hello", "world", 42 }` {
		t.Errorf("\nunexpected output: |%v|", s)
	}
}
