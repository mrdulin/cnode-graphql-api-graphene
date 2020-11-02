/*
Given a string, capitalize the letters that occupy even indexes and odd indexes separately, and return as shown below. Index 0 will be considered even.

For example, capitalize("abcdef") = ['AbCdEf', 'aBcDeF']. See test cases for more examples.

The input will be a lowercase string with no spaces.
*/
package kata

import "strings"

func Capitalize(st string) []string {
  var (
    odd, even string
  )
  for i, r := range st {
    if i%2 == 0 {
      even += string(r)
      odd += strings.ToUpper(string(r))
    } else {
      even += strings.ToUpper(string(r))
      odd += string(r)
    }
  }
  return []string{odd, even}
}