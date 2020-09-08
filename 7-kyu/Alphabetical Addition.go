package __kyu

import (
	"unicode/utf8"
)

func AddLetters(letters []rune) rune {
	if len(letters) == 0 {
		r, _ := utf8.DecodeRuneInString("z")
		return r
	}
	var (
		t, tt int
		base  = 96
	)
	for _, letter := range letters {
		t += int(letter) - base
	}
	if tt = t % 26; tt == 0 {
		tt = 26
	}
	return rune(tt + base)
}
