/*
Write a function that takes in a string of one or more words, and returns the same string, but with all five or more letter words reversed (Just like the name of this Kata). Strings passed in will consist of only letters and spaces. Spaces will be included only when more than one word is present.

Examples: spinWords( "Hey fellow warriors" ) => returns "Hey wollef sroirraw" spinWords( "This is a test") => returns "This is a test" spinWords( "This is another test" )=> returns "This is rehtona test"
*/
package kata

import (
  "strings"
)

func SpinWords(str string) (rval string) {
  arr := strings.Split(str, " ")
  var rvalArr []string
  for _, word := range arr {
    if len(word) < 5 {
      rvalArr = append(rvalArr, word)
    } else {
      rvalArr = append(rvalArr, reverse(word))
    }
  }
  return strings.Join(rvalArr, " ")
}

func reverse(s string) (r string) {
  for _, v := range s {
    r = string(v) + r
  }
  return
}