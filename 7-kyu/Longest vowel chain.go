/*
The vowel substrings in the word codewarriors are o,e,a,io. The longest of these has a length of 2. Given a lowercase string that has alphabetic characters only (both vowels and consonants) and no spaces, return the length of the longest vowel substring. Vowels are any of aeiou.

Good luck!
*/
package kata

import (
  "strings"
)

func Solve(s string) int {
  f := func(c rune) bool {
    return !strings.Contains("aeiou", string(c))
  }
  vowels := strings.FieldsFunc(s, f)
  var l int
  for _, v := range vowels {
    if len(v) > l {
      l = len(v)
    }
  }
  return l
}