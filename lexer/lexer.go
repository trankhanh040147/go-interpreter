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

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.input[l.readPosition]) {
		l.readChar()
	}
	return l.input[position:l.readPosition]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isNumber(l.input[l.readPosition]) {
		l.readChar()
	}
	return l.input[position:l.readPosition]
}

func (l *Lexer) skipSpacingCharacter() {
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
	l.skipSpacingCharacter()
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
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.MULTIPLY, l.ch)
	case '/':
		tok = newToken(token.DIV, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookUpIdentType(tok.Literal)
			break
		}
		if isNumber(l.ch) {
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

func isNumber(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
