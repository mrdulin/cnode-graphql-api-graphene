/*
Given a string of words, you need to find the highest scoring word.

Each letter of a word scores points according to its position in the alphabet: a = 1, b = 2, c = 3 etc.

You need to return the highest scoring word as a string.

If two words score the same, return the word that appears earliest in the original string.

All letters will be lowercase and all inputs will be valid.
*/
package kata

import (
  "strings"
)

func High(s string) string {
  var (
    max int
    idx int
  )
  ss := strings.Split(s, " ")
  for i, substr := range ss {
    var sum int
    for _, r := range substr {
      sum += int(r) - 96
    }
    if sum > max {
      max = sum
      idx = i
    }
  }
  return ss[idx]
}