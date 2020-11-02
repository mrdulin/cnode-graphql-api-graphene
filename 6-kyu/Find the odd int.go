/*
Given an array of integers, find the one that appears an odd number of times.

There will always be only one integer that appears an odd number of times.
*/
package kata

func FindOdd(seq []int) int {
  var m = make(map[int]int)
  for _, v := range seq {
    m[v] += 1
  }
  var r int
  for k, v := range m {
    if v % 2 != 0 {
      r = k
      break
    }
  }
  return r
}