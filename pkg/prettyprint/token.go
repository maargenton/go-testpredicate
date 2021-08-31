package prettyprint

import (
	"strings"
)

// token represent one parsed item of the printed value
type token struct {
	level int
	str   string
	kvSep int

	trailing string
	sub      []token
	subCount int // Original # of sub before collapse
}

// makeToken creates a new token from a string and scans for its key-value
// separator if any
func makeToken(s string) token {
	return token{
		str:   s,
		kvSep: scanKVSep(s),
	}
}

func scanKVSep(s string) int {
	for i, c := range s {
		if c == ':' {
			return i
		}
		if c == '"' {
			return -1
		}
	}
	return -1
}

// isOpening is true if the token is the beginning of a nested sequence of sub
// tokens
func (t *token) isOpening() bool {
	return strings.HasSuffix(t.str, "{")
}

// isClosing is true if the token is the end of a nested sequence of sub tokens
func (t *token) isClosing() bool {
	return strings.HasPrefix(t.str, "}")
}

// isCollapsable if token contains sub-tokens that are plain values, i.e.
// neither key-value pairs not containers with nested tokens
func (t *token) isCollapsable() bool {
	if len(t.sub) == 0 {
		return false
	}
	for _, s := range t.sub {
		if len(s.sub) > 0 || s.kvSep != -1 {
			return false
		}
	}
	return true
}

// setLevel assign a level to a token and depper level to each sub-token
// recursively
func (t *token) setLevel(l int) {
	t.level = l
	for i := range t.sub {
		t.sub[i].setLevel(l + 1)
	}
}

// alignValue aligns the value part of the token to the given column
func (t *token) alignValue(c int) {
	if t.kvSep < 0 {
		return
	}

	i := t.kvSep + 1
	j := scanSkipSpaces(t.str, i)
	k := t.str[:i]
	v := t.str[j:]
	pad := c - i
	if pad < 0 {
		pad = 1
	}
	t.str = k + strings.Repeat(" ", pad) + v
}

// ---------------------------------------------------------------------------

// tokenize parses the result of a '%#v' formatting and generates a list of
// tokens that represent each element of the string
func tokenize(str string) []token {
	var tokens []token
	for len(str) > 0 {
		var t token
		t, str = nextToken(str)
		tokens = append(tokens, t)
	}

	return tokens
}

func nextToken(is string) (t token, os string) {
	var l = len(is)
	var i = scanSkipSpaces(is, 0)
	var j = i
	for j < l {
		if is[j] == '"' {
			j = scanQuotedString(is, j+1)
		} else if is[j] == ',' || is[j] == '{' {
			ts := is[i : j+1]
			os = is[j+1:]
			if isStruct(ts) {
				var next token
				oss, empty := skipStruct(os)
				next, os = nextToken(oss)
				if empty {
					ts += "} " + next.str
				} else {
					ts += "...} " + next.str
				}
			} else if isInterface(ts) {
				var next token
				oss, empty := skipStruct(os)
				next, os = nextToken(oss)
				if empty {
					t = next
					return
				} else {
					ts += "...} " + next.str
				}
			}
			t = makeToken(ts)
			return
		} else if is[j] == '}' {
			if j == 0 {
				if l > 1 {
					if is[1] == ',' {
						return makeToken("},"), is[2:]
					}
					return makeToken("},"), is[1:]
				}
				return makeToken("}"), is[j+1:]
			}
			return makeToken(is[i:j] + ","), is[j:]
		} else {
			j++
		}
	}
	return makeToken(is[i:]), ""
}

func isSpace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\r' || b == '\n'
}

func isAlnum(b byte) bool {
	return b >= '0' && b <= '9' ||
		b >= 'A' && b <= 'Z' ||
		b >= 'a' && b <= 'z'
}

func containsWord(s, word string) bool {
	l := len(word)
	i := strings.Index(s, word)
	if i < 0 ||
		i > 0 && isAlnum(s[i-1]) ||
		i+l < len(s) && isAlnum(s[i+l]) {
		return false
	}
	return true
}
func isStruct(s string) bool {
	return containsWord(s, "struct")
}

func isInterface(s string) bool {
	return containsWord(s, "interface")
}

func skipStruct(s string) (string, bool) {
	var l = len(s)
	var nested = 1
	var count = 0

	for i := 0; i < l; i++ {
		if s[i] == '{' {
			nested++
		}
		if s[i] == '}' {
			nested--
			if nested == 0 {
				return s[i+1:], count == 0
			}
		}
		if s[i] != ' ' && s[i] != '\t' {
			count++
		}
	}
	return "", count == 0
}

func scanSkipSpaces(s string, i int) int {
	var l = len(s)
	for i < l && isSpace(s[i]) {
		i++
	}
	return i
}

func scanQuotedString(s string, i int) int {
	l := len(s)

	for i < l && s[i] != '"' {
		if s[i] == '\\' {
			i++
		}
		i++
	}
	if i < l {
		return i + 1
	}
	return l
}

// ---------------------------------------------------------------------------

// buildTokenTree takes a flat list of tokens from tokenize() and rebuilds the
// intrinsic hierarchy of tokens and nested tokens.
func buildTokenTree(tokens []token) []token {
	var result []token
	var t token

	for len(tokens) > 0 {
		t, tokens = buildTokenBranch(tokens)
		result = append(result, t)
	}
	for i := range result {
		result[i].setLevel(0)
	}
	return result
}

func buildTokenBranch(tokens []token) (token, []token) {
	var t token
	t, tokens = tokens[0], tokens[1:]
	if !t.isOpening() {
		return t, tokens
	}

	t0 := t
	for len(tokens) > 0 {
		t, tokens = buildTokenBranch(tokens)
		if t.isClosing() {
			t0.trailing = t.str
			return t0, tokens
		}
		t0.sub = append(t0.sub, t)
	}
	return t0, tokens
}
