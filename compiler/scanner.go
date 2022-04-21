package compiler

import (
	"strconv"
	"unicode"
)

type Source []byte

func (src Source) GetLine(ln int) Line {
	var counter int
	start, end := -1, -1
	if ln == 1 {
		start = 0
	}
	for i, v := range src {
		if start == -1 && counter == ln-1 {
			start = i
			continue
		}
		if v == '\n' {
			counter++
		}
		if counter == ln {
			end = i
			break
		}
	}
	if start != -1 {
		if end != -1 {
			return NewLine(ln, start, end, src[start:end])
		}
		return NewLine(ln, start, len(src), src[start:])
	}
	return NewLine(ln, 0, 0, []byte{})
}

func (src Source) NextMatch(curr int, exp rune) bool {
	return rune(src[curr+1]) == exp
}

func (src Source) NextFunc(curr int, f func(rune) bool) bool {
	return f(rune(src[curr+1]))
}

// FindNextOccur returns -1 if it doesn't find.
// Otherwise, return the index of the next occurrence of exp.
func (src Source) FindNextOccur(curr int, exp rune) int {
	for i, v := range src[curr+1:] {
		if rune(v) == exp {
			return curr + i
		}
	}
	return -1
}

// IterateFunc return the first index where all functions return false
func (src Source) IterateFunc(curr int, arrF ...func(int, rune) bool) int {
	for i, v := range src[curr:] {
		var ok bool
		for _, f := range arrF {
			if !f(curr+i, rune(v)) {
				ok = false
				continue
			}
			ok = true
			break
		}
		if !ok {
			return curr + i
		}
	}
	return len(src)
}

func (src Source) Scan() ([]*Token, error) {
	tks := make([]*Token, 0)
	var curr int
	ln := 1 // line counter

	validFloatPoint := func(i int, r rune) bool {
		return r == '.' && unicode.IsDigit(rune(src[i+1]))
	}

	for curr < len(src) {
		// Literals
		if unicode.IsDigit(rune(src[curr])) {
			i := src.IterateFunc(curr, isDigit, validFloatPoint)
			v, err := strconv.ParseFloat(string(src[curr:i]), 32)
			if err != nil {
				return tks, NewSyntaxErr(src.GetLine(ln), curr)
			}
			tks = append(tks, NewToken(Number, v, ln))
			curr = i
			continue
		}
		if unicode.IsLetter(rune(src[curr])) {
			i := src.IterateFunc(curr, isLetter)
			if i == -1 {
				return tks, NewSyntaxErr(src.GetLine(ln), curr)
			}
			switch v := string(src[curr:i]); v {
			case "and":
				tks = append(tks, NewToken(And, nil, ln))
			case "or":
				tks = append(tks, NewToken(Or, nil, ln))
			case "if":
				tks = append(tks, NewToken(If, nil, ln))
			case "else":
				tks = append(tks, NewToken(Else, nil, ln))
			case "false":
				tks = append(tks, NewToken(False, nil, ln))
			case "true":
				tks = append(tks, NewToken(True, nil, ln))
			case "class":
				tks = append(tks, NewToken(Class, nil, ln))
			case "super":
				tks = append(tks, NewToken(Super, nil, ln))
			case "this":
				tks = append(tks, NewToken(This, nil, ln))
			case "func":
				tks = append(tks, NewToken(Func, nil, ln))
			case "for":
				tks = append(tks, NewToken(For, nil, ln))
			case "nil":
				tks = append(tks, NewToken(Nil, nil, ln))
			case "print":
				tks = append(tks, NewToken(Print, nil, ln))
			case "return":
				tks = append(tks, NewToken(Return, nil, ln))
			case "var":
				tks = append(tks, NewToken(Var, nil, ln))
			case "while":
				tks = append(tks, NewToken(While, nil, ln))
			default:
				tks = append(tks, NewToken(Identifier, v, ln))
			}
			curr = i
			continue
		}

		switch src[curr] {
		// Escape codes.
		case ' ':
		case '\r':
		case '\t':
			break
		case '\n':
			ln++

		// Single-characters.
		case '(':
			tks = append(tks, NewToken(LeftParen, nil, ln))
		case ')':
			tks = append(tks, NewToken(RightParen, nil, ln))
		case '{':
			tks = append(tks, NewToken(LeftBrace, nil, ln))
		case '}':
			tks = append(tks, NewToken(RightBrace, nil, ln))
		case ',':
			tks = append(tks, NewToken(Comma, nil, ln))
		case '-':
			tks = append(tks, NewToken(Minus, nil, ln))
		case '+':
			tks = append(tks, NewToken(Plus, nil, ln))
		case ';':
			tks = append(tks, NewToken(SemiColon, nil, ln))
		case '/':
			tks = append(tks, NewToken(Slash, nil, ln))
		case '*':
			tks = append(tks, NewToken(Star, nil, ln))

		// Composite characters
		case '!':
			if src.NextMatch(curr, '=') {
				tks = append(tks, NewToken(BangEqual, nil, ln))
				curr++
				break
			}
			tks = append(tks, NewToken(Bang, nil, ln))
		case '=':
			if src.NextMatch(curr, '=') {
				tks = append(tks, NewToken(EqualEqual, nil, ln))
				curr++
				break
			}
			tks = append(tks, NewToken(Equal, nil, ln))
		case '>':
			if src.NextMatch(curr, '=') {
				tks = append(tks, NewToken(GreaterEqual, nil, ln))
				curr++
				break
			}
			tks = append(tks, NewToken(Greater, nil, ln))
		case '<':
			if src.NextMatch(curr, '=') {
				tks = append(tks, NewToken(LessEqual, nil, ln))
				curr++
				break
			}
			tks = append(tks, NewToken(Less, nil, ln))

		// Literals
		case '.':
			// Abbreviation for float lower than 1. Example: .5 == 0.5
			if src.NextFunc(curr, unicode.IsDigit) {
				curr++
				i := src.IterateFunc(curr, isDigit)
				v, err := strconv.ParseFloat(string(src[curr-1:i]), 32)
				if err != nil {
					return tks, NewSyntaxErr(src.GetLine(ln), curr)
				}
				tks = append(tks, NewToken(Number, v, ln))
				curr = i
				break
			}
			tks = append(tks, NewToken(Dot, nil, ln))
		case '"':
			i := src.FindNextOccur(curr, '"')
			if i == -1 {
				return tks, NewSyntaxErr(src.GetLine(ln), curr)
			}
			tks = append(tks, NewToken(String, string(src[curr+1:i+1]), ln))
			curr = i + 1
		}
		curr++
	}
	tks = append(tks, NewToken(Eof, nil, ln))
	return tks, nil
}
