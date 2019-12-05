package prettyprint

// Option ...
// type Option func(o *opts)

// type opts struct {
// 	width      int
// 	wrapPrefix string
// 	wrapSuffix string
// }

// func defaultOpts() *opts {
// 	return &opts{
// 		width:      80,
// 		wrapPrefix: "↩",
// 		wrapSuffix: "↪",
// 	}
// }

// // Width sets the maximum width for the formating of the values
// func Width(w int) Option {
// 	return func(o *opts) {
// 		o.width = w
// 	}
// }

// // WrapMarkers specify a pair or marker strings to be inserted around a line
// // wrap
// func WrapMarkers(before, after string) Option {
// 	return func(o *opts) {
// 		o.wrapPrefix = before
// 		o.wrapSuffix = after
// 	}
// }

// // // TruncateString ...
// // func TruncateString(l int) {
// // }

// // // TrimDepth ...
// // func TrimDepth(maxDepth int) {
// // }

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
