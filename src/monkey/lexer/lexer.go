package lexer

type Lexer struct {
    input        string
    position     int     // Lexer position in input
    readPosition int     // Position after position
    ch           byte    // char at position
}

func New(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar() // @NOTE: How does this initialize Lexer state?
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


