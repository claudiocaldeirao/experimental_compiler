package lexical

import "github.com/claudiocaldeirao/experimental_compiler/internal/token"

var Keywords = map[string]token.TokenType{
	"IF":    token.IF,
	"THEN":  token.THEN,
	"ELSE":  token.ELSE,
	"WHILE": token.WHILE,
	"DO":    token.DO,
	"BEGIN": token.BEGIN,
	"END":   token.END,
	"PRINT": token.PRINT,
}
