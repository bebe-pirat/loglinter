package analyzer_test

import (
	"testing"

	"github.com/bebe-pirat/loglinter/pkg/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAll(t *testing.T) {
	analysistest.RunWithSuggestedFixes(t, analysistest.TestData(), analyzer.New(), "./src/p")
}
