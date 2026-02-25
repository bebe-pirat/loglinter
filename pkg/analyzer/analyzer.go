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

var Analyzer = &analysis.Analyzer{
	Name:     "loglinter",
	Doc:      "Checks that log messages starts with lowercase, uses only english symbols, doesn't use any special symbols or emojii and have no sensitive data",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func New() *analysis.Analyzer {
	return Analyzer
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			callExpr, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}

			selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			obj, ok := pass.TypesInfo.Uses[selExpr.Sel]
			if !ok {
				return true
			}
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

			firstParam := callExpr.Args[0]
			var message string
			switch arg := firstParam.(type) {
			case *ast.BasicLit:
				message, _ = strconv.Unquote(arg.Value)
			case *ast.BinaryExpr:
				if lit, ok := arg.X.(*ast.BasicLit); ok {
					message, _ = strconv.Unquote(lit.Value)
				}
			}

			if message != "" {
				return checkLogMessage(pass, firstParam, message)
			}

			return true
		})
	}

	return nil, nil
}

func checkLogMessage(pass *analysis.Pass, node ast.Expr, message string) bool {
	runes := []rune(message)

	if !isFirstSymbolLowerCase(runes) {
		newParam := strconv.Quote(firstSymbolToLowercase(runes))
		pass.Report(analysis.Diagnostic{
			Pos:     node.Pos(),
			End:     node.End(),
			Message: "first letter in log message should be lowercase",
			SuggestedFixes: []analysis.SuggestedFix{
				{
					TextEdits: []analysis.TextEdit{
						{
							Pos:     node.Pos(),
							End:     node.End(),
							NewText: []byte(newParam),
						},
					},
				},
			},
		})
		return true
	}

	if !isStringEnglish(runes) {
		pass.Reportf(node.Pos(), "log message should contain only English characters")
		return true
	}

	if containsSpecialSymbolsOrEmoji(message) {
		pass.Reportf(node.Pos(), "log message contains special symbols or emoji")
		return true
	}

	if containsSensitiveKeyword(message) {
		pass.Reportf(node.Pos(), "log message may contain sensitive data")
		return true
	}

	return true
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

func containsSpecialSymbolsOrEmoji(str string) bool {
	pattern := `^[a-z0-9\s_=:"]+$`
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}

	mathed := re.MatchString(str)

	return !mathed
}

func containsSensitiveKeyword(s string) bool {
	sensiviteKeywords := []string{
		"password",
		"api_key",
		"apikey",
		"secret",
		":",
		"=",
	}

	s = strings.ToLower(s)
	for _, word := range sensiviteKeywords {
		if strings.Contains(s, word) {
			return true
		}
	}

	return false
}

func firstSymbolToLowercase(runes []rune) string {
	if len(runes) > 0 {
		runes[0] = unicode.ToLower(runes[0])
	}

	return string(runes)
}
