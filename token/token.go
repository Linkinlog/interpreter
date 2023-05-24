package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// IDENT Identifiers add, foobar, x, y, ...
	IDENT = "IDENT"
	// INT literals 12345
	INT = "INT"

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LSQUIGGLE = "{"
	RSQUIGGLE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
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
	0:   EOF,
}

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if toke, ok := keywords[ident]; ok {
		return toke
	}
	return IDENT
}
