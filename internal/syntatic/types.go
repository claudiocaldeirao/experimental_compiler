package syntatic

import "github.com/claudiocaldeirao/experimental_compiler/internal/types"

type Parser struct {
	tokens []types.Token
	pos    int
}
