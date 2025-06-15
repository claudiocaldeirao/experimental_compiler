package syntactic

import "github.com/claudiocaldeirao/experimental_compiler/internal/token"

type Parser struct {
	tokens []token.Token
	pos    int
}
