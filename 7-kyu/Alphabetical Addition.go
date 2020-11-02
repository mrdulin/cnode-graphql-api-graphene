/*
Your task is to add up letters to one letter.

The function will be given a slice of letters (runes), each one being a letter to add. Return a rune.

Notes:
Letters will always be lowercase.
Letters can overflow (see second to last example of the description)
If no letters are given, the function should return 'z'
Examples:
AddLetters([]rune{'a', 'b', 'c'}) = 'f'
AddLetters([]rune{'a', 'b'}) = 'c'
AddLetters([]rune{'z'}) = 'z'
AddLetters([]rune{'z', 'a'}) = 'a'
AddLetters([]rune{'y', 'c', 'b'}) = 'd' // notice the letters overflowing
AddLetters([]rune{}) = 'z'

*/
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
