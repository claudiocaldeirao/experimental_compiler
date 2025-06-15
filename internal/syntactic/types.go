package syntactic

import "github.com/claudiocaldeirao/experimental_compiler/internal/token"

type Parser struct {
	tokens     []token.Token
	pos        int
	Statements []Statement
}

type Statement interface{}

type AssignStmt struct {
	VarName          string
	ExpressionTokens []token.Token
}

type PrintStmt struct {
	VarName string
}
