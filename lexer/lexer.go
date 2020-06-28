package lexer

type Lexer struct {
	input string
}

func MakeLexer(input string) Lexer {
	return Lexer{input: input}
}

func (l Lexer) NextToken() {

}
