package require_test

import (
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/require"
)

func TestExample(t *testing.T) {
	v := 123
	require.That(t, v).ToString().Length().Eq(3)
	require.That(t, v,
		require.Context{Name: "double", Value: v * 2},
	).ToString().Length().Eq(3)
}
