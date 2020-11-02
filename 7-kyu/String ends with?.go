/*
Complete the solution so that it returns true if the first argument(string) passed in ends with the 2nd argument (also a string).

Examples:

solution("abc", "bc") // returns true
solution("abc", "d") // returns false
*/
package kata

import (
  "strings"
)

func solution(str, ending string) bool {
  var idx = strings.LastIndex(str, ending)
  return idx != -1 && idx + len(ending) == len(str)
}