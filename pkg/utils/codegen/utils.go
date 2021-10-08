// Package codegen is a small package intended to help with integrating custom
// code generation tools in your project. It provides a 5 steps standardized
// workflow as follows:
//   - Load a package, parse its source code, produce an AST.
//   - Scan the loaded package and AST to produce an information tree (custom step).
//   - Load a (custom) template file and apply it to the extracted information tree.
//   - Format the generate output through go-imports.
//   - Save the generated file to disk.
package codegen

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"strings"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/imports"
)

// LoadPackage loads and parses the source code of a package names relative to
// the current working directory
func LoadPackage(name string) (pkg *packages.Package, err error) {
	cfg := &packages.Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedCompiledGoFiles |
			packages.NeedImports |
			packages.NeedTypes |
			packages.NeedTypesSizes |
			packages.NeedSyntax |
			packages.NeedTypesInfo |
			packages.NeedModule,
		Tests:      false,
		BuildFlags: []string{},
	}

	pkgs, err := packages.Load(cfg, name)
	if err != nil {
		return nil,
			fmt.Errorf("failed to load package containing target file: %w", err)
	}
	if len(pkgs) < 1 {
		return nil, fmt.Errorf("no package loaded")
	}
	pkg = pkgs[0]
	return
}

// ApplyTemplate loads a template file, executes it on the `data` object,
// formats the output through go-imports, and save the result alongside the
// original template file with a .go extension.
func ApplyTemplate(templateFn string, data interface{}) error {
	tmpl, err := template.ParseFiles(templateFn)
	if err != nil {
		return fmt.Errorf("failed to load template: %w", err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	output, err := imports.Process("", buf.Bytes(), &imports.Options{
		AllErrors: true, Comments: true, TabIndent: true, TabWidth: 8,
	})
	if err != nil {
		return fmt.Errorf("failed to format generate code: %w", err)
	}

	outputFn := strings.ReplaceAll(templateFn, ".tmpl", ".go")
	err = ioutil.WriteFile(outputFn, output, 0644)
	return err
}
