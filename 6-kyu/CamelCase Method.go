/*
Write simple .camelCase method (camel_case function in PHP, CamelCase in C# or camelCase in Java) for strings. All words must have their first letter capitalized without spaces.

For instance:

CamelCase("hello case")      // => "HelloCase"
CamelCase("camel case word") // => "CamelCaseWord"
Don't forget to rate this kata! Thanks :)
*/
package kata

import (
  "strings"
)

func CamelCase(s string) string {
  s = strings.TrimSpace(s)
  if len(s) == 0 {
    return s
  }
  r := []string{}
  ss := strings.Split(s, " ")
  for _, substr := range ss {
    r = append(r, strings.ToUpper(string(substr[0]))+substr[1:])
  }
  return strings.Join(r, "")
}