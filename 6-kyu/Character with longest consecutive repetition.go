/*
For a given string s find the character c (or C) with longest consecutive repetition and return:

type Result struct {
    C rune // character
    L int  // count
}
where l (or L) is the length of the repetition. If there are two or more characters with the same l return the first in order of appearance.

For empty string return:

Happy coding! :)

*/
package kata

import (
  "sort"
  "fmt"
)

func LongestRepetition(text string) Result {
  if len(text) == 0 {
    return Result{}
  }
  results := []Result{}
  i := 0
  for j, r := range text {
    if j == 0 {
      results = append(results, Result{C: r, L: 1})
    } else if string(results[i].C) == string(r) {
      results[i].L++
    } else {
      results = append(results, Result{C: r, L: 1})
      i ++  
    }
  }
  
  sort.SliceStable(results, func(i, j int) bool {
    if results[i].L > results[j].L {
      return true
    }
    if results[i].L < results[j].L {
      return false
    }
    return results[i].C == results[j].C
  })
  fmt.Println(text, results)
  return results[0]
}