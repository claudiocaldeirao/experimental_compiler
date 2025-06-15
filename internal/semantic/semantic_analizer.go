package semantic

import (
	"fmt"

	"github.com/claudiocaldeirao/experimental_compiler/internal/syntactic"
	"github.com/claudiocaldeirao/experimental_compiler/internal/token"
)

func NewSemanticAnalyzer(statements []syntactic.Statement) *SemanticAnalyzer {
	return &SemanticAnalyzer{
		Variables:  make(map[string]*Variable),
		Statements: statements,
	}
}

func (s *SemanticAnalyzer) Analyze() {
	for _, stmt := range s.Statements {
		switch node := stmt.(type) {
		case syntactic.AssignStatement:
			s.handleAssignment(node)
		case syntactic.PrintStatement:
			s.handleUsage(node.VarName)
		}
	}
	s.checkUnused()
}

func (s *SemanticAnalyzer) handleAssignment(stmt syntactic.AssignStatement) {
	/// Mark the variable as declared
	if _, exists := s.Variables[stmt.VarName]; !exists {
		s.Variables[stmt.VarName] = &Variable{Name: stmt.VarName}
	}
	s.Variables[stmt.VarName].Declared = true

	// Validate variables used in the expression
	for _, et := range stmt.ExpressionTokens {
		if et.Type == token.IDENTIFIER {
			s.handleUsage(et.Lexeme)
		}
	}
}

func (s *SemanticAnalyzer) handleUsage(name string) {
	variable, exists := s.Variables[name]
	if !exists {
		s.Errors = append(s.Errors, fmt.Sprintf("Variable '%s' used before declaration", name))
		s.Variables[name] = &Variable{Name: name, Used: true}
	} else {
		variable.Used = true
	}
}

func (s *SemanticAnalyzer) checkUnused() {
	for _, v := range s.Variables {
		if v.Declared && !v.Used {
			s.Errors = append(s.Errors, fmt.Sprintf("Variable '%s' declared but never used", v.Name))
		}
	}
}

func (s *SemanticAnalyzer) HasErrors() bool {
	return len(s.Errors) > 0
}

func (s *SemanticAnalyzer) ReportErrors() {
	for _, msg := range s.Errors {
		fmt.Println("Semantic Error:", msg)
	}
}
