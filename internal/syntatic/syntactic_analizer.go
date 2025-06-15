package syntatic

import (
	"fmt"

	"github.com/claudiocaldeirao/experimental_compiler/internal/types"
)

func NewParser(tokens []types.Token) *Parser {
	return &Parser{tokens: tokens, pos: 0}
}

func (p *Parser) match(expected types.TokenType) bool {
	if p.pos < len(p.tokens) && p.tokens[p.pos].Type == expected {
		p.pos++
		return true
	}
	return false
}

func (p *Parser) current() types.Token {
	if p.pos < len(p.tokens) {
		return p.tokens[p.pos]
	}
	return types.Token{Type: "EOF"}
}

// Command := ID '=' Expression ';'
func (p *Parser) parseCommand() bool {
	if !p.match("ID") {
		return false
	}
	if !p.match("ASSIGN") {
		return false
	}
	if !p.parseExpression() {
		return false
	}
	return p.match("SEMICOLON")
}

// Expression := ID | NUM | ID op NUM
func (p *Parser) parseExpression() bool {
	if p.match("ID") || p.match("NUM") {
		if p.match("PLUS") || p.match("MINUS") || p.match("MULTIPLY") || p.match("DIVIDE") {
			return p.match("ID") || p.match("NUM")
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

	for p.current().Type != "END" && p.current().Type != "EOF" {
		if !p.parseCommand() {
			fmt.Println("Error parsing command near:", p.current().Lexeme)
			return false
		}
	}

	if !p.match("END") {
		fmt.Println("Expected END")
		return false
	}

	return true
}
