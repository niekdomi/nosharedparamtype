package analyzer

import (
	"fmt"
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
			for _, name := range field.Names {
				fmt.Print("Param-Name: ", name.Name)
			}

			fmt.Print("Param-Type: ", field.Type)
		}

		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}

	return nil, nil
}
