/*
You are given an input string.

For each symbol in the string if it's the first character occurrence, replace it with a '1', else replace it with the amount of times you've already seen it...

But will your code be performant enough?

Examples:
input   =  "Hello, World!"
result  =  "1112111121311"

input   =  "aaaaaaaaaaaa"
result  =  "123456789101112"
There might be some non-ascii characters in the string.*/
package kata

import (
  "strings"
  "strconv"
)

func Numericals(s string) string {
    var m = make(map[rune]int)
  var nums []int
    for _, v := range s {
      m[v] += 1
      nums = append(nums, m[v])
    }
  var r strings.Builder
  for _, n := range nums {
    str := strconv.Itoa(n)
    r.WriteString(str)
  }
  return r.String()
}