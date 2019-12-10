package p

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/maargenton/go-testpredicate/pkg/predicate"
	"github.com/maargenton/go-testpredicate/pkg/prettyprint"
)

// Matches returns a predicate that checks if a string matches a regular
// expression
func Matches(re string) predicate.T {

	return predicate.Make(
		fmt.Sprintf("value matches /%v/", re),
		func(value interface{}) (predicate.Result, error) {

			s, ok := value.(string)
			if !ok {
				return predicate.Invalid, fmt.Errorf(
					"value of type '%T' cannot be matched against a regexp", value)
			}
			m, err := regexp.MatchString(re, s)
			if err != nil {
				return predicate.Invalid,
					fmt.Errorf("failed to compile regexp, %v", err)
			}

			if m {
				return predicate.Passed, nil
			}
			return predicate.Failed, nil
		})
}

// ToUpper returns a predicate that checks if the uppercase version of a string
// passes the nested predicate
func ToUpper(p predicate.T) predicate.T {
	return predicate.Make(
		strings.Replace(p.String(), "value", "ToUpper(value)", -1),
		func(value interface{}) (predicate.Result, error) {

			s, ok := value.(string)
			if !ok {
				return predicate.Invalid, fmt.Errorf(
					"value of type '%T' cannot be transformed to uppercase", value)
			}

			s = strings.ToUpper(s)
			r, err := p.Evaluate(s)
			err = predicate.WrapError(err, "uppercase: %v", prettyprint.FormatValue(s))
			return r, err
		})
}

// ToLower returns a predicate that checks if the lowercase version of a string
// passes the nested predicate
func ToLower(p predicate.T) predicate.T {
	return predicate.Make(
		strings.Replace(p.String(), "value", "ToLower(value)", -1),
		func(value interface{}) (predicate.Result, error) {

			s, ok := value.(string)
			if !ok {
				return predicate.Invalid, fmt.Errorf(
					"value of type '%T' cannot be transformed to uppercase", value)
			}

			s = strings.ToLower(s)
			r, err := p.Evaluate(s)
			err = predicate.WrapError(err, "lowercase: %v", prettyprint.FormatValue(s))
			return r, err
		})
}

// ToString returns a predicate that checks if a value converted to a string
// passes the nested predicate. Value is conversted using fmt.Sprintf("%v", ...)
func ToString(p predicate.T) predicate.T {
	return predicate.Make(
		strings.Replace(p.String(), "value", "ToString(value)", -1),
		func(value interface{}) (predicate.Result, error) {

			s := fmt.Sprintf("%v", value)
			r, err := p.Evaluate(s)
			err = predicate.WrapError(err, "string: %v", prettyprint.FormatValue(s))
			return r, err
		})
}

// ToGoString returns a predicate that checks
// func ToGoString(p predicate.T) predicate.T {
// 	return testpredicate.MakeUnimplemented()
// }
