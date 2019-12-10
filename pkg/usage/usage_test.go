package usage_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/asserter"
	"github.com/maargenton/go-testpredicate/pkg/p"
)

var skip = true

func TestEmpty(t *testing.T) {
	if skip {
		t.Skip()
	}
	assert := asserter.New(t)
	assert.That(nil, p.IsNil())

	v := []int{1, 2, 3}
	assert.That(v, p.IsEmpty())
}

func TestContains(t *testing.T) {
	if skip {
		t.Skip()
	}
	assert := asserter.New(t)
	assert.That(nil, p.IsNil())

	v := []int{1, 2, 3, 4, 5, 6, 7, 8}
	w := []int{3, 4, 6}
	assert.That(v, p.Contains(w))
}

func TestContainsWrongType(t *testing.T) {
	if skip {
		t.Skip()
	}
	assert := asserter.New(t)
	assert.That(nil, p.IsNil())

	v := []int{1, 2, 3, 4, 5, 6, 7, 8}
	w := struct{ v []int }{v: []int{3, 4, 6}}
	assert.That(v, p.Contains(w))
}

func TestLength(t *testing.T) {
	if skip {
		t.Skip()
	}
	assert := asserter.New(t)
	assert.That(nil, p.IsNil())

	v := []int{1, 2, 3, 4, 5, 6, 7, 8}
	assert.That(v, p.Length(p.Lt(4)))
}

func TestError(t *testing.T) {
	if skip {
		t.Skip()
	}
	assert := asserter.New(t)

	_, err := strconv.ParseInt("-zzt", 10, 0)
	err = fmt.Errorf("failed to parse value for 'limit' parameter, %w", err)
	assert.That(err, p.IsNoError())
	assert.That(err, p.IsError(strconv.ErrSyntax))
}
