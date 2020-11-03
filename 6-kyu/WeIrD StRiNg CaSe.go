/*
Write a function toWeirdCase (weirdcase in Ruby) that accepts a string, and returns the same string with all even indexed characters in each word upper cased, and all odd indexed characters in each word lower cased. The indexing just explained is zero based, so the zero-ith index is even, therefore that character should be upper cased.

The passed in string will only consist of alphabetical characters and spaces(' '). Spaces will only be present if there are multiple words. Words will be separated by a single space(' ').

Examples:
toWeirdCase("String") // => returns "StRiNg"
toWeirdCase("Weird string case") // => returns "WeIrD StRiNg CaSe"
*/
package kata

import (
  "strings"
)

func toWeirdCase(str string) string {
  var r = []string{}
  var lowerStr = strings.ToLower(str)
  var strArr = strings.Split(lowerStr, " ")
  for _, v := range strArr {
    tmp := ""
    for k, vv := range v {
      if k == 0 || k % 2 == 0 {
        tmp += strings.ToUpper(string(vv))
      } else {
        tmp += string(vv)
      }
    }
    r = append(r, tmp)
  }
  return strings.Join(r, " ")
}