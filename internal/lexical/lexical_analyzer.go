package lexical

import (
	"unicode"

	"github.com/claudiocaldeirao/experimental_compiler/internal/token"
)

func Tokenize(source string) []token.Token {
	var tokens []token.Token
	var pos int

	for pos < len(source) {
		ch := source[pos]

		if isWhitespace(ch) {
			pos++
			continue
		}

		if isLetter(ch) {
			start := pos
			for pos < len(source) && isLetterOrDigit(source[pos]) {
				pos++
			}
			lexeme := source[start:pos]
			tokType, ok := Keywords[lexeme]
			if !ok {
				tokType = token.IDENTIFIER
			}
			tokens = append(tokens, token.Token{Type: tokType, Lexeme: lexeme})
			continue
		}

		if isDigit(ch) {
			start := pos
			for pos < len(source) && isDigit(source[pos]) {
				pos++
			}
			lexeme := source[start:pos]
			tokens = append(tokens, token.Token{Type: token.NUMBER, Lexeme: lexeme})
			continue
		}

		if pos+1 < len(source) {
			twoChar := source[pos : pos+2]
			if tokType, ok := Symbols[twoChar]; ok {
				tokens = append(tokens, token.Token{Type: tokType, Lexeme: twoChar})
				pos += 2
				continue
			}
		}

		if tokType, ok := Symbols[string(ch)]; ok {
			tokens = append(tokens, token.Token{Type: tokType, Lexeme: string(ch)})
			pos++
			continue
		}

		tokens = append(tokens, token.Token{Type: token.ILLEGAL, Lexeme: string(ch)})
		pos++
	}

	tokens = append(tokens, token.Token{Type: token.EOF, Lexeme: ""})
	return tokens
}

func isLetter(r byte) bool {
	return unicode.IsLetter(rune(r))
}

func isDigit(r byte) bool {
	return '0' <= r && r <= '9'
}

func isLetterOrDigit(r byte) bool {
	return isLetter(r) || isDigit(r)
}

func isWhitespace(r byte) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}
