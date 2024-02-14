package lexer

import "monkey/token"

type Lexer struct {
    input        string
    position     int     // Lexer position in input
    readPosition int     // Position after position
    ch           byte    // char at position
}

func New(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar() 
    return l
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

func (l *Lexer) NextToken() token.Token {
    var tok token.Token

    l.skipWhitespace()

    switch l.ch {
    case '=':
        tok = newToken(token.ASSIGN, l.ch)
    case ';':
        tok = newToken(token.SEMICOLON, l.ch)
    case '(':
        tok = newToken(token.LPAREN, l.ch)
    case ')':
        tok = newToken(token.RPAREN, l.ch)
    case ',':
        tok = newToken(token.COMMA, l.ch)
    case '+':
        tok = newToken(token.PLUS, l.ch)
    case '{':
        tok = newToken(token.LBRACE, l.ch)
    case '}':
        tok = newToken(token.RBRACE, l.ch)
    case 0:
        tok = token.Token{Type: token.EOF, Literal: ""}
    default:
        if isLetter(l.ch) {
            tok.Literal = l.readIdentifier()
            tok.Type = LookupIdent(tok.Literal)
            return tok
        } else if isDigit(l.ch) {
            tok.Literal = l.readInt()
            tok.Type = token.INT
            return tok
        } else {
            tok = newToken(token.ILLEGAL, l.ch)
        }
    }

    l.readChar()
    return tok
}

func (l *Lexer) readIdentifier() string {
    var startPos = l.position
    for isLetter(l.ch) {
        l.readChar()
    }
    return l.input[startPos:l.position]
}

func (l *Lexer) readInt() string {
    var startPos = l.position
    for isDigit(l.ch) {
        l.readChar()
    }
    return l.input[startPos:l.position]
}

var keywords = map[string]token.TokenType{
    "fn": token.FUNCTION,
    "let": token.LET,
}

func LookupIdent(ident string) token.TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return token.IDENT
}

func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
    for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
        l.readChar()
    }
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
    return token.Token{Type: tokenType, Literal: string(ch)}
}
