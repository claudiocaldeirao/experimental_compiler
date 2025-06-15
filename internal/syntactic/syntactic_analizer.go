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
	// Check if it's a 'print' command
	if p.current().Type == token.PRINT {
		p.pos++ // consume 'print'

		// Parse expression after print
		var exprTokens []token.Token
		if p.match(token.IDENTIFIER) || p.match(token.NUMBER) {
			exprTokens = append(exprTokens, p.tokens[p.pos-1])

			if p.match(token.PLUS) || p.match(token.MINUS) || p.match(token.MULTIPLY) || p.match(token.DIVIDE) {
				exprTokens = append(exprTokens, p.tokens[p.pos-1])
				if p.match(token.IDENTIFIER) || p.match(token.NUMBER) {
					exprTokens = append(exprTokens, p.tokens[p.pos-1])
				} else {
					return false
				}
			}
		} else {
			return false
		}

		if !p.match(token.SEMICOLON) {
			return false
		}

		stmt := PrintStatement{
			ExpressionTokens: exprTokens,
		}
		p.Statements = append(p.Statements, stmt)
		return true
	}

	// Otherwise, assume it's an assignment
	idToken := p.current()
	if !p.match(token.IDENTIFIER) {
		return false
	}
	if !p.match(token.ASSIGN) {
		return false
	}

	// Capture the tokens that make up the expression
	var exprTokens []token.Token
	if p.match(token.IDENTIFIER) || p.match(token.NUMBER) {
		exprTokens = append(exprTokens, p.tokens[p.pos-1])

		if p.match(token.PLUS) || p.match(token.MINUS) || p.match(token.MULTIPLY) || p.match(token.DIVIDE) {
			exprTokens = append(exprTokens, p.tokens[p.pos-1])
			if p.match(token.IDENTIFIER) || p.match(token.NUMBER) {
				exprTokens = append(exprTokens, p.tokens[p.pos-1])
			} else {
				return false
			}
		}
	} else {
		return false
	}

	// Ensure command ends with a semicolon
	if !p.match(token.SEMICOLON) {
		return false
	}

	// Add statement to the AST
	stmt := AssignStatement{
		VarName:          idToken.Lexeme,
		ExpressionTokens: exprTokens,
	}
	p.Statements = append(p.Statements, stmt)
	return true
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
