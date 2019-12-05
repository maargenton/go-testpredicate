package prettyprint

import (
	"fmt"
	"testing"
)

// ---------------------------------------------------------------------------
// makeToken()

func TestMakeToken(t *testing.T) {
	tok := makeToken("123")

	if tok.str != "123" {
		t.Errorf("\nExpected tok.str to be '123', was '%v'", tok.str)
	}
	if tok.kvSep != -1 {
		t.Errorf("\nExpected tok.kvSep to be -1, was %v", tok.kvSep)
	}
}

func TestMakeTokenWithKV(t *testing.T) {
	tok := makeToken("value: 123")

	if tok.str != "value: 123" {
		t.Errorf("\nExpected tok.str to be 'value: 123', was '%v'", tok.str)
	}
	if tok.kvSep != 5 {
		t.Errorf("\nExpected tok.kvSep to be 6, was %v", tok.kvSep)
	}
}

func TestMakeTokenWithStringLiteral(t *testing.T) {
	tok := makeToken(`"value: 123"`)

	if tok.str != `"value: 123"` {
		t.Errorf("\nExpected tok.str to be 'value: 123', was '%v'", tok.str)
	}
	if tok.kvSep != -1 {
		t.Errorf("\nExpected tok.kvSep to be -1, was %v", tok.kvSep)
	}
}

// ---------------------------------------------------------------------------
// token.is...()

func TestIsOpeningIsClosing(t *testing.T) {
	tokOpen := makeToken("[]int {")
	tok := makeToken("value: 123")
	tokClose := makeToken("}, ")

	if !tokOpen.isOpening() {
		t.Errorf("\ntoken '%v' should be opening", tokOpen.str)
	}
	if tokOpen.isClosing() {
		t.Errorf("\ntoken '%v' should not be closing", tokOpen.str)
	}

	if tok.isOpening() {
		t.Errorf("\ntoken '%v' should not be opening", tok.str)
	}
	if tok.isClosing() {
		t.Errorf("\ntoken '%v' should not be closing", tok.str)
	}

	if tokClose.isOpening() {
		t.Errorf("\ntoken '%v' should not be opening", tokClose.str)
	}
	if !tokClose.isClosing() {
		t.Errorf("\ntoken '%v' should be closing", tokClose.str)
	}
}

func TestIsCollapsable(t *testing.T) {
	tok := makeToken("[]interface{} {")
	tok.trailing = "}"
	tok.sub = []token{
		makeToken("123,"),
		makeToken("456,"),
	}

	if !tok.isCollapsable() {
		t.Errorf("\ntoken '%v' should be collapsable", tok.str)
	}

	if tok.sub[0].isCollapsable() {
		t.Errorf("\nleaf token '%v' should be collapsable", tok.sub[0])
	}

	nested := makeToken("struct{...} {")
	nested.trailing = "}"
	nested.sub = []token{
		makeToken("value1: 123,"),
		makeToken("value2: 456,"),
	}

	if nested.isCollapsable() {
		t.Errorf("\ntoken '%v' should not be collapsable", tok.str)
	}

	tok.sub = append(tok.sub, nested)
	if nested.isCollapsable() {
		t.Errorf("\ntoken '%v' should no longer be collapsable", tok.str)
	}

}

// ---------------------------------------------------------------------------
// token.setLevel()

func TestSetLevel(t *testing.T) {
	tok := makeToken("[]interface{} {")
	tok.trailing = "}"
	tok.sub = []token{
		makeToken("123,"),
		makeToken("456,"),
	}

	tok.setLevel(2)

	if tok.level != 2 {
		t.Errorf("\ntoken '%v' should be level 2, was level %v", tok.str, tok.level)
	}
	for _, tok := range tok.sub {
		if tok.level != 3 {
			t.Errorf("\nsub-token '%v' should be level 3, was level %v", tok.str, tok.level)
		}

	}
}

// ---------------------------------------------------------------------------
// token.alignValue()

