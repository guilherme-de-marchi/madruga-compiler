package compiler

type Lexer interface {
	Scan() ([]*Token, error)
}

type Compiler struct {
	Lexer Lexer
}

func NewCompiler(lexer Lexer) *Compiler {
	return &Compiler{
		Lexer: lexer,
	}
}
