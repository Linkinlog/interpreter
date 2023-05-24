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

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) NextToken() (toke token.Token) {
	l.skipWhitespace()
	tokenType, ok := token.TokenTypes[l.char]
	if !ok && (!isDigit(l.char) && !isLetter(l.char)) {
		return newToken(token.ILLEGAL, l.char)
	}
	if isTwoCharToken(l.char, l.peekChar()) {
		toke = makeTwoCharToken(l.char, l.peekChar())
		l.readChar()
		l.readChar()
		return toke
	}
	if isLetter(l.char) {
		toke.Literal = l.readNumberOrIdentifier(isLetter)
		toke.Type = token.LookupIdent(toke.Literal)
		return toke
	}
	if isDigit(l.char) {
		toke.Type = token.INT
		toke.Literal = l.readNumberOrIdentifier(isDigit)
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

func (l *Lexer) readNumberOrIdentifier(fn func(byte) bool) string {
	position := l.position
	for fn(l.char) {
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

func isTwoCharToken(char byte, next byte) bool {
	if (char == '=' || char == '!') && next == '=' {
		return true
	}
	return false
}

func makeTwoCharToken(first byte, second byte) (toke token.Token) {
	toke.Literal = string(first) + string(second)
	if first == '!' {
		toke.Type = token.NOT_EQ
	} else if first == '=' {
		toke.Type = token.EQ
	}
	return toke
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}
