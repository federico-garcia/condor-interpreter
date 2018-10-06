package token

// Type represents the type of token the interpreter supports: indentifier, keyword, literal, operators, etc
type Type string

// Token stores its type and its value
type Token struct {
	Type    Type
	Literal string
}

const (
	// Special tokens
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"
	// Operators
	ASSIGN = "="
	PLUS   = "+"
	// Delimeters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
