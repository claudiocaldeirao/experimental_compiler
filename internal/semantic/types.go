package semantic

import "github.com/claudiocaldeirao/experimental_compiler/internal/syntactic"

type Variable struct {
	Name     string
	Declared bool
	Used     bool
}

type SemanticAnalyzer struct {
	Variables  map[string]*Variable
	Statements []syntactic.Statement
	Errors     []string
}
