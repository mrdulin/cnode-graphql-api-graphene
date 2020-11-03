/*
Complete the method which returns the number which is most frequent in the given input array. If there is a tie for most frequent number, return the largest number among them.

Note: no empty arrays will be given.

Examples
[12, 10, 8, 12, 7, 6, 4, 10, 12]              -->  12
[12, 10, 8, 12, 7, 6, 4, 10, 12, 10]          -->  12
[12, 10, 8, 8, 3, 3, 3, 3, 2, 4, 10, 12, 10]  -->   3
*/
package kata

func HighestRank(nums []int) int {
  var m = make(map[int]int)
  for _, n := range nums {
    m[n] += 1
  }
  var max, count int
  for n, v := range m {
    if v > count {
      count = v
      max = n
    } else if v == count && n > max {
      max = n
    }
  }
  return max
}