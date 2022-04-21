package compiler

import (
	"fmt"
)

type Line struct {
	Num        int
	Start, End int
	Content    []byte
}

func NewLine(num, start, end int, content []byte) Line {
	return Line{
		Num:     num,
		Start:   start,
		End:     end,
		Content: content,
	}
}

type CompilerError struct {
	Line Line
	At   int
}

type SyntaxErr struct {
	CompilerError
}

func NewSyntaxErr(ln Line, at int) SyntaxErr {
	return SyntaxErr{
		CompilerError: CompilerError{
			Line: ln,
			At:   at,
		},
	}
}

func (se SyntaxErr) Error() string {
	msg := fmt.Sprintf(
		"[SYNTAX ERROR]\nln %v|>> %v\n",
		se.Line.Num,
		string(se.Line.Content),
	)
	return msg
	// offset := se.At - se.Line.Start
	// return msg + strings.Join(make([]string, offset), " ") + "^"
}
