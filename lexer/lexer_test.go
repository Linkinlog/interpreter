package lexer

import (
	"reflect"
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

func TestNew(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want *Lexer
	}{
		{
			name: "Test_New_01",
			args: args{input: "let foo = 5;"},
			want: &Lexer{
				input:        "let foo = 5;",
				position:     0,
				readPosition: 1,
				char:         byte('l'),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_readChar(t *testing.T) {
	type fields struct {
		input        string
		position     int
		readPosition int
		char         byte
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Test_readChar_01",
			fields: fields{
				input:        "let foo = 5;",
				position:     0,
				readPosition: 1,
				char:         byte('l'),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lexer{
				input:        tt.fields.input,
				position:     tt.fields.position,
				readPosition: tt.fields.readPosition,
				char:         tt.fields.char,
			}
			l.readChar()
		})
	}
}

func TestLexer_NextToken(t *testing.T) {
	type fields struct {
		input        string
		position     int
		readPosition int
		char         byte
	}
	tests := []struct {
		name     string
		fields   fields
		wantToke token.Token
	}{
		{
			name: "Test_NextToken_ILLEGAL",
			fields: fields{
				input:        "!foo",
				position:     0,
				readPosition: 1,
				char:         byte('!'),
			},
			wantToke: token.Token{
				Type:    token.ILLEGAL,
				Literal: "!",
			},
		},
		{
			name: "Test_NextToken_IDENT",
			fields: fields{
				input:        "foo",
				position:     0,
				readPosition: 1,
				char:         byte('f'),
			},
			wantToke: token.Token{
				Type:    token.IDENT,
				Literal: "foo",
			},
		},
		{
			name: "Test_NextToken_IDENT",
			fields: fields{
				input:        "let",
				position:     0,
				readPosition: 1,
				char:         byte('l'),
			},
			wantToke: token.Token{
				Type:    token.LET,
				Literal: "let",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lexer{
				input:        tt.fields.input,
				position:     tt.fields.position,
				readPosition: tt.fields.readPosition,
				char:         tt.fields.char,
			}
			if gotToke := l.NextToken(); !reflect.DeepEqual(gotToke, tt.wantToke) {
				t.Errorf("Lexer.NextToken() = %v, want %v", gotToke, tt.wantToke)
			}
		})
	}
}

func TestLexer_readNumberOrIdentifier(t *testing.T) {
	type fields struct {
		input        string
		position     int
		readPosition int
		char         byte
	}
	tests := []struct {
		name   string
		fields fields
		fn     func(byte) bool
		want   string
	}{
		{
			name: "Test_readNumberOrIdentifier_FUNCTION",
			fields: fields{
				input:        "function",
				position:     0,
				readPosition: 1,
				char:         byte('f'),
			},
			fn:   isLetter,
			want: "function",
		},
		{
			name: "Test_readNumberOrIdentifier_LET",
			fields: fields{
				input:        "let",
				position:     0,
				readPosition: 1,
				char:         byte('l'),
			},
			fn:   isLetter,
			want: "let",
		},
		{
			name: "Test_readNumber_5",
			fields: fields{
				input:        "5",
				position:     0,
				readPosition: 1,
				char:         byte('5'),
			},
			fn:   isDigit,
			want: "5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lexer{
				input:        tt.fields.input,
				position:     tt.fields.position,
				readPosition: tt.fields.readPosition,
				char:         tt.fields.char,
			}
			if got := l.readNumberOrIdentifier(tt.fn); got != tt.want {
				t.Errorf("Lexer.readNumberOrIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newToken(t *testing.T) {
	type args struct {
		tokenType token.TokenType
		char      byte
	}
	tests := []struct {
		name string
		args args
		want token.Token
	}{
		{
			name: "Test_newToken_ASSIGN",
			args: args{
				tokenType: token.ASSIGN,
				char:      '=',
			},
			want: token.Token{
				Type:    token.ASSIGN,
				Literal: "=",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newToken(tt.args.tokenType, tt.args.char); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isLetter(t *testing.T) {
	type args struct {
		char byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test_isLetter_5",
			args: args{
				char: '5',
			},
			want: false,
		},
		{
			name: "Test_isLetter_a",
			args: args{
				char: 'a',
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLetter(tt.args.char); got != tt.want {
				t.Errorf("isLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isDigit(t *testing.T) {
	type args struct {
		char byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test_isDigit_a",
			args: args{
				char: 'a',
			},
			want: false,
		},
		{
			name: "Test_isDigit_0",
			args: args{
				char: '0',
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDigit(tt.args.char); got != tt.want {
				t.Errorf("isDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_skipWhitespace(t *testing.T) {
	type fields struct {
		input        string
		position     int
		readPosition int
		char         byte
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Test_SkipWhitespace_01",
			fields: fields{
				input:        "let     foo = 5;",
				position:     0,
				readPosition: 1,
				char:         byte('l'),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lexer{
				input:        tt.fields.input,
				position:     tt.fields.position,
				readPosition: tt.fields.readPosition,
				char:         tt.fields.char,
			}
			l.skipWhitespace()
		})
	}
}
