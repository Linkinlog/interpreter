package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// IDENT Identifiers add, foobar, x, y, ...
	IDENT = "IDENT"
	// INT literals 12345
	INT = "INT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="

	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LSQUIGGLE = "{"
	RSQUIGGLE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var TokenTypes = map[byte]TokenType{
	'=': ASSIGN,
	';': SEMICOLON,
	'(': LPAREN,
	')': RPAREN,
	',': COMMA,
	'+': PLUS,
	'{': LSQUIGGLE,
	'}': RSQUIGGLE,
	'-': MINUS,
	'!': BANG,
	'*': ASTERISK,
	'/': SLASH,
	'<': LT,
	'>': GT,
	0:   EOF,
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if toke, ok := keywords[ident]; ok {
		return toke
	}
	return IDENT
}
