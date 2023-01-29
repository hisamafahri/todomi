package token

var keywords = map[string]TokenType {
    "fn": FUNCTION,
    "let": LET,
}

func LookupIdentifier(identifier string) TokenType {
    if tok, ok := keywords[identifier]; ok {
       return tok 
    }
    return IDENTIFIER
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	IDENTIFIER = "IDENTIFIER"
	INTEGER    = "INTEGER"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

