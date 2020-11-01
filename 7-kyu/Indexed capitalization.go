/*
Given a string and an array of integers representing indices, capitalize all letters at the given indices.

For example:

capitalize("abcdef",[1,2,5]) = "aBCdeF"
capitalize("abcdef",[1,2,5,100]) = "aBCdeF". There is no index 100.
The input will be a lowercase string with no spaces and an array of digits.

Good luck!
*/
package kata

import (
  "strings"
)

func Capitalize(s string, arr []int) string {
  var res string
  for i, r := range s {
    if IntContains(arr, i) {
      res += strings.ToUpper(string(r))
    } else {
      res += string(r)
    }
  }
  return res
}

func IntContains(arr []int, i int) bool {
  for _, s := range arr {
    if s == i {
      return true
    }
  }
  return false
}