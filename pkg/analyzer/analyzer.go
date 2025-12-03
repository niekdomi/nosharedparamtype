// Package analyzer implements the linting logic of the no shared parameter type linter.
package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "nosharedparamtype",
	Doc:  "enforces explicit type declarations for all function parameters",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	inspect := func(node ast.Node) bool {
		funcDecl, ok := node.(*ast.FuncDecl)
		if !ok {
			return true
		}

		for _, field := range funcDecl.Type.Params.List {
			if len(field.Names) > 1 {
				pass.Reportf(
					node.Pos(),
					"function %s has type sharing parameters",
					funcDecl.Name.Name,
				)
			}
		}

		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}

	return nil, nil
}
