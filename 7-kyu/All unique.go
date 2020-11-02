/*
Write a program to determine if a string contains only unique characters. Return true if it does and false otherwise.

The string may contain any of the 128 ASCII characters. Characters are case-sensitive, e.g. 'a' and 'A' are considered different characters.
*/
package kata

import "strings"

func HasUniqueChar (str string) bool {
  for i, r := range str {
    if strings.Contains(str[i+1:], string(r)) {
      return false
    }
  }
  return true
}