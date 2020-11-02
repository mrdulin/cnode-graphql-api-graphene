/*
Write a function called that takes a string of parentheses, and determines if the order of the parentheses is valid. The function should return true if the string is valid, and false if it's invalid.

Examples
"()"              =>  true
")(()))"          =>  false
"("               =>  false
"(())((()())())"  =>  true
Constraints
0 <= input.length <= 100
*/
package kata

import (
  "regexp"
)

func ValidParentheses(parens string) bool {
  var r = `\(\)`
  var re = regexp.MustCompile(r)
  parens = re.ReplaceAllString(parens, "");
  matched, _ := regexp.MatchString(r, parens)
  if(matched) {
    return ValidParentheses(parens)
  } else if len(parens) != 0 {
    return false
  }
  return true;
}