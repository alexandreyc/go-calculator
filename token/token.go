package token

// Token identifiers
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Literals
	INT = "INT"

	// Operators
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"

	LPAREN = "("
	RPAREN = ")"
)

// TokenType ...
type TokenType string

// Token ...
type Token struct {
	Type    TokenType
	Literal string
}
