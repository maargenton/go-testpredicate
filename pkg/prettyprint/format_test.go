package prettyprint_test

import (
	"strings"
	"testing"

	"github.com/maargenton/go-testpredicate"
	"github.com/maargenton/go-testpredicate/pkg/prettyprint"
	"github.com/maargenton/go-testpredicate/pred"
)

func makeInts(n int) []int {
	var r []int
	for i := 0; i < n; i++ {
		r = append(r, i)
	}
	return r
}

func TestFormatValueWithShortListCollapsesValuesIntoASingleLine(t *testing.T) {
	assert := testpredicate.NewAsserter(t)
	assert.That(true, pred.Eq(true))

	v := makeInts(3)
	s := prettyprint.FormatValue(v)

	lines := strings.Split(s, "\n")
	assert.That(lines, pred.Length(pred.Eq(2)))
	assert.That(lines, pred.All(pred.Length(pred.Le(82))))
}

func TestFormatValueCollapsesNestedLists(t *testing.T) {
	assert := testpredicate.NewAsserter(t)
	assert.That(true, pred.Eq(true))

	v := struct{ i []int }{i: makeInts(3)}
	s := prettyprint.FormatValue(v)

	lines := strings.Split(s, "\n")
	assert.That(lines, pred.Length(pred.Eq(4)))
	assert.That(lines, pred.All(pred.Length(pred.Le(82))))
}

func TestFormatValueWithLongListCollapsesValuesIntoMultipleLines(t *testing.T) {
	assert := testpredicate.NewAsserter(t)
	assert.That(true, pred.Eq(true))

	v := makeInts(30)
	s := prettyprint.FormatValue(v)

	lines := strings.Split(s, "\n")
	assert.That(lines, pred.Length(pred.Eq(6)))
	assert.That(lines, pred.All(pred.Length(pred.Le(82))))
}

func TestFormatValueWithWeryLongListTruncatesCollapsedLines(t *testing.T) {
	assert := testpredicate.NewAsserter(t)
	assert.That(true, pred.Eq(true))

	v := makeInts(200)
	s := prettyprint.FormatValue(v)

	lines := strings.Split(s, "\n")
	assert.That(lines, pred.Length(pred.Eq(15)))
	assert.That(lines, pred.All(pred.Length(pred.Le(82))))
}

func TestFormatValueWithWeryLongListTruncatesCollapsedLinesToMaxWrapped(t *testing.T) {
	assert := testpredicate.NewAsserter(t)
	assert.That(true, pred.Eq(true))

	pp := prettyprint.New()
	pp.MaxWrapped = 5
	v := makeInts(200)
	s := pp.FormatValue(v)

	lines := strings.Split(s, "\n")
	assert.That(lines, pred.Length(pred.Eq(10)))
	assert.That(lines, pred.All(pred.Length(pred.Le(82))))
}

// ---------------------------------------------------------------------------
// Test alignment of values in key-value tokens

func TestFormatValueAlignsValues(t *testing.T) {
	assert := testpredicate.NewAsserter(t)
	assert.That(true, pred.Eq(true))

	v := struct {
		a          int
		bbbbbbbbbb int
	}{
		a:          1,
		bbbbbbbbbb: 2,
	}
	s := prettyprint.FormatValue(v)

	lines := strings.Split(s, "\n")
	assert.That(lines[1], pred.Eq("\ta:          1,"))
	assert.That(lines[2], pred.Eq("\tbbbbbbbbbb: 2,"))
}

func TestFormatValueAlignsValuesInChunks(t *testing.T) {
	assert := testpredicate.NewAsserter(t)
	assert.That(true, pred.Eq(true))

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
	assert.That(lines[1], pred.Eq("\ta:          1,"))
	assert.That(lines[2], pred.Eq("\tbbbbbbbbbb: 2,"))
	assert.That(lines[7], pred.Eq("\td:   4,"))
	assert.That(lines[8], pred.Eq("\teee: 5,"))
}

// ---------------------------------------------------------------------------
// Test long tokens are wrapped onto multiple lines

func TestFormatValueWrapsLongStringTokens(t *testing.T) {
	assert := testpredicate.NewAsserter(t)
	assert.That(true, pred.Eq(true))

	v := strings.Repeat("abcdefghi ", 20)
	s := prettyprint.FormatValue(v)

	lines := strings.Split(s, "\n")
	assert.That(lines, pred.Length(pred.Gt(1)))
	assert.That(lines, pred.All(pred.Length(pred.Le(82))))
}

func TestFormatValueForceWrapsLongStringTokens(t *testing.T) {
	assert := testpredicate.NewAsserter(t)
	assert.That(true, pred.Eq(true))

	v := strings.Repeat("abcdefghi ", 20) +
		strings.Repeat("_aaabbb_", 20) +
		strings.Repeat("abcdefghi ", 20)

	s := prettyprint.FormatValue(v)
	lines := strings.Split(s, "\n")
	assert.That(lines, pred.Length(pred.Eq(9)))
	assert.That(lines, pred.All(pred.Length(pred.Le(82))))
}

func TestFormatValueCanBreakOnEscapedCharacters(t *testing.T) {
	assert := testpredicate.NewAsserter(t)
	assert.That(true, pred.Eq(true))

	v := strings.Repeat("abcdefghi\n", 20)
	s := prettyprint.FormatValue(v)
	// fmt.Println(s)
	// t.Fail()
	lines := strings.Split(s, "\n")
	assert.That(lines, pred.Length(pred.Eq(5)))
	assert.That(lines, pred.All(pred.Length(pred.Le(82))))
	assert.That(lines[:3], pred.All(pred.EndsWith("\\n↩")))
}

func TestFormatValueBreakAlignment(t *testing.T) {
	assert := testpredicate.NewAsserter(t)
	assert.That(true, pred.Eq(true))

	v := strings.Repeat("a bb ccc dddd eeeee ffffff ggggg hhhh iii jj k ", 80)
	s := prettyprint.FormatValue(v)
	// fmt.Println(s)
	// t.Fail()
	lines := strings.Split(s, "\n")
	assert.That(lines, pred.All(pred.Length(pred.Le(82))))
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
