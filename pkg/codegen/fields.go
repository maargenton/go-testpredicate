package codegen

import (
	"go/ast"
	"go/printer"
	"go/token"
	"strings"
)

// Field captures one fragment of the textual representation of a function
// argument list, return list or receiver list
type Field struct {
	Names []string
	Type  string
}

// Fields captures all fragments of the textual representation of a function
// argument list, return list or receiver list
type Fields []Field

// FieldsFromAST capture the textual components of an ast.FieldList and returns
// the corresponding codegen.Fields.
func FieldsFromAST(fset *token.FileSet, l *ast.FieldList) (fields Fields) {
	if l == nil {
		return nil
	}
	for _, ff := range l.List {
		var f Field
		for _, name := range ff.Names {
			f.Names = append(f.Names, name.String())
		}
		var b strings.Builder
		printer.Fprint(&b, fset, ff.Type)
		f.Type = b.String()

		fields = append(fields, f)
	}
	return
}

// String returns a string representation of the fields
func (ff Fields) String() string {
	var r strings.Builder
	for i, f := range ff {
		if i > 0 {
			r.WriteString(", ")
		}
		if len(f.Names) > 0 {
			r.WriteString(strings.Join(f.Names, ", "))
			r.WriteString(" ")
		}
		r.WriteString(f.Type)
	}
	return r.String()
}

// Fwd returns a string representation of a forwarding list for the
// fields, assuming all fields have a name
func (ff Fields) Fwd() string {
	var names []string
	for _, f := range ff {
		names = append(names, f.Names...)
	}
	return strings.Join(names, ", ")
}
