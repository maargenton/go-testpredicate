package utils_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/maargenton/go-testpredicate/utils"
)

// ---------------------------------------------------------------------------
// utils.WrapError()
// ---------------------------------------------------------------------------

func TestWrapError(t *testing.T) {
	err1 := fmt.Errorf("error1")
	err2 := utils.WrapError(err1, "error2: %v", 123)
	if !strings.HasPrefix(err2.Error(), "error2: 123") {
		t.Errorf("unexpected error: '%v'", err2)

	}
	if !strings.HasSuffix(err2.Error(), "error1") {
		t.Errorf("unexpected error: '%v'", err2)
	}
}

func TestWrapErrorWithNoBaseError(t *testing.T) {
	err2 := utils.WrapError(nil, "error2: %v", 123)
	if !strings.HasPrefix(err2.Error(), "error2: 123") ||
		strings.Contains(err2.Error(), "\n") {
		t.Errorf("unexpected error: '%v'", err2)
	}
}

// ---------------------------------------------------------------------------
// utils.FormatDetails()
// ---------------------------------------------------------------------------

func TestFormatDetailsEmpty(t *testing.T) {
	s := utils.FormatDetails()
	if s != "" {
		t.Errorf("unexpected details output: %v", s)
	}
}

func TestFormatDetailsList(t *testing.T) {
	s := utils.FormatDetails(1, 2, 3)
	if s != "1 2 3" {
		t.Errorf("unexpected details output: '%v'", s)
	}
}

func TestFormatDetailsFormat(t *testing.T) {
	fmt := "a:%v b:%v c:%v"
	s := utils.FormatDetails(fmt, 1, 2, 3)
	if s != "a:1 b:2 c:3" {
		t.Errorf("unexpected details output: '%v'", s)
	}
}
