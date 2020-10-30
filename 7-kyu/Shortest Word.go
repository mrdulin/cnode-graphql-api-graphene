/*
Simple, given a string of words, return the length of the shortest word(s).

String will never be empty and you do not need to account for different data types.
*/
package kata

import (
  "strings"
  "fmt"
)

func FindShort(s string) int {
  strArr := strings.Split(s, " ")
  fmt.Println(strArr)
  var minLen int
  for i, subStr := range strArr {
    subStrLen := len(subStr)
    if i == 0 {
      minLen = subStrLen
    }
    if minLen > subStrLen {
      minLen = subStrLen
    }
  }
  return minLen
}