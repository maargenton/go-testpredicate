package prettyprint

import (
	"fmt"
	"strings"
)

// Option ...
// type Option func(o *opts)

// Formatter contains the configuration
type Formatter struct {
	Width      int
	WrapPrefix string
	WrapSuffix string

	IndentWidth int
	IndentStr   string
	NewlineStr  string
}

// New return a new pretty-printer that can be customized and used locally
func New() *Formatter {
	return &Formatter{
		Width:       80,
		WrapPrefix:  "↩",
		WrapSuffix:  "↪",
		IndentWidth: 4,
		IndentStr:   "\t",
		NewlineStr:  "\n",
	}
}

// Default is the default pretty-printer with globally chared config
var Default *Formatter = New()

// FormatValue return a formated value using the default shared pretty-printer
func FormatValue(v interface{}) string {
	return Default.FormatValue(v)
}

// FormatValue return the value formated according to the local settings
func (f *Formatter) FormatValue(v interface{}) string {
	str := fmt.Sprintf("%#v", v)
	tokenList := tokenize(str)
	tokens := buildTokenTree(tokenList)

	f.collapseLeaves(tokens)
	f.alignValues(tokens)
	f.applyToTokens(tokens, f.wrapToken())

	var buf strings.Builder
	f.writeTokens(&buf, tokens)

	return buf.String()
}

