package utils

import (
	"fmt"

	"github.com/maargenton/go-testpredicate/pkg/prettyprint"
)

// WrapError appends the nested error after the formated message, on a new line
func WrapError(nestedErr error, format string, a ...interface{}) error {
	if nestedErr != nil {
		msg := fmt.Sprintf(format, a...)
		return fmt.Errorf("%v\n%v", msg, nestedErr)
	}

	return fmt.Errorf(format, a...)
}

// FormatValue retruns a string representing the value, truncated
// to a maximum length of 80.
func FormatValue(v interface{}) string {
	return prettyprint.FormatValue(v)
}

func formatValue(value interface{}) string {
	if _, ok := value.(string); ok {
		s := fmt.Sprintf("%#v", value)
		l := len(s)
		if l <= 80 {
			return s
		}
		return s[0:76] + "..." + s[l-1:l]
	}

	s := fmt.Sprintf("%#v", value)
	l := len(s)
	if l <= 80 {
		return s
	}

	s = fmt.Sprintf("%+v", value)
	l = len(s)
	if l <= 80 {
		return s
	}

	return s[0:77] + "..."
}

// FormatDetails formats a list of assertion details into a string. When details
// starts with a string, it is interpreted as a format string using the rest of
// their details as argumetns. Otherwise, details are printed as a space
// separated list of values.
func FormatDetails(details ...interface{}) string {
	if len(details) == 0 {
		return ""
	}

	if s, ok := details[0].(string); ok {
		return fmt.Sprintf(s, details[1:]...)
	}

	return fmt.Sprint(details...)
}
