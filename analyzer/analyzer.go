package analyzer

import (
	"go/ast"
	"go/types"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

func NewLogChecker() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "loglinter",
		Doc:      "short documentation about linter",
		Run:      run,
		URL:      "https://documentation.com",
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			callExpr, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}

			selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
			obj := pass.TypesInfo.Uses[selExpr.Sel]

			fn, ok := obj.(*types.Func)
			if !ok {
				return true
			}
			pkg := fn.Pkg()
			if pkg == nil {
				return true
			}
			if pkg.Path() != "log/slog" && pkg.Path() != "go.uber.org/zap" {
				return true
			}

			method := fn.Name()
			if method != "Info" && method != "Error" && method != "Warn" && method != "Debug" {
				return true
			}

			if len(callExpr.Args) <= 0 {
				return true
			}

			for _, arg := range callExpr.Args {
				switch arg.(type) {
				case *ast.BasicLit:
					n := arg.(*ast.BasicLit)
					str, _ := strconv.Unquote(n.Value)
					runes := []rune(str)

					if !isFirstSymbolLowerCase(runes) {
						pass.Reportf(callExpr.Pos(), "first letter in log message is not lower case")
						return true
					}

					if !isStringEnglish(runes) {
						pass.Reportf(callExpr.Pos(), "log message contains non-english characters")
						return true
					}

					if !doesNotContainEmojiiAndSpecialSymbols(str) {
						pass.Reportf(callExpr.Pos(), "log message contains special symbols or emojii")
						return true
					}

					if doesContainSensitiveKeyWord(str) {
						pass.Reportf(callExpr.Pos(), "log message may contain sensetive or dynamic data")
					}
				case *ast.BinaryExpr:
					pass.Reportf(callExpr.Pos(), "log message may contain sensetive or dynamic data")
				default:
				}
			}

			return true
		})
	}

	return nil, nil
}

func isFirstSymbolLowerCase(str []rune) bool {
	if len(str) <= 0 {
		return false
	}

	if unicode.IsLetter(str[0]) {
		if unicode.IsUpper(str[0]) {
			return false
		}
	}

	return true
}

func isStringEnglish(str []rune) bool {
	if len(str) <= 0 {
		return true
	}

	pattern := `[a-zA-Z]`
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}
	for _, letter := range str {
		if unicode.IsLetter(letter) && !re.MatchString(string(letter)) {
			return false
		}
	}

	return true
}

func doesNotContainEmojiiAndSpecialSymbols(str string) bool {
	pattern := `^[a-z0-9\s"]+$`
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}

	mathed := re.MatchString(str)

	return mathed
}

func doesContainSensitiveKeyWord(s string) bool {
	sensiviteKeywords := []string{
		"password",
		"api_key",
		"apikey",
		"secret",
	}

	s = strings.ToLower(s)
	for _, word := range sensiviteKeywords {
		if strings.Contains(s, word) {
			return true
		}
	}

	return false
}
