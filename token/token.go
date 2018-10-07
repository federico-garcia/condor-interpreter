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

var keywords = map[string]Type{
	"fn":  "FUNCTION",
	"let": "LET",
}

// LookupIdentifier returns the token type (keyword or variable name) of a given identifier
func LookupIdentifier(s string) Type {
	tokenType, ok := keywords[s]
	if ok {
		return tokenType
	}
	return IDENT
}
