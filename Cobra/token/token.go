package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// Identifiers + Literals
	IDENT = "IDENT"
	INT = "int"

	// Operators
	ASSIGN = "<="
	PLUS = "+"

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"

	// Parnathesis and braces

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "function"
	VAR = "var"

)