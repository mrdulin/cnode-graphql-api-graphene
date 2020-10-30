/*
In this Kata, we will check if a string contains consecutive letters as they appear in the English alphabet and if each letter occurs only once.

Rules are: (1) the letters are adjacent in the English alphabet, and (2) each letter occurs only once.

For example: 
solve("abc") = True, because it contains a,b,c
solve("abd") = False, because a, b, d are not consecutive/adjacent in the alphabet, and c is missing.
solve("dabc") = True, because it contains a, b, c, d
solve("abbc") = False, because b does not occur once.
solve("v") = True
All inputs will be lowercase letters.
*/
package kata

import (
  "sort"
)

func Solve(s string) bool {
  ints := []int{}
  for _, r := range s {
    ints = append(ints, int(r))
  }
  sort.Sort(sort.IntSlice(ints))
  var r int
  for i, c := range ints {
    if i == 0 {
      continue
    }
    r += int(c) - int(ints[i-1])
  }
  if r == len(s) - 1 {
    return true
  }
  return false
}