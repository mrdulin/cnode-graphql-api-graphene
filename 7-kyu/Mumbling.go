/*
This time no story, no theory. The examples below show you how to write function accum:

Examples:

accum("abcd") -> "A-Bb-Ccc-Dddd"
accum("RqaEzty") -> "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
accum("cwAt") -> "C-Ww-Aaa-Tttt"
The parameter of accum is a string which includes only letters from a..z and A..Z.
*/
package kata

import (
  "strings"
  "fmt"
)

func Accum(s string) string {
  var strs []string
  for k, v := range s {
    str := strings.Repeat(strings.ToLower(string(v)), k + 1)
    str = strings.Title(str)
    fmt.Println(k, str)
    strs = append(strs, str)
  }
  return strings.Join(strs, "-")
}