package lexer

import (
	"lexxy/tokens"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func (l *Lexer) currentChIs(ch byte) bool {
	return l.ch == ch
}

func (l *Lexer) peekChIs(ch byte) bool {
	return l.peekChar() == ch
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readField() string {
	pos := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(tokenType tokens.TokenType, ch byte) tokens.Token {
	return tokens.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

func (l *Lexer) NextToken() tokens.Token {
	var tok tokens.Token

	l.skipWhitespace()

	switch l.ch {
	case '(':
		tok = newToken(tokens.LPAREN, l.ch)
	case ')':
		tok = newToken(tokens.RPAREN, l.ch)
	case '[':
		tok = newToken(tokens.LBRACKET, l.ch)
	case ']':
		tok = newToken(tokens.RBRACKET, l.ch)
	case '{':
		tok = newToken(tokens.LBRACE, l.ch)
	case '}':
		tok = newToken(tokens.RBRACE, l.ch)
	case ',':
		tok = newToken(tokens.COMMA, l.ch)
	case ':':
		tok = newToken(tokens.LOL, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = tokens.EOF

	default:
		if isLetter(l.ch) {
			tok.Literal = l.readField()
			if l.ch == ':' {
				tok.Type = tokens.ROOT
			} else {
				tok.Type = tokens.FIELD
			}
			return tok
		} else {
			tok = newToken(tokens.ILLEGAL, l.ch)
		}

	}

	l.readChar()
	return tok
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}
