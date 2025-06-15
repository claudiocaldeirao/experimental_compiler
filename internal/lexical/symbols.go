package lexical

import (
	"github.com/claudiocaldeirao/experimental_compiler/internal/token"
)

var Symbols = map[string]token.TokenType{
	"+":  token.PLUS,
	"-":  token.MINUS,
	"*":  token.MULTIPLY,
	"/":  token.DIVIDE,
	"=":  token.ASSIGN,
	"==": token.EQUAL,
	">":  token.GT,
	"<":  token.LT,
	"(":  token.LPAREN,
	")":  token.RPAREN,
	";":  token.SEMICOLON,
}
