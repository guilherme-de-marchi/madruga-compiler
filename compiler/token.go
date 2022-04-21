package compiler

import (
	"fmt"
)

const (
	// Single-character
	LeftParen = iota
	RightParen
	LeftBrace
	RightBrace
	Comma
	Dot
	Minus
	Plus
	SemiColon
	Slash
	Star
	Bang
	Equal
	Greater
	Less

	// Composite characters
	BangEqual
	EqualEqual
	GreaterEqual
	LessEqual

	// Literals
	Identifier
	String
	Number

	// Keywords
	And
	Or
	If
	Else
	False
	True
	Class
	Super
	This
	Func
	For
	Nil
	Print
	Return
	Var
	While

	Eof
)

type Token struct {
	Type    int
	Literal interface{}
	Line    int
}

func NewToken(tp int, lt interface{}, ln int) *Token {
	return &Token{
		Type:    tp,
		Literal: lt,
		Line:    ln,
	}
}

func (t *Token) String() string {
	return fmt.Sprintf("\nType ID: %v\nLiteral: %v", t.Type, t.Literal)
}

func PrintTokens(tks ...*Token) {
	for _, tk := range tks {
		fmt.Println(tk.String())
	}
}
