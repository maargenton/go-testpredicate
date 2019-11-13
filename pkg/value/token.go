package value

import (
	"io"
	"strings"
)

type token struct {
	level int
	str   string
	kvSep int

	trailing string
	sub      []token
}

func makeToken(s string) token {
	return token{
		str:   s,
		kvSep: scanKVSep(s),
	}
}

func (t *token) isOpening() bool {
	return strings.HasSuffix(t.str, "{")
}

func (t *token) isClosing() bool {
	return strings.HasPrefix(t.str, "}")
}

func (t *token) setLevel(l int) {
	t.level = l
	for i := range t.sub {
		t.sub[i].setLevel(l + 1)
	}
}

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

var indent = []byte{'\t'}
var newline = []byte{'\n'}

func (t *token) dump(w io.Writer) {
	for i := 0; i < t.level; i++ {
		w.Write(indent)
	}

	w.Write([]byte(t.str))

	if len(t.trailing) > 0 || len(t.sub) > 0 {
		w.Write(newline)

		for i, s := range t.sub {
			s.dump(w)
			if i != len(t.sub)-1 {
				w.Write(newline)
			}
		}

		if len(t.trailing) > 0 {
			w.Write(newline)
			for i := 0; i < t.level; i++ {
				w.Write(indent)
			}
			w.Write([]byte(t.trailing))
		}
	}
}

// ---------------------------------------------------------------------------

func dumpTokens(tokens []token, w io.Writer) {
	for _, t := range tokens {
		t.dump(w)
	}
}

func tokenize(str string, w int) []token {
	var tokens []token
	for len(str) > 0 {
		var t token
		t, str = nextToken(str)
		tokens = append(tokens, t)
	}

	return tokens
}

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
	for len(tokens) > 0 {
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

	return t, tokens
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
				next, os = nextToken(skipStruct(os))
				if next.str == "{" {
					ts += "...} {"
				} else {
					ts += "...}" + next.str
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

// ---------------------------------------------------------------------------

func isSpace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\r' || b == '\n'
}

func isAlnum(b byte) bool {
	return b >= '0' && b <= '9' ||
		b >= 'A' && b <= 'Z' ||
		b >= 'a' && b <= 'z'
}

func isStruct(s string) bool {
	i := strings.Index(s, "struct")
	if i < 0 ||
		i > 0 && isAlnum(s[i-1]) ||
		i+6 < len(s) && isAlnum(s[i+6]) {
		return false
	}
	return true
}

func skipStruct(s string) string {
	var l = len(s)
	var nested = 1

	for i := 0; i < l; i++ {
		if s[i] == '{' {
			nested++
		}
		if s[i] == '}' {
			nested--
			if nested == 0 {
				return s[i+1:]
			}
		}
	}
	return ""
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
