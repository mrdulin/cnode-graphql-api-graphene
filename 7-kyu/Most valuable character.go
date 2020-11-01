/*
In this Kata, you will be given a string and your task is to return the most valuable character. The value of a character is the difference between the index of its last occurrence and the index of its first occurrence. Return the character that has the highest value. If there is a tie, return the alphabetically lowest character. [For Golang return rune]

All inputs will be lower case.

For example:
solve('a') = 'a'
solve('ab') = 'a'. Last occurrence is equal to first occurrence of each character. Return lexicographically lowest.
solve("axyzxyz") = 'x'
More examples in test cases. Good luck!
*/
package kata

import (
  "strings"
  "sort"
)

type Tuple struct {
  R rune
  Diff int
}

func Solve(s string) rune {
  arr := []Tuple{}
  for i, r := range s {
    lastIdx := strings.LastIndex(s, string(r))
    arr = append(arr, Tuple{R: r, Diff: lastIdx - i})
  }

  sort.Slice(arr, func(i, j int) bool {
    if arr[i].Diff > arr[j].Diff {
      return true
    }
    if arr[i].Diff < arr[j].Diff {
      return false
    }
    return arr[i].R < arr[j].R
  })
  return arr[0].R
}