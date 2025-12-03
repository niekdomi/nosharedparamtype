// Package main provides the command-line interface for the nosharedparamtype linter.
package main

import (
	"github.com/niekdomi/nosharedparamtype/pkg/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
