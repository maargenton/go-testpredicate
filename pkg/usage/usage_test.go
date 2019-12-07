package usage_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate"
	"github.com/maargenton/go-testpredicate/pred"
)

var skip = true

func TestEmpty(t *testing.T) {
	if skip {
		t.Skip()
	}
	assert := testpredicate.NewAsserter(t)
	assert.That(nil, pred.IsNil())

	v := []int{1, 2, 3}
	assert.That(v, pred.IsEmpty())
}

func TestContains(t *testing.T) {
	if skip {
		t.Skip()
	}
	assert := testpredicate.NewAsserter(t)
	assert.That(nil, pred.IsNil())

	v := []int{1, 2, 3, 4, 5, 6, 7, 8}
	w := []int{3, 4, 6}
	assert.That(v, pred.Contains(w))
}

func TestContainsWrongType(t *testing.T) {
	if skip {
		t.Skip()
	}
	assert := testpredicate.NewAsserter(t)
	assert.That(nil, pred.IsNil())

	v := []int{1, 2, 3, 4, 5, 6, 7, 8}
	w := struct{ v []int }{v: []int{3, 4, 6}}
	assert.That(v, pred.Contains(w))
}

func TestLength(t *testing.T) {
	if skip {
		t.Skip()
	}
	assert := testpredicate.NewAsserter(t)
	assert.That(nil, pred.IsNil())

	v := []int{1, 2, 3, 4, 5, 6, 7, 8}
	assert.That(v, pred.Length(pred.Lt(4)))
}
