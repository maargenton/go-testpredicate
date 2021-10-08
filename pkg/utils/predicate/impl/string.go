package impl

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/maargenton/go-testpredicate/pkg/utils/predicate"
)

// (desc string, f PredicateFunc)

// Matches tests if a string matches a regular expression
func Matches(re string) (desc string, f predicate.PredicateFunc) {
	desc = fmt.Sprintf("{} =~ /%v/", re)
	f = func(v interface{}) (r bool, ctx []predicate.ContextValue, err error) {
		s, ok := v.(string)
		if !ok {
			return false, nil, fmt.Errorf(
				"value of type '%T' cannot be matched against a regexp", v)
		}
		m, err := regexp.MatchString(re, s)
		if err != nil {
			return false, nil, fmt.Errorf("failed to compile regexp: %w", err)
		}
		return m, nil, nil
	}
	return
}

// ---------------------------------------------------------------------------
// Transformation predicates on strings or producing strings

// ToString is a transformation predicate that converts any value to a string
// representation using `%v` formatting option.
func ToString() (desc string, f predicate.TransformFunc) {
	desc = "{}.String()"
	f = func(v interface{}) (r interface{}, ctx []predicate.ContextValue, err error) {
		s := fmt.Sprintf("%v", v)
		return s, []predicate.ContextValue{
			{Name: "string", Value: s, Pre: false},
		}, nil
	}
	return
}

// ToLower is a transformation predicate that converts any value to a string
// representation using `%v` formatting option.
func ToLower() (desc string, f predicate.TransformFunc) {
	desc = "ToLower({})"
	f = func(v interface{}) (r interface{}, ctx []predicate.ContextValue, err error) {
		s, ok := v.(string)
		if !ok {
			return nil, nil, fmt.Errorf(
				"value of type '%T' cannot be transformed to lowercase", v)
		}
		s = strings.ToLower(s)
		return s, []predicate.ContextValue{
			{Name: "lower", Value: s, Pre: false},
		}, nil
	}
	return
}

// ToUpper is a transformation predicate that converts any value to a string
// representation using `%v` formatting option.
func ToUpper() (desc string, f predicate.TransformFunc) {
	desc = "ToUpper({})"
	f = func(v interface{}) (r interface{}, ctx []predicate.ContextValue, err error) {
		s, ok := v.(string)
		if !ok {
			return nil, nil, fmt.Errorf(
				"value of type '%T' cannot be transformed to uppercase", v)
		}
		s = strings.ToUpper(s)
		return s, []predicate.ContextValue{
			{Name: "upper", Value: s, Pre: false},
		}, nil
	}
	return
}

// ToGoString returns a predicate that checks
// func ToGoString(p predicate.T) predicate.T {
// 	return testpredicate.MakeUnimplemented()
// }
