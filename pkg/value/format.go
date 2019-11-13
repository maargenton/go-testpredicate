package value

import (
	"fmt"
	"strings"
)

// Todo :
// - String wrapping, force wrap when no reasonable break point
// - multi-line collapse of short nested values in groups of 10
//   - condition: all groups of 10 fit in line
// - truncate list of sub based on line count after collapse (10 max)
//   - end with: ...(230 total)

// Option ...
type Option func(o *opts)

type opts struct {
	width      int
	wrapPrefix string
	wrapSuffix string
}

func defaultOpts() *opts {
	return &opts{
		width:      80,
		wrapPrefix: "↩",
		wrapSuffix: "↪",
	}
}

// Width sets the maximum width for the formating of the values
func Width(w int) Option {
	return func(o *opts) {
		o.width = w
	}
}

// WrapMarkers specify a pair or marker strings to be inserted around a line
// wrap
func WrapMarkers(before, after string) Option {
	return func(o *opts) {
		o.wrapPrefix = before
		o.wrapSuffix = after
	}
}

// // TruncateString ...
// func TruncateString(l int) {
// }

// // TrimDepth ...
// func TrimDepth(maxDepth int) {
// }

// Format ....
func Format(v interface{}, options ...Option) string {

	opts := defaultOpts()
	for _, o := range options {
		o(opts)
	}

	str := fmt.Sprintf("%#v", v)
	tokenList := tokenize(str, 80)
	tokenTree := buildTokenTree(tokenList)

	collapseLeaves(tokenTree, 80)
	alignValues(tokenTree)
	applyToTokens(tokenTree, wrapToken(opts))

	var buf strings.Builder
	dumpTokens(tokenTree, &buf)
	// for _, t := range tokenTree {
	// 	t.dump(&buf)
	// }

	return buf.String()
}

// ---------------------------------------------------------------------------
// Collapse leaf nodes based on formating width
// ---------------------------------------------------------------------------

func collapseLeaves(tokens []token, w int) {
outer_loop:
	for i := range tokens {
		var l = len(tokens[i].str) + len(tokens[i].trailing)
		if len(tokens[i].sub) == 0 {
			continue
		}

		for j := range tokens[i].sub {
			if len(tokens[i].sub[j].sub) > 0 || tokens[i].sub[j].kvSep != -1 {
				collapseLeaves(tokens[i].sub, w)
				continue outer_loop
			}
			l += len(tokens[i].sub[j].str) + 1
		}

		if l < w-4*tokens[i].level {
			var buf strings.Builder
			buf.WriteString(tokens[i].str)
			for j := range tokens[i].sub {
				buf.WriteByte(' ')
				buf.WriteString(tokens[i].sub[j].str)
			}
			str := buf.String()
			if strings.HasSuffix(str, ",") {
				str = str[:len(str)-1]
			}
			tokens[i].str = str + " " + tokens[i].trailing
			tokens[i].trailing = ""
			tokens[i].sub = []token{}
		}
	}
}

// ---------------------------------------------------------------------------
// Key / Value pairs alignment
// ---------------------------------------------------------------------------

type chunkBounds struct {
	s, e int
}

// chunkBy splits a slice into contiguous ranges of similar property
func chunkBy(l int, f func(int) int) []chunkBounds {
	var bb []chunkBounds

	if l == 0 {
		return bb
	}

	var s = 0
	var v0 = f(0)

	for i := 1; i < l; i++ {
		v := f(i)
		if v != v0 {
			bb = append(bb, chunkBounds{s: s, e: i})
			s = i
			v0 = v
		}
	}
	if s < l {
		bb = append(bb, chunkBounds{s: s, e: l})
	}
	return bb
}

func alignValues(tokens []token) {
	for i := range tokens {
		alignValues(tokens[i].sub)

		idx := 0
		splitBeforeNested := func(j int) int {
			if len(tokens[i].sub[j].sub) > 0 {
				idx++
				return idx - 1
			}
			return idx
		}

		for _, b := range chunkBy(len(tokens[i].sub), splitBeforeNested) {
			alignTokenValues(tokens[i].sub[b.s:b.e])
		}
	}
}

func alignTokenValues(tokens []token) {
	c := -1
	for i := range tokens {
		if tokens[i].kvSep > c {
			c = tokens[i].kvSep
		}
	}

	for i := range tokens {
		tokens[i].alignValue(c + 2)
	}
}

// ---------------------------------------------------------------------------
// Split long tokens
// ---------------------------------------------------------------------------

func applyToTokens(tokens []token, f func(t *token)) {

	for i := range tokens {
		if len(tokens[i].sub) > 0 {
			applyToTokens(tokens[i].sub, f)
		}
		f(&tokens[i])
	}
}

func wrapToken(opts *opts) func(t *token) {
	return func(t *token) {
		w := opts.width - 4*t.level
		if len(t.str) > w {
			lines := wrapString(t.str, w)
			t.str = lines[0] + opts.wrapPrefix
			for i, l := range lines[1:] {
				var str string
				if i == len(lines)-2 {
					str = opts.wrapSuffix + l
				} else {
					str = opts.wrapSuffix + l + opts.wrapPrefix
				}
				st := makeToken(str)
				st.level = t.level + 1
				t.sub = append(t.sub, st)
			}
		}
	}
}

func nextBreakPoint(s string, i int) int {
	var l = len(s)
	for i < l-1 {
		if isSpace(s[i]) && !isSpace(s[i+1]) {
			return i + 1
		}
		i++
		if i < l-2 && s[i] == '\\' {
			return i + 2
		}
	}
	return l
}

func wrapString(s string, w int) []string {
	var result []string
	var i, j int
	var l = len(s)

	for j < l {
		k := nextBreakPoint(s, j)
		if k-i < w+1 {
			j = k
		} else {
			result = append(result, s[i:j])
			i = j
			j = k
		}
	}
	if i != j {
		result = append(result, s[i:j])
	}

	return result
}
