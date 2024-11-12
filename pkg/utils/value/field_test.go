package value_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/utils/value"
)

type TestCase struct {
	name     string
	value    interface{}
	keypath  string
	expected interface{}
}

func verifyFieldTestCase(t *testing.T, tc TestCase) {
	t.Helper()
	t.Run(fmt.Sprintf("Given %v", tc.name), func(t *testing.T) {
		t.Helper()
		t.Run(fmt.Sprintf("when calling Field(\"%v\")", tc.keypath), func(t *testing.T) {
			t.Helper()
			r := value.Field(tc.value, tc.keypath)
			t.Run("then expected value is returned", func(t *testing.T) {
				t.Helper()

				if !reflect.DeepEqual(r, tc.expected) {
					t.Errorf("\nexpected value: %#+v\nactual value:   %#+v",
						tc.expected, r)
				}
			})
		})
	})
}

func TestFieldOnNonFieldTypes(t *testing.T) {
	verifyFieldTestCase(t, TestCase{
		name:     "an integer value",
		value:    123,
		keypath:  "Name",
		expected: nil,
	})
}

func TestFieldOnStructTypes(t *testing.T) {
	var v = []TestStruct{
		{Name: "aaa"},
		{Name: "bbb"},
		{Name: "ccc", private: ""},
	}
	_ = v

	verifyFieldTestCase(t, TestCase{
		name:     "an array of struct and valid field name",
		value:    v,
		keypath:  "Name",
		expected: []interface{}{"aaa", "bbb", "ccc"},
	})

	verifyFieldTestCase(t, TestCase{
		name:     "an array of struct and invalid field name",
		value:    v,
		keypath:  "Names",
		expected: []interface{}{nil, nil, nil},
	})

	verifyFieldTestCase(t, TestCase{
		name:     "an array of struct and invalid keypath suffix",
		value:    v,
		keypath:  "Name.Value",
		expected: []interface{}{nil, nil, nil},
	})

	verifyFieldTestCase(t, TestCase{
		name:     "an array of struct and un-exported keypath",
		value:    v,
		keypath:  "private",
		expected: []interface{}{nil, nil, nil},
	})

	verifyFieldTestCase(t, TestCase{
		name:     "a func()(v) keypath",
		value:    v,
		keypath:  "Func",
		expected: []interface{}{"Func", "Func", "Func"},
	})

	verifyFieldTestCase(t, TestCase{
		name:     "a func()(v,err) method keypath",
		value:    v,
		keypath:  "FuncErr",
		expected: []interface{}{"FuncErr", "FuncErr", "FuncErr"},
	})
	verifyFieldTestCase(t, TestCase{
		name:     "a func(v)(v,v) method keypath",
		value:    v,
		keypath:  "FuncOther",
		expected: []interface{}{"FuncOther", "FuncOther", "FuncOther"},
	})

	verifyFieldTestCase(t, TestCase{
		name:     "a func(v)(v,err) method keypath",
		value:    v,
		keypath:  "FuncArgs",
		expected: []interface{}{nil, nil, nil},
	})
}

func TestFieldOnArrayOfStructTypes(t *testing.T) {
	verifyFieldTestCase(t, TestCase{
		name: "an array of pointers to struct",
		value: []*TestStruct{
			{Name: "aaa"},
			{Name: "bbb"},
			{Name: "ccc"},
		},
		keypath:  "Name",
		expected: []interface{}{"aaa", "bbb", "ccc"},
	})
}

func TestFieldOnMapTypes(t *testing.T) {
	verifyFieldTestCase(t, TestCase{
		name: "an map with matching keys",
		value: []obj{
			{"Name": "aaa"},
			{"Name": "bbb"},
			{"Name": "ccc"},
		},
		keypath:  "Name",
		expected: []interface{}{"aaa", "bbb", "ccc"},
	})

	verifyFieldTestCase(t, TestCase{
		name: "an map with non-matching keys",
		value: []obj{
			{"Name": "aaa"},
			{"Value": "bbb"},
			{"Name": "ccc"},
		},
		keypath:  "Name",
		expected: []interface{}{"aaa", nil, "ccc"},
	})

	verifyFieldTestCase(t, TestCase{
		name:     "an map with non-matching keys",
		value:    []int{1, 2, 3},
		keypath:  "Name",
		expected: []interface{}{nil, nil, nil},
	})
}

func TestFieldOnArrayOfMapTypes(t *testing.T) {
	verifyFieldTestCase(t, TestCase{
		name: "an map with non-matching keys",
		value: []interface{}{
			obj{"Name": "aaa"},
			obj{"Names": "bbb"},
			obj{"Name": "ccc"},
		},
		keypath:  "Name",
		expected: []interface{}{"aaa", nil, "ccc"},
	})

	verifyFieldTestCase(t, TestCase{
		name: "an map with non-matching keys",
		value: []interface{}{
			obj{"Name": "aaa"},
			obj{"Name": obj{"aaa": "bbb"}},
			obj{"Name": "ccc"},
		},
		keypath:  "Name.aaa",
		expected: []interface{}{nil, "bbb", nil},
	})
	verifyFieldTestCase(t, TestCase{
		name: "an map with non-matching keys",
		value: []interface{}{
			obj{"Name": "aaa"},
			obj{"Name": []obj{
				{"aaa": "bbb"},
				{"aaa": "ccc"},
			}},
			obj{"Name": "ccc"},
		},
		keypath:  "Name.aaa",
		expected: []interface{}{nil, []interface{}{"bbb", "ccc"}, nil},
	})

}

func TestFieldWithNiladicFunctionTarget(t *testing.T) {
	var sentinel = errors.New("sentinel")
	var err = fmt.Errorf("error: %w", sentinel)

	verifyFieldTestCase(t, TestCase{
		name:     "Error() value on an error",
		value:    err,
		keypath:  "Error",
		expected: "error: sentinel",
	})
}

// ---------------------------------------------------------------------------
// Support type

type obj map[string]interface{}

type TestStruct struct {
	Name    string
	Values  []string
	private string
}

func (t TestStruct) Func() string {
	return "Func"
}

func (t TestStruct) FuncErr() (string, error) {
	return "FuncErr", fmt.Errorf("error")
}

func (t TestStruct) FuncOther() (string, string) {
	return "FuncOther", "error"
}

func (t TestStruct) FuncArgs(i int) (string, error) {
	return "FuncArgs", nil
}
