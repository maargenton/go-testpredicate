package utils

import "fmt"

// WrapError appends the nested error after the formated message, on a new line
func WrapError(nestedErr error, format string, a ...interface{}) error {
	if nestedErr != nil {
		msg := fmt.Sprintf(format, a...)
		return fmt.Errorf("%v\n%v", msg, nestedErr)
	}

	return fmt.Errorf(format, a...)
}

// FormatValue retruns a string representing the value, truncated
// to a reasonable max size of 50
func FormatValue(value interface{}) string {
	s := fmt.Sprintf("%#v", value)
	l := len(s)

	if l > 50 {
		s = s[0:24] + "..." + s[l-23:l]
	}
	return s
}

// FormatDetails formats a list of assertion details into a string. When details
// starts with a string, it is interpreted as a format string using the rest of
// ther details as argumetns. Otherwise, details are printed as a space
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
