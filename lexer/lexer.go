package lexer

import "github.com/Linkinlog/MagLang/token"

// Lexer struct 
type Lexer struct {
	input        string
	position     int  // current position we are at in the input
	readPosition int  // position we will be reading from
	char         byte // current char we are examining
}

// New function 
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar method 
// Sets Lexer.char to be the character at Lexer.readPosition.
// At the EOF we set Lexer.char to "NUL".
func (l *Lexer) readChar() {
	eof_byte := byte(0)
	if l.readPosition >= len(l.input) {
		l.char = eof_byte
	} else {
		l.char = l.input[l.readPosition]
	}
	// advance position and readPosition
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() (toke token.Token) {
	l.skipWhitespace()
	tokenType, ok := token.TokenTypes[l.char]
	if !ok && (!isDigit(l.char) && !isLetter(l.char)) {
		return newToken(token.ILLEGAL, l.char)
	}
	if isLetter(l.char) {
		toke.Literal = l.readIdentifier()
		toke.Type = token.LookupIdent(toke.Literal)
		return toke
	}
	if isDigit(l.char) {
		toke.Type = token.INT
		toke.Literal = l.readNumber()
		return toke
	}
	if tokenType == token.EOF {
		toke.Literal = ""
		toke.Type = tokenType
		return toke
	}
	toke = newToken(tokenType, l.char)
	l.readChar()
	return toke
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}
