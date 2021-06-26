package main

import (
	"fmt"
	"go/ast"
	"os"
	"path/filepath"

	"golang.org/x/tools/go/packages"

	"github.com/maargenton/go-testpredicate/pkg/codegen"
)

// FileDecl captures the relevant function declarations for one source file of
// the package.
type FileDecl struct {
	Name  string
	Funcs []FuncDecl
}

// FuncDecl captures the details of one relevant function declaration.
type FuncDecl struct {
	Pkg         string
	Name        string
	Comment     string
	Args        codegen.Fields
	Rets        codegen.Fields
	Transformer bool
}

func extractAPI(pkg *packages.Package) (files []FileDecl, err error) {
	if pkg.Module == nil {
		return nil, fmt.Errorf("target package is not part of a module")
	}
	basePath := pkg.Module.Dir

	for i, file := range pkg.Syntax {
		filename := pkg.CompiledGoFiles[i]
		filename, err := filepath.Rel(basePath, filename)
		if err != nil {
			return nil, fmt.Errorf("invalid filename: %w", err)
		}

		var funcs []FuncDecl
		for _, decl := range file.Decls {
			if decl, ok := decl.(*ast.FuncDecl); ok {
				if decl == nil || decl.Recv != nil {
					continue
				}

				var f = &FuncDecl{}
				f.Pkg = pkg.Name
				f.Name = decl.Name.String()
				f.Rets = codegen.FieldsFromAST(pkg.Fset, decl.Type.Results)
				if len(f.Rets) != 2 {
					continue
				}
				t := f.Rets[1].Type
				if t == "predicate.PredicateFunc" {
					f.Transformer = false
				} else if t == "predicate.TransformFunc" {
					f.Transformer = true
				} else {
					continue
				}

				f.Args = codegen.FieldsFromAST(pkg.Fset, decl.Type.Params)
				comment := ""
				if decl.Doc != nil {
					for _, doc := range decl.Doc.List {
						comment += doc.Text + "\n"
					}
				}
				f.Comment = comment
				funcs = append(funcs, *f)
			}
		}

		files = append(files, FileDecl{
			Name:  filename,
			Funcs: funcs,
		})
	}
	return
}

func run() error {
	if len(os.Args) < 3 {
		return fmt.Errorf("usage: forward_api <source_package> <template_name>")
	}
	sourcePath := os.Args[1]
	templatePath := os.Args[2]

	pkg, err := codegen.LoadPackage(sourcePath)
	if err != nil {
		return err
	}

	files, err := extractAPI(pkg)
	if err != nil {
		return err
	}

	return codegen.ApplyTemplate(templatePath, files)
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
