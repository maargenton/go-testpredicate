package p_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/predicate"
)

type predicateExpectation struct {
	value        interface{}
	result       predicate.Result
	descMatchers []string
	errMatchers  []string
}

func validatePredicate(t *testing.T, p predicate.T, exp *predicateExpectation) {
	t.Helper()
	if p == nil {
		t.Errorf("\npredicate should not be nil")
	}

	desc := fmt.Sprintf("%v", p)
	for _, m := range exp.descMatchers {

		// matched, rerr := regexp.MatchString(m, desc)
		matched, rerr := checkConatinsOrRegexpMatch(desc, m)
		if rerr != nil {
			t.Errorf("\nfailed to match description with %v,\nregexp error: %v", m, rerr)
		} else if !matched {
			t.Errorf("\nfailed to match description with %v,\ndescription: %v", m, desc)
		}
	}

	r, err := p.Evaluate(exp.value)
	if r != exp.result {
		t.Errorf(
			"\nexpected predicate result to be %v, \nwas %v",
			exp.result, r)
	}
	n := len(exp.errMatchers)
	if n != 0 && err == nil {
		t.Errorf("\nexpected predicate to return an error")
	} else if n == 0 && err != nil {
		t.Errorf("\nexpected predicate to return no error,\nerror: %v", err)
	} else {
		for _, m := range exp.errMatchers {

			matched, rerr := checkConatinsOrRegexpMatch(err.Error(), m)
			if rerr != nil {
				t.Errorf("\nfailed to match error with '%v',\nregexp error: %v", m, rerr)
			} else if !matched {
				t.Errorf("\nfailed to match error with '%v',\npredicate error: %v", m, err)
			}
		}
	}
}

func checkConatinsOrRegexpMatch(s, match string) (bool, error) {
	if strings.HasPrefix(match, "/") && strings.HasSuffix(match, "/") {
		re := "(?s)" + match[1:len(match)-1]
		return regexp.MatchString(re, s)
	}

	return strings.Contains(s, match), nil
}
