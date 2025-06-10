package lexer

type Lexer struct {
	curpos   int8 //currently the file size cannot exceed 256 lines
	ch       byte
	filename string
}

func Init(fname string) Lexer {
	return Lexer{0, 0, fname}
}

func NextChar() byte {
	return 0
}
