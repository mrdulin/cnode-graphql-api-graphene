/*
The input is a string str of digits. Cut the string into chunks (a chunk here is a substring of the initial string) of size sz (ignore the last chunk if its size is less than sz).

If a chunk represents an integer such as the sum of the cubes of its digits is divisible by 2, reverse that chunk; otherwise rotate it to the left by one position. Put together these modified chunks and return the result as a string.

If

sz is <= 0 or if str is empty return ""
sz is greater (>) than the length of str it is impossible to take a chunk of size sz hence return "".
Examples:
revrot("123456987654", 6) --> "234561876549"
revrot("123456987653", 6) --> "234561356789"
revrot("66443875", 4) --> "44668753"
revrot("66443875", 8) --> "64438756"
revrot("664438769", 8) --> "67834466"
revrot("123456779", 8) --> "23456771"
revrot("", 8) --> ""
revrot("123456779", 0) --> "" 
revrot("563000655734469485", 4) --> "0365065073456944"
*/
package kata

import (
  "math"
)

func Revrot(s string, n int) string {
  if len(s) == 0 || n <= 0 || n > len(s) {
    return ""
  }
  var r string
  ss := ToChunk(s, n)
  for _, substr := range ss {
    if CubeSumOfDigit(substr)%2 == 0 {
      r += Reverse(substr)
    } else {
      r += substr[1:] + string(substr[0])
    }
  }
  return r
}


func ToChunk(s string, n int) []string {
  r := []string{}
  for i := 0; i < len(s); i += n {
    if i+n > len(s) {
      break
    }
    r = append(r, s[i:i+n])
  }
  return r
}


func CubeSumOfDigit(s string) int {
  var sum int
  for _, r := range s {
    sum += int(math.Pow(float64(r), 3))
  }
  return sum
}

func Reverse(s string) (result string) {
  for _, v := range s {
    result = string(v) + result
  }
  return
}