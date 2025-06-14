package lexical

import (
	"strings"
	"unicode"
)

type Token struct {
	Type  string
	Value string
}

var keywords = map[string]string{
	"IF": "IF", "THEN": "THEN", "ELSE": "ELSE",
	"WHILE": "WHILE", "DO": "DO",
	"BEGIN": "BEGIN", "END": "END",
	"PRINT": "PRINT",
}

var symbols = map[string]string{
	"+": "PLUS", "-": "MINUS", "*": "MULTIPLY", "/": "DIVIDE",
	"=": "ASSIGN", "==": "EQUAL", ">": "GT", "<": "LT",
	"(": "LPAREN", ")": "RPAREN", ";": "SEMICOLON",
}

func Tokenize(input string) []Token {
	var tokens []Token
	var i int
	for i < len(input) {
		c := rune(input[i])

		if unicode.IsSpace(c) {
			i++
			continue
		}

		if unicode.IsLetter(c) {
			start := i
			for i < len(input) && (unicode.IsLetter(rune(input[i])) || unicode.IsDigit(rune(input[i]))) {
				i++
			}
			word := input[start:i]
			upper := strings.ToUpper(word)
			if t, found := keywords[upper]; found {
				tokens = append(tokens, Token{Type: t, Value: word})
			} else {
				tokens = append(tokens, Token{Type: "ID", Value: word})
			}
			continue
		}

		if unicode.IsDigit(c) {
			start := i
			for i < len(input) && (unicode.IsDigit(rune(input[i])) || rune(input[i]) == '.') {
				i++
			}
			tokens = append(tokens, Token{Type: "NUM", Value: input[start:i]})
			continue
		}

		if i+1 < len(input) {
			two := input[i : i+2]
			if t, found := symbols[two]; found {
				tokens = append(tokens, Token{Type: t, Value: two})
				i += 2
				continue
			}
		}

		if t, found := symbols[string(c)]; found {
			tokens = append(tokens, Token{Type: t, Value: string(c)})
			i++
			continue
		}

		i++
	}

	return tokens
}
