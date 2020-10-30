/*
The two oldest ages function/method needs to be completed. It should take an array of numbers as its argument and return the two highest numbers within the array. The returned value should be an array in the format [second oldest age, oldest age].

The order of the numbers passed in could be any order. The array will always include at least 2 items.

For example:

TwoOldestAges([]int{1, 5, 87, 45, 8, 8}) // should return [2]int{45, 87}
*/
package kata

func TwoOldestAges(ages []int) [2]int {
  var maxAge1, maxAge2 int
  for _, age := range ages {
    if age > maxAge1 {
      maxAge2 = maxAge1
      maxAge1 = age
    } else if age > maxAge2 {
      maxAge2 = age
    }
  }
  return [2]int{maxAge2, maxAge1}
}