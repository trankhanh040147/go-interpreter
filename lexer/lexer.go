package lexer

import "go-interpreter/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
		return
	}
	l.position = l.readPosition
	l.readPosition++
	l.ch = l.input[l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for l.readPosition < len(l.input) && isLetter(l.input[l.readPosition]) {
		l.readChar()
	}
	return l.input[position:l.readPosition]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for l.readPosition < len(l.input) && isDigit(l.input[l.readPosition]) {
		l.readChar()
	}
	return l.input[position:l.readPosition]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
	return
}

func newToken(tokType token.TokType, ch byte) token.Token {
	return token.Token{
		Type:    tokType,
		Literal: string(ch),
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	switch l.ch {
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '=':
		ch := string(l.ch)
		tok.Type = token.ASSIGN
		tok.Literal = ch
		if l.peekChar() == '=' {
			l.readChar()
			tok.Type = token.EQ
			tok.Literal = ch + string(l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '!':
		ch := string(l.ch)
		tok.Type = token.BANG
		tok.Literal = string(l.ch)
		if l.peekChar() == '=' {
			l.readChar()
			tok.Type = token.NOT_EQ
			tok.Literal = ch + string(l.ch)
		}
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookUpIdentType(tok.Literal)
			break
		}
		if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			break
		}
		tok.Literal = ""
		tok.Type = token.ILLEGAL
	}
	l.readChar()
	return tok
}

func isLetter(ch byte) bool {
	return ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' || ch == '_'
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
