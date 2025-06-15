package generator

import (
	"fmt"
	"strings"

	"github.com/claudiocaldeirao/experimental_compiler/internal/syntactic"
)

func GenerateJS(statements []syntactic.Statement) string {
	var builder strings.Builder

	for _, stmt := range statements {
		switch s := stmt.(type) {
		case syntactic.AssignStatement:
			builder.WriteString(fmt.Sprintf("let %s = ", s.VarName))
			for i, t := range s.ExpressionTokens {
				if i > 0 {
					builder.WriteString(" ")
				}
				builder.WriteString(t.Lexeme)
			}
			builder.WriteString(";\n")
		case syntactic.PrintStatement:
			builder.WriteString("console.log(")
			for i, t := range s.ExpressionTokens {
				if i > 0 {
					builder.WriteString(" ")
				}
				builder.WriteString(t.Lexeme)
			}
			builder.WriteString(");\n")
		}
	}

	return builder.String()
}
