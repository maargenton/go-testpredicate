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

// var anonymousStruct = &struct {
// 	name    string
// 	path    string
// 	options []string
// 	nested  []struct {
// 		name  string
// 		value float64
// 	}
// 	index   []int
// 	details string
// }{
// 	name:    "my name",
// 	path:    "/aaa/bbb/ccc/d",
// 	options: []string{"readable", "aaa{", "bbb}", "build-target"},
// 	nested: []struct {
// 		name  string
// 		value float64
// 	}{
// 		struct {
// 			name  string
// 			value float64
// 		}{
// 			name:  "aaa",
// 			value: 123.456,
// 		},
// 	},
// 	index:   []int{1, 5, 2, 3, ')', 12, 1, 5, 2, 3, ')', 12, 1, 5, 2, 3, ')', 12, 1, 5, 2, 3, ')', 12, 1, 5, 2, 3, ')', 12},
// 	details: "MakeBoolPredicate wraps a predicate function\nreturning bool into a predicateinterface. Any error returned from the function is interpreted as an invalid evaluation.xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
// }

// func TestFormatValueWithStruct(t *testing.T) {
// 	assert := testpredicate.NewAsserter(t)
// 	assert.That(true, pred.Eq(true))

// 	s := prettyprint.FormatValue(anonymousStruct)
// 	_ = s
// 	fmt.Print(s)

// 	lines := strings.Split(s, "\n")
// 	assert.That(lines, pred.Length(pred.Eq(18)))
// 	assert.That(lines, pred.All(pred.Length(pred.Le(82))))
// }

// func TestFormatString(t *testing.T) {
// 	assert := testpredicate.NewAsserter(t)
// 	assert.That(true, pred.Eq(true))

// 	str := "MakeBoolPredicate wraps a predicate function\nreturning bool into a predicateinterface. Any error returned from the function is interpreted as an invalid evaluation."
// 	s := value.Format(str)
// 	lines := strings.Split(s, "\n")
// 	assert.That(lines, pred.Length(pred.GreaterThan(1)))
// 	assert.That(lines, pred.All(pred.Length(pred.Le(82))))
// }

// func makeInts(m, n int) []int {
// 	var result = make([]int, 0, n)
// 	for i := 0; i < n; i++ {
// 		result = append(result, rand.Intn(m))
// 	}

// 	return result
// }

// func TestMultilineCollapse(t *testing.T) {
// 	assert := testpredicate.NewAsserter(t)
// 	assert.That(true, pred.IsTrue())

// 	v := struct {
// 		values []int
// 	}{
// 		values: makeInts(1000000000000000000, 100),
// 	}

// 	s := value.Format(v)
// 	// fmt.Print(s)
// 	// assert.That(s, pred.Eq(""))

// 	lines := strings.Split(s, "\n")
// 	assert.That(lines, pred.Length(pred.GreaterThan(1)))
// 	assert.That(lines, pred.All(pred.Length(pred.Lt(80))))
// }
