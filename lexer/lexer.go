package lexer

// TODO: Change ch to support full unicode and not just UTF-8
type Lexer struct {
	input        string
	position     int  // Current pos in input (current char)
	readPosition int  // Current reading pos (after current char)
	ch           byte // Current char
}

func new(input string) *Lexer {
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
