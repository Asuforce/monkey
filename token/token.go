package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal String
}

const {
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	INDENT = "INDENT"
	INT = "INT"

	ASSIGN = "="
	PLUS = "+"

	CONMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
}
