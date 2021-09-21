package bdd

type state struct {
	current int
	count   int
}

// tracker defines both a root level and sub-level bifurcated evaluation
// context. A newly initialized tracker is initially invalid and can only be
// accessed after a first successful call to `Next()`.
type tracker struct {
	state *[]state
	level int
}

// Next either initializes the tracker and returns true, or advances the tracker
// to the next branch and returns true. It returns false after the last branch
// has been evaluated.
func (t *tracker) Next() bool {
	if t.state == nil {
		s := make([]state, 1)
		t.state = &s
		return true
	}
	for i := len(*t.state) - 1; i >= 0; i-- {
		(*t.state)[i].current++
		if (*t.state)[i].current < (*t.state)[i].count {
			t.resetCounters(i)
			return true
		}
	}
	return false
}

func (t *tracker) resetCounters(k int) {
	for i := range *t.state {
		(*t.state)[i].count = 0
		if i > k {
			(*t.state)[i].current = 0
		}
	}
}

// SubTracker returns a new tracker object, linked to the current tracker, that
// captures the next nested level i n the bifurcated evaluation tree.
func (t *tracker) SubTracker() *tracker {
	if t.state == nil {
		panic("tracker.SubTracker() called before interation started with tracker.Next()")
	}

	for len(*t.state) < t.level+2 {
		*t.state = append(*t.state, state{})
	}
	return &tracker{
		state: t.state,
		level: t.level + 1,
	}
}

// Active increments both the current index and count at the current level and
// returns true is the current branch is active and its evaluation should
// continue.
func (t *tracker) Active() bool {
	if t.state == nil {
		panic("tracker.Active() called before interation started with tracker.Next()")
	}

	active := (*t.state)[t.level].count == (*t.state)[t.level].current
	(*(t.state))[t.level].count++
	return active
}