func TestAlignValue(t *testing.T) {
	tok := makeToken("value1: 123,")

	tok.alignValue(3)
	if tok.str != "value1: 123," {
		t.Errorf("\ntoken value should be at least one space away, was: '%v'", tok.str)
	}
	tok.alignValue(10)
	if tok.str != "value1:   123," {
		t.Errorf("\ntoken value should be aligned at 10, was: '%v'", tok.str)
	}
	tok.alignValue(3)
	if tok.str != "value1: 123," {
		t.Errorf("\nlarge alignments should be reversible, was: '%v'", tok.str)
	}

	tok2 := makeToken("123,")
	tok2.alignValue(20)
	if tok2.str != "123," {
		t.Errorf("\nexpected no change on non-kv tokens, was: '%v'", tok.str)
	}
}

// ---------------------------------------------------------------------------
// tokenize()

var structValue = struct {
	i      int
	s      string
	is     []int
	nested struct {
		i int
	}
}{ // token 0: struct{...} {
	i: 123,            // token 1
	s: "aaa\n\"bbb\"", // token 2
	is: []int{ // token 3
		1, 2, 3, // tokens 4, 5, 6
	}, // token 7
	// token 8: nested {...} {
	//token 9: i:0
	// token 10: } of nested
} // token 11

func TestTokenize(t *testing.T) {
	var s = fmt.Sprintf("%#v", structValue)
	var tokens = tokenize(s)

	if len(tokens) != 12 {
		t.Fatalf("\nexpected 12 token, got: '%v'", len(tokens))
	}
	if tokens[8].str != "nested:struct {...} {" {
		t.Fatalf("\nunexpected token[8]: '%v'", tokens[8].str)
	}
}

func TestTokenizeWithTruncatedStringLiteral(t *testing.T) {
	var s = `"aaa\n\"bbb\"`
	var tokens = tokenize(s)

	if l := len(tokens); l != 1 {
		t.Fatalf("\nexpected 1 token, got: '%v'", l)
	}
	if str := tokens[0].str; str != s {
		t.Fatalf("\nunexpected token[0]: '%v'", str)
	}
}

func TestTokenizeWithTruncatedStruct(t *testing.T) {
	var s = `struct { i int; s string; is []int`
	var tokens = tokenize(s)

	if l := len(tokens); l != 1 {
		t.Fatalf("\nexpected 1 token, got: '%v'", l)
	}
	if str := tokens[0].str; str != "struct {...}" {
		t.Fatalf("\nunexpected token[0]: '%v'", str)
	}
}

// ---------------------------------------------------------------------------
// buildTokenTree()

func TestBuildTokenTree(t *testing.T) {
	var s = fmt.Sprintf("%#v", structValue)
	var tokens = tokenize(s)
	var tree = buildTokenTree(tokens)

	if l := len(tree); l != 1 {
		t.Fatalf("\nexpected 1 tree root, got: '%v'", l)
	}
	if l := len(tree[0].sub); l != 4 {
		t.Fatalf("\nexpected 4 root branches, got: '%v'", l)
	}

	nested := tree[0].sub[3]
	if l := len(nested.sub); l != 1 {
		t.Fatalf("\nexpected 1 branch within nested struct, got: '%v'", l)
	}
}

func TestBuildTokenTreeWithPartialTokens(t *testing.T) {
	var s = fmt.Sprintf("%#v", structValue)
	var tokens = tokenize(s)
	var tree = buildTokenTree(tokens[:5])

	if l := len(tree); l != 1 {
		t.Fatalf("\nexpected 1 tree root, got: '%v'", l)
	}
	if l := len(tree[0].sub); l != 3 {
		t.Fatalf("\nexpected 3 root branches, got: '%v'", l)
	}
}

func TestBuildTokenTreeWithNoToken(t *testing.T) {
	var s = fmt.Sprintf("%#v", structValue)
	var tokens = tokenize(s)
	var tree = buildTokenTree(tokens[:0])

	if l := len(tree); l != 0 {
		t.Fatalf("\nexpected no root tokens, got: '%v'", l)
	}
}
