/*
In this Kata, we are going to reverse a string while maintaining the spaces (if any) in their original place.

For example:

solve("our code") = "edo cruo"
-- Normal reversal without spaces is "edocruo". 
-- However, there is a space at index 3, so the string becomes "edo cruo"

solve("your code rocks") = "skco redo cruoy". 
solve("codewars") = "srawedoc"
More examples in the test cases. All input will be lower case letters and in some cases spaces.

Good luck!
*/
package kata

import (
  "strings"
)

func Solve(s string) string {
  r := make([]string, len(s))
  for i, v := range s {
    if string(v) == " " {
      r[i] = " "
    }
  }
  j := len(s) - 1
  for i, v := range s {
    if string(v) != " " {
      if r[j] == " " {
        j--
      }
      r[i] = string(s[j])
      j--
    }
  }
  return strings.Join(r, "")
}