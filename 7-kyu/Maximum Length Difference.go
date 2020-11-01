/*
You are given two arrays a1 and a2 of strings. Each string is composed with letters from a to z. Let x be any string in the first array and y be any string in the second array.

Find max(abs(length(x) − length(y)))

If a1 and/or a2 are empty return -1 in each language except in Haskell (F#) where you will return Nothing (None).

Example:
a1 = ["hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"]
a2 = ["cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"]
mxdiflg(a1, a2) --> 13
Bash note:
input : 2 strings with substrings separated by ,
output: number as a string
*/
package kata

import "math"

func MxDifLg(a1 []string, a2 []string) int {
  var maxDiff = -1
  for _, x := range a1 {
    xLen := len(x)
    if xLen == 0 {
      break;
    }
    for _, y := range a2 {
      yLen := len(y)
      if yLen == 0 { 
        break;
      }
      diff := int(math.Abs(float64(xLen - yLen)))
      if maxDiff < diff {
        maxDiff = diff
      }
    }
  }
  return maxDiff
}