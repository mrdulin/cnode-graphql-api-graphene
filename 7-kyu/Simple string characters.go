/*
In this Kata, you will be given a string and your task will be to return a list of ints detailing the count of uppercase letters, lowercase, numbers and special characters, as follows.

Solve("*'&ABCDabcde12345") = [4,5,5,3]. 
--the order is: uppercase letters, lowercase, numbers and special characters.
More examples in the test cases.
*/
package kata

import "regexp"

var (
  nRe = regexp.MustCompile(`[0-9]`)
  upperRe = regexp.MustCompile(`[A-Z]`)
  lowerRe = regexp.MustCompile(`[a-z]`)
)


func Solve(s string) []int {
  uc := len(upperRe.FindAllString(s, -1))
  nc := len(nRe.FindAllString(s, -1))
  lc := len(lowerRe.FindAllString(s, -1))
  sc := len(s) - (uc + nc + lc)
  return []int{uc, lc, nc, sc}
}