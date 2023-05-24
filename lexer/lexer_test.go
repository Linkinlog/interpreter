package lexer

import (
	"testing"

	"github.com/Linkinlog/MagLang/token"
)

func TestNextTokenSimple(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LSQUIGGLE, "{"},
		{token.RSQUIGGLE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	i := New(input)

	for idx, tt := range tests {
		token := i.NextToken()
		if token.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong, exptected=%q, received=%q",
				idx, tt.expectedType, token.Type)
		}

		if token.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - tokenLiteral wrong, exptected=%q, received=%q",
				idx, tt.expectedLiteral, token.Literal)
		}
	}
}

func TestNextTokenFull(t *testing.T) {
	input := `let five = 5;
	let ten = 10;

	let add = fn(x, y) {
		x + y;
	};

	let result = add(five, ten);
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LSQUIGGLE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RSQUIGGLE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	i := New(input)

	for idx, tt := range tests {
		token := i.NextToken()
		if token.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong, exptected=%q, received=%q",
				idx, tt.expectedType, token.Type)
		}

		if token.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - tokenLiteral wrong, exptected=%q, received=%q",
				idx, tt.expectedLiteral, token.Literal)
		}
	}
}