// collapseLeaves either folds leaf sub-tokens into the parent or group them
// into lines that lines that fir into the prescribed width.
func (f *Formatter) collapseLeaves(tokens []token) {
	for i := range tokens {
		if tokens[i].isCollapsable() {
			f.tryCollapseToken(&tokens[i])
		} else {
			f.collapseLeaves(tokens[i].sub)
		}
	}
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func (f *Formatter) tryCollapseToken(t *token) {
	if !t.isCollapsable() {
		return
	}

	var availableWidth = f.Width - 4*t.level
	var baseWidth = len(t.str) + len(t.trailing)
	var maxWidth = 0
	var totalWidth = 0
	for _, s := range t.sub {
		l := len(s.str) + 1
		totalWidth += l
		maxWidth = max(maxWidth, l)
	}
	if baseWidth+totalWidth < availableWidth {
		f.collapseToken(t)
	} else {
		var avgWidth = totalWidth / len(t.sub)
		if avgWidth < availableWidth/3 && maxWidth < availableWidth/2 {
			f.collapseTokenMultiline(t)
		}
	}
}

func (f *Formatter) collapseToken(t *token) {
	var buf strings.Builder
	buf.WriteString(t.str)
	for j := range t.sub {
		buf.WriteByte(' ')
		buf.WriteString(t.sub[j].str)
	}
	str := buf.String()
	if strings.HasSuffix(str, ",") {
		str = str[:len(str)-1]
	}
	t.str = str + " " + t.trailing
	t.trailing = ""
	t.subCount = len(t.sub)
	t.sub = []token{}
}

func (f *Formatter) collapseTokenMultiline(t *token) {
	var buf strings.Builder
	var sub []token
	var availableWidth = f.Width - 4*(t.level+1)

	for _, s := range t.sub {
		if buf.Len()+len(s.str)+1 > availableWidth {
			sub = append(sub, makeToken(buf.String()))
			buf.Reset()
		}
		if buf.Len() > 0 {
			buf.WriteByte(' ')
		}
		buf.WriteString(s.str)
	}
	if buf.Len() > 0 {
		sub = append(sub, makeToken(buf.String()))
		buf.Reset()
	}
	t.subCount = len(t.sub)
	t.sub = sub

	truncateSubTokens(t, 10)
}

func truncateSubTokens(t *token, m int) {
	sub := make([]token, 0)
	if len(t.sub) > m {
		n2 := m / 2
		n1 := m - n2

		sub1 := t.sub[:n1]
		sub2 := t.sub[len(t.sub)-n2:]
		sub = append(sub, sub1...)
		sub = append(sub, makeToken("...,"))
		sub = append(sub, sub2...)
	} else {
		sub = t.sub
	}

	if t.subCount >= 7 || len(sub) > 3 {
		cnt := fmt.Sprintf("// len() = %v", t.subCount)
		sub = append(sub, makeToken(cnt))
	}
	t.sub = sub
	t.setLevel(t.level)
}

// alignValues aligns the values of consecutive key-value tokens
func (f *Formatter) alignValues(tokens []token) {

}

// applyToTokens applies a transformation recursively to all tokens
func (f *Formatter) applyToTokens(tokens []token, op func(t *token)) {

}

// wrapToken fits long token into the prescribed width, splitting it in multiple
// lines if needed, and surrounding each inserted line break with configured
// markers
func (f *Formatter) wrapToken() func(t *token) {
	return func(t *token) {}
}

// writeTokens dump a processed token tree into the string builder
func (f *Formatter) writeTokens(buf *strings.Builder, tokens []token) {
	for _, t := range tokens {
		f.writeIndent(buf, t.level)
		buf.WriteString(t.str)
		buf.WriteString(f.NewlineStr)

		f.writeTokens(buf, t.sub)

		if len(t.trailing) > 0 {
			f.writeIndent(buf, t.level)
			buf.WriteString(t.trailing)
			buf.WriteString(f.NewlineStr)
		}
	}
}

func (f *Formatter) writeIndent(buf *strings.Builder, level int) {
	for i := 0; i < level; i++ {
		buf.WriteString(f.IndentStr)
	}
}

// var indent = []byte{'\t'}
// var newline = []byte{'\n'}

// func (t *token) dump(w io.Writer) {
// 	for i := 0; i < t.level; i++ {
// 		w.Write(indent)
// 	}
// 	w.Write([]byte(t.str))
// 	w.Write(newline)

// 	for _, s := range t.sub {
// 		s.dump(w)
// 	}

// 	if len(t.trailing) > 0 {
// 		for i := 0; i < t.level; i++ {
// 			w.Write(indent)
// 		}
// 		w.Write([]byte(t.trailing))
// 		w.Write(newline)
// 	}
// }

// func dumpTokens(tokens []token, w io.Writer) {
// 	for _, t := range tokens {
// 		t.dump(w)
// 	}
// }

// // Format ....
// func Format(v interface{}, options ...Option) string {

// 	opts := defaultOpts()
// 	for _, o := range options {
// 		o(opts)
// 	}

// 	str := fmt.Sprintf("%#v", v)
// 	tokenList := tokenize(str)
// 	tokenTree := buildTokenTree(tokenList)

// 	collapseLeaves(tokenTree, 80)
// 	alignValues(tokenTree)
// 	applyToTokens(tokenTree, wrapToken(opts))

// 	var buf strings.Builder
// 	var w = makeTokenWriter(&buf)
// 	w.writeTokens(tokenTree)

// 	return buf.String()
// }

// // ---------------------------------------------------------------------------
// // Collapse leaf nodes based on formating width
// // ---------------------------------------------------------------------------

// func collapseLeaves(tokens []token, w int) {
// 	// outer_loop:
// 	for i := range tokens {
// 		if tokens[i].isCollapsable() {
// 			tryCollapseToken(&tokens[i], w)
// 		} else {
// 			collapseLeaves(tokens[i].sub, w)
// 		}
// 	}
// }

// func max(x, y int) int {
// 	if x < y {
// 		return y
// 	}
// 	return x
// }

// func tryCollapseToken(t *token, w int) {
// 	if !t.isCollapsable() {
// 		return
// 	}

// 	var availableWidth = w - 4*t.level
// 	var baseWidth = len(t.str) + len(t.trailing)
// 	var maxWidth = 0
// 	var totalWidth = 0
// 	for _, s := range t.sub {
// 		l := len(s.str) + 1
// 		totalWidth += l
// 		maxWidth = max(maxWidth, l)
// 	}
// 	if baseWidth+totalWidth < availableWidth {
// 		collapseToken(t, w)
// 	} else {
// 		var avgWidth = totalWidth / len(t.sub)
// 		if avgWidth < availableWidth/3 && maxWidth < availableWidth/2 {
// 			collapseTokenMultiline(t, w)
// 		}
// 	}
// }

// func collapseToken(t *token, w int) {
// 	var buf strings.Builder
// 	buf.WriteString(t.str)
// 	for j := range t.sub {
// 		buf.WriteByte(' ')
// 		buf.WriteString(t.sub[j].str)
// 	}
// 	str := buf.String()
// 	if strings.HasSuffix(str, ",") {
// 		str = str[:len(str)-1]
// 	}
// 	t.str = str + " " + t.trailing
// 	t.trailing = ""
// 	t.subCount = len(t.sub)
// 	t.sub = []token{}
// }

// func collapseTokenMultiline(t *token, w int) {
// 	var buf strings.Builder
// 	var sub []token
// 	var availableWidth = w - 4*(t.level+1)

// 	for _, s := range t.sub {
// 		if buf.Len()+len(s.str)+1 > availableWidth {
// 			sub = append(sub, makeToken(buf.String()))
// 			buf.Reset()
// 		}
// 		if buf.Len() > 0 {
// 			buf.WriteByte(' ')
// 		}
// 		buf.WriteString(s.str)
// 	}
// 	if buf.Len() > 0 {
// 		sub = append(sub, makeToken(buf.String()))
// 		buf.Reset()
// 	}
// 	t.subCount = len(t.sub)
// 	t.sub = sub

// 	truncateSubTokens(t, 10)
// }

// func truncateSubTokens(t *token, m int) {
// 	sub := make([]token, 0)
// 	if len(t.sub) > m {
// 		n2 := m / 2
// 		n1 := m - n2

// 		sub1 := t.sub[:n1]
// 		sub2 := t.sub[len(t.sub)-n2:]
// 		sub = append(sub, sub1...)
// 		sub = append(sub, makeToken("...,"))
// 		sub = append(sub, sub2...)
// 	} else {
// 		sub = t.sub
// 	}

// 	if t.subCount >= 7 || len(sub) > 3 {
// 		cnt := fmt.Sprintf("// len() = %v", t.subCount)
// 		sub = append(sub, makeToken(cnt))
// 	}
// 	t.sub = sub
// 	t.setLevel(t.level)
// }

// // ---------------------------------------------------------------------------
// // Key / Value pairs alignment
// // ---------------------------------------------------------------------------

// type chunkBounds struct {
// 	s, e int
// }

// // chunkBy splits a slice into contiguous ranges of similar property
// func chunkBy(l int, f func(int) int) []chunkBounds {
// 	var bb []chunkBounds

// 	if l == 0 {
// 		return bb
// 	}

// 	var s = 0
// 	var v0 = f(0)

// 	for i := 1; i < l; i++ {
// 		v := f(i)
// 		if v != v0 {
// 			bb = append(bb, chunkBounds{s: s, e: i})
// 			s = i
// 			v0 = v
// 		}
// 	}
// 	if s < l {
// 		bb = append(bb, chunkBounds{s: s, e: l})
// 	}
// 	return bb
// }

// func alignValues(tokens []token) {
// 	for i := range tokens {
// 		alignValues(tokens[i].sub)

// 		idx := 0
// 		splitBeforeNested := func(j int) int {
// 			if len(tokens[i].sub[j].sub) > 0 {
// 				idx++
// 				return idx - 1
// 			}
// 			return idx
// 		}

// 		for _, b := range chunkBy(len(tokens[i].sub), splitBeforeNested) {
// 			alignTokenValues(tokens[i].sub[b.s:b.e])
// 		}
// 	}
// }

// func alignTokenValues(tokens []token) {
// 	c := -1
// 	for i := range tokens {
// 		if tokens[i].kvSep > c {
// 			c = tokens[i].kvSep
// 		}
// 	}

// 	for i := range tokens {
// 		tokens[i].alignValue(c + 2)
// 	}
// }

// // ---------------------------------------------------------------------------
// // Split long tokens
// // ---------------------------------------------------------------------------

// func applyToTokens(tokens []token, f func(t *token)) {
// 	for i := range tokens {
// 		if len(tokens[i].sub) > 0 {
// 			applyToTokens(tokens[i].sub, f)
// 		}
// 		f(&tokens[i])
// 	}
// }

// func wrapToken(opts *opts) func(t *token) {
// 	return func(t *token) {
// 		w := opts.width - 4*t.level
// 		if len(t.str) > w {
// 			lines := wrapString(t.str, w)
// 			t.str = lines[0] + opts.wrapPrefix
// 			for i, l := range lines[1:] {
// 				var str string
// 				if i == len(lines)-2 {
// 					str = opts.wrapSuffix + l
// 				} else {
// 					str = opts.wrapSuffix + l + opts.wrapPrefix
// 				}
// 				st := makeToken(str)
// 				st.level = t.level + 1
// 				t.sub = append(t.sub, st)
// 			}
// 		}
// 	}
// }

// func nextBreakPoint(s string, i int) int {
// 	var l = len(s)
// 	for i < l-1 {
// 		if isSpace(s[i]) && !isSpace(s[i+1]) {
// 			return i + 1
// 		}
// 		i++
// 		if i < l-2 && s[i] == '\\' {
// 			return i + 2
// 		}
// 	}
// 	return l
// }

// func wrapString(s string, w int) []string {
// 	var result []string
// 	var i, j int
// 	var l = len(s)
// 	var firstLine = true

// 	for j < l {
// 		k := nextBreakPoint(s, j)
// 		if k-i < w+1 {
// 			j = k
// 		} else if j-i < w/2 || k-j > w/2 {
// 			j = i + w - 1
// 			result = append(result, s[i:j])
// 			i = j
// 			j = k
// 		} else {
// 			result = append(result, s[i:j])
// 			i = j
// 			j = k
// 		}
// 		if firstLine && len(result) > 0 {
// 			firstLine = false
// 			w -= 4
// 		}
// 	}
// 	if i != j {
// 		result = append(result, s[i:j])
// 	}

// 	return result
// }
