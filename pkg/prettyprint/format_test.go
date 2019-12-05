package prettyprint_test

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/maargenton/go-testpredicate"
	"github.com/maargenton/go-testpredicate/pkg/prettyprint"
	"github.com/maargenton/go-testpredicate/pkg/value"
	"github.com/maargenton/go-testpredicate/pred"
)

var anonymousStruct = &struct {
	name    string
	path    string
	options []string
	nested  []struct {
		name  string
		value float64
	}
	index   []int
	details string
}{
	name:    "my name",
	path:    "/aaa/bbb/ccc/d",
	options: []string{"readable", "aaa{", "bbb}", "build-target"},
	nested: []struct {
		name  string
		value float64
	}{
		struct {
			name  string
			value float64
		}{
			name:  "aaa",
			value: 123.456,
		},
	},
	index:   []int{1, 5, 2, 3, ')', 12, 1, 5, 2, 3, ')', 12, 1, 5, 2, 3, ')', 12, 1, 5, 2, 3, ')', 12, 1, 5, 2, 3, ')', 12},
	details: "MakeBoolPredicate wraps a predicate function\nreturning bool into a predicateinterface. Any error returned from the function is interpreted as an invalid evaluation.xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
}

func TestFormatValueWithStruct(t *testing.T) {
	assert := testpredicate.NewAsserter(t)
	assert.That(true, pred.Eq(true))

	s := prettyprint.FormatValue(anonymousStruct)
	_ = s
	fmt.Print(s)
	// assert.That(s, pred.Eq(""))

	lines := strings.Split(s, "\n")
	assert.That(lines, pred.Length(pred.GreaterThan(1)))
	assert.That(lines, pred.All(pred.Length(pred.Lt(82))))
}

func TestFormatString(t *testing.T) {
	assert := testpredicate.NewAsserter(t)
	assert.That(true, pred.Eq(true))

	str := "MakeBoolPredicate wraps a predicate function\nreturning bool into a predicateinterface. Any error returned from the function is interpreted as an invalid evaluation."
	s := value.Format(str)
	lines := strings.Split(s, "\n")
	assert.That(lines, pred.Length(pred.GreaterThan(1)))
	assert.That(lines, pred.All(pred.Length(pred.Le(82))))
}

func makeInts(m, n int) []int {
	var result = make([]int, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, rand.Intn(m))
	}

	return result
}

func TestMultilineCollapse(t *testing.T) {
	assert := testpredicate.NewAsserter(t)
	assert.That(true, pred.IsTrue())

	v := struct {
		values []int
	}{
		values: makeInts(1000000000000000000, 100),
	}

	s := value.Format(v)
	// fmt.Print(s)
	// assert.That(s, pred.Eq(""))

	lines := strings.Split(s, "\n")
	assert.That(lines, pred.Length(pred.GreaterThan(1)))
	assert.That(lines, pred.All(pred.Length(pred.Lt(80))))
}
