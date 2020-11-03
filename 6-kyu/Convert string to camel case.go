/*
Complete the method/function so that it converts dash/underscore delimited words into camel casing. The first word within the output should be capitalized only if the original word was capitalized (known as Upper Camel Case, also often referred to as Pascal case).

Examples
ToCamelCase("the-stealth-warrior"); // returns "theStealthWarrior"

ToCamelCase("The_Stealth_Warrior"); // returns "TheStealthWarrior"
*/
package kata

import (
  "strings"
  "unicode"
)

func ToCamelCase(s string) string {
  var isToUpper = false
  var r string
  for k, v := range s {
    if k == 0 && unicode.IsUpper(v) {
      r += string(v)
      continue
    }
    if v == '_' || v == '-' {
      isToUpper = true
      continue
    } 
    if isToUpper {
      r += strings.ToUpper(string(v))
      isToUpper = false
    } else {
      r += string(v)
    }
  }
  return r
}