package value_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/maargenton/go-testpredicate"
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
	index:   []int{1, 5, 2, 3, ')', 12},
	details: "MakeBoolPredicate wraps a predicate function\nreturning bool into a predicateinterface. Any error returned from the function is interpreted as an invalid evaluation.",
}

func TestXXX(t *testing.T) {
	assert := testpredicate.NewAsserter(t)
	assert.That(true, pred.Eq(true))

	s := value.Format(anonymousStruct)
	fmt.Println(s)
	assert.That(s, pred.Eq(""))

}

func TestFormatString(t *testing.T) {
	assert := testpredicate.NewAsserter(t)
	assert.That(true, pred.Eq(true))

	str := "MakeBoolPredicate wraps a predicate function\nreturning bool into a predicateinterface. Any error returned from the function is interpreted as an invalid evaluation."
	s := value.Format(str)
	lines := strings.Split(s, "\n")
	assert.That(lines, pred.Length(pred.GreaterThan(1)))
	assert.That(lines, pred.All(pred.Length(pred.Lt(80))))
}
