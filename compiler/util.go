package compiler

import "unicode"

func isDigit(_ int, r rune) bool {
	return unicode.IsDigit(r)
}

func isLetter(_ int, r rune) bool {
	return unicode.IsLetter(r)
}
