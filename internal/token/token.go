package token

type TokenType string

type Token struct {
	Type   TokenType
	Lexeme string
}

const (
	// Operators and delimiters
	PLUS      TokenType = "PLUS"
	MINUS     TokenType = "MINUS"
	MULTIPLY  TokenType = "MULTIPLY"
	DIVIDE    TokenType = "DIVIDE"
	ASSIGN    TokenType = "ASSIGN"
	EQUAL     TokenType = "EQUAL"
	GT        TokenType = "GT"
	LT        TokenType = "LT"
	LPAREN    TokenType = "LPAREN"
	RPAREN    TokenType = "RPAREN"
	SEMICOLON TokenType = "SEMICOLON"

	// Keywords
	IF    TokenType = "IF"
	THEN  TokenType = "THEN"
	ELSE  TokenType = "ELSE"
	WHILE TokenType = "WHILE"
	DO    TokenType = "DO"
	BEGIN TokenType = "BEGIN"
	END   TokenType = "END"
	PRINT TokenType = "PRINT"

	// General types
	IDENTIFIER TokenType = "IDENTIFIER"
	NUMBER     TokenType = "NUMBER"
	EOF        TokenType = "EOF"
	ILLEGAL    TokenType = "ILLEGAL"
)
