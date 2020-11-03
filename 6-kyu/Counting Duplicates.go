/*
Count the number of Duplicates
Write a function that will return the count of distinct case-insensitive alphabetic characters and numeric digits that occur more than once in the input string. The input string can be assumed to contain only alphabets (both uppercase and lowercase) and numeric digits.

Example
"abcde" -> 0 # no characters repeats more than once
"aabbcde" -> 2 # 'a' and 'b'
"aabBcde" -> 2 # 'a' occurs twice and 'b' twice (`b` and `B`)
"indivisibility" -> 1 # 'i' occurs six times
"Indivisibilities" -> 2 # 'i' occurs seven times and 's' occurs twice
"aA11" -> 2 # 'a' and '1'
"ABBA" -> 2 # 'A' and 'B' each occur twice
*/
package kata

import "strings"

func duplicate_count(s1 string) int {
  var m = make(map[rune]int)
  var s1Lower = strings.ToLower(s1)
  for _, v := range s1Lower {
    m[v] += 1
  }
  var chars []rune
  for k, v := range m {
    if v > 1 {
      chars = append(chars, k)
    }
  }
  return len(chars)
}
