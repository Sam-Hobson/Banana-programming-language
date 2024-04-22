package token

type TokenType uint8

var tokenNames = [...]string{
	"ILLEGAL",
	"EOF",
	"IDENT",
	"INT",
	"ASSIGN",
	"PLUS",
	"COMMA",
	"SEMICOLON",
	"LPAREN",
	"RPAREN",
	"LBRACE",
	"RBRACE",
	"FUNCTION",
	"LET",
}

const (
	ILLEGAL = iota
	EOF     = iota

	// Identifiers & literals
	IDENT = iota
	INT   = iota

	// Operators
	ASSIGN = iota
	PLUS   = iota

	// Delimiters
	COMMA     = iota
	SEMICOLON = iota

	LPAREN = iota
	RPAREN = iota
	LBRACE = iota
	RBRACE = iota

	// Keywords
	FUNCTION = iota
	LET      = iota
)

func TokenTypeToString(t TokenType) string {
	return tokenNames[t]
}
