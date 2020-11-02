/*
Write a function that when given a number >= 0, returns an Array of ascending length subarrays.

pyramid(0) => [ ]
pyramid(1) => [ [1] ]
pyramid(2) => [ [1], [1, 1] ]
pyramid(3) => [ [1], [1, 1], [1, 1, 1] ]
Note: the subarrays should be filled with 1s
*/
package kata

func Pyramid(n int) [][]int {
  var r = [][]int{}
  var el = []int{}
  for {
    if len(r) == n {
      break
    }
    el = append(el, 1)
    r = append(r, el)
  }
  return r
}