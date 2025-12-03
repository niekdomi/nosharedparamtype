// Package analyzer implements the linting logic of the no shared parameter type linter.
package analyzer

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// Analyzer reports shared parameter type declarations.
var Analyzer = &analysis.Analyzer{
	Name: "nosharedparamtype",
	Doc:  "enforces explicit type declarations for all function parameters",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			funcDecl, ok := node.(*ast.FuncDecl)
			if !ok || funcDecl.Type.Params == nil {
				return true
			}

			checkParams(pass, funcDecl)

			return true
		})
	}

	return nil, nil
}

func checkParams(pass *analysis.Pass, funcDecl *ast.FuncDecl) {
	for _, field := range funcDecl.Type.Params.List {
		if len(field.Names) <= 1 {
			continue
		}

		pass.Report(analysis.Diagnostic{
			Pos:     field.Pos(),
			Message: fmt.Sprintf("function %s has parameters using shared type", funcDecl.Name.Name),
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "add explicit type for each parameter",
					TextEdits: []analysis.TextEdit{
						{
							Pos:     field.Pos(),
							End:     field.End(),
							NewText: []byte(expandParams(field)),
						},
					},
				},
			},
			End:      0,
			Category: "",
			URL:      "",
			Related:  []analysis.RelatedInformation{},
		})
	}
}

func expandParams(field *ast.Field) string {
	typeStr := types.ExprString(field.Type)
	params := make([]string, 0, len(field.Names))

	for _, name := range field.Names {
		params = append(params, fmt.Sprintf("%s %s", name.Name, typeStr))
	}

	return strings.Join(params, ", ")
}
