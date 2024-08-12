package lexer

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  //current reading position in input (after current char)
	ch           byte //Current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

// Exercise: change to support unicode and emojis.
func (l *Lexer) readCHar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
