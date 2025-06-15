package syntactic

import (
	"fmt"

	"github.com/claudiocaldeirao/experimental_compiler/internal/token"
)

func NewParser(tokens []token.Token) *Parser {
	return &Parser{tokens: tokens, pos: 0}
}

// Program := 'BEGIN' {Command} 'END'
func (p *Parser) ParseProgram() []Statement {
	if !p.match(token.BEGIN) {
		fmt.Println("Expected BEGIN")
		return nil
	}

	for p.current().Type != token.END && p.current().Type != token.EOF {
		if !p.parseCommand() {
			fmt.Println("Error parsing command near:", p.current().Lexeme)
			return nil
		}
	}

	if !p.match(token.END) {
		fmt.Println("Expected END")
		return nil
	}

	return p.Statements
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

// Command := printStatement | assignmentStatement
func (p *Parser) parseCommand() bool {
	switch p.current().Type {
	case token.PRINT:
		return p.parsePrint()
	case token.IDENTIFIER:
		return p.parseAssignment()
	default:
		fmt.Printf("Unexpected token: %v\n", p.current().Lexeme)
		return false
	}
}

// AssignStatement := ID '=' Expression ';'
func (p *Parser) parseAssignment() bool {
	idToken := p.current()
	if !p.match(token.IDENTIFIER) {
		return false
	}
	if !p.match(token.ASSIGN) {
		return false
	}

	exprTokens, ok := p.parseExpressionTokens()
	if !ok || !p.match(token.SEMICOLON) {
		return false
	}

	stmt := AssignStatement{
		VarName:          idToken.Lexeme,
		ExpressionTokens: exprTokens,
	}
	p.Statements = append(p.Statements, stmt)
	return true
}

// PrintStatement := 'PRINT' Expression ';'
func (p *Parser) parsePrint() bool {
	p.pos++

	exprTokens, ok := p.parseExpressionTokens()
	if !ok || !p.match(token.SEMICOLON) {
		return false
	}

	stmt := PrintStatement{
		ExpressionTokens: exprTokens,
	}
	p.Statements = append(p.Statements, stmt)
	return true
}

// ExpressionTokens := ID | NUMBER (OPERATOR ID | NUMBER)*
func (p *Parser) parseExpressionTokens() ([]token.Token, bool) {
	var tokens []token.Token

	if p.match(token.IDENTIFIER) || p.match(token.NUMBER) {
		tokens = append(tokens, p.tokens[p.pos-1])

		if p.match(token.PLUS) || p.match(token.MINUS) || p.match(token.MULTIPLY) || p.match(token.DIVIDE) {
			tokens = append(tokens, p.tokens[p.pos-1])
			if p.match(token.IDENTIFIER) || p.match(token.NUMBER) {
				tokens = append(tokens, p.tokens[p.pos-1])
			} else {
				return nil, false
			}
		}
		return tokens, true
	}

	return nil, false
}
