package plugin

import (
	"github.com/bebe-pirat/loglinter/pkg/analyzer"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
	"github.com/golangci/golangci-lint/v2/pkg/lint/linter"
	"golang.org/x/tools/go/analysis"
)

func New() *linter.Config {
	return &linter.Config{
		Linter: goanalysis.NewLinter(
			"loglinter",
			"Custom linter for log checking",
			[]*analysis.Analyzer{analyzer.Analyzer},
			nil,
		).WithLoadMode(goanalysis.LoadModeSyntax),
	}
}
