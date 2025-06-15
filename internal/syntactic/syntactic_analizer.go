package syntactic

import (
	"fmt"

	"github.com/claudiocaldeirao/experimental_compiler/internal/token"
)

func NewParser(tokens []token.Token) *Parser {
	return &Parser{tokens: tokens, pos: 0}
}

func (p *Parser) match(expected token.TokenType) bool {
	if p.pos < len(p.tokens) && p.tokens[p.pos].Type == expected {
		p.pos++
		return true
	}
	return false
}

func (p *Parser) current() token.Token {
	if p.pos < len(p.tokens) {
		return p.tokens[p.pos]
	}
	return token.Token{Type: token.EOF}
}

// Command := ID '=' Expression ';'
func (p *Parser) parseCommand() bool {
	if !p.match(token.IDENTIFIER) {
		return false
	}
	if !p.match(token.ASSIGN) {
		return false
	}
	if !p.parseExpression() {
		return false
	}
	return p.match(token.SEMICOLON)
}

// Expression := ID | NUM | ID op NUM
func (p *Parser) parseExpression() bool {
	if p.match(token.IDENTIFIER) || p.match(token.NUMBER) {
		if p.match(token.PLUS) || p.match(token.MINUS) || p.match(token.MULTIPLY) || p.match(token.DIVIDE) {
			return p.match(token.IDENTIFIER) || p.match(token.NUMBER)
		}
		return true
	}
	return false
}

// Program := 'BEGIN' {Command} 'END'
func (p *Parser) ParseProgram() bool {
	if !p.match("BEGIN") {
		fmt.Println("Expected BEGIN")
		return false
	}

	for p.current().Type != token.END && p.current().Type != token.EOF {
		if !p.parseCommand() {
			fmt.Println("Error parsing command near:", p.current().Lexeme)
			return false
		}
	}

	if !p.match(token.END) {
		fmt.Println("Expected END")
		return false
	}

	return true
}
