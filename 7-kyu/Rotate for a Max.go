/*
Let us begin with an example:

Take a number: 56789. Rotate left, you get 67895.

Keep the first digit in place and rotate left the other digits: 68957.

Keep the first two digits in place and rotate the other ones: 68579.

Keep the first three digits and rotate left the rest: 68597. Now it is over since keeping the first four it remains only one digit which rotated is itself.

You have the following sequence of numbers:

56789 -> 67895 -> 68957 -> 68579 -> 68597

and you must return the greatest: 68957.

Task
Write function max_rot(n) which given a positive integer n returns the maximum number you got doing rotations similar to the above example.

So max_rot (or maxRot or ... depending on the language) is such as:

max_rot(56789) should return 68957

max_rot(38458215) should return 85821534
*/
package kata

import (
  "strconv"
  "sort"
)

func MaxRot(n int64) int64 {
  s := strconv.Itoa(int(n))
  arr := []int64{n}
  for i := 0; i < len(s)-1; i++ {
    m := s[:i] + Swap(s[i:])
    if ms, err := strconv.Atoi(m); err == nil {
      arr = append(arr, int64(ms))
    }
    s = m
  }
  sort.Slice(arr, func(i, j int) bool {
    return arr[i] < arr[j]
  }) 
  return arr[len(arr)-1]
}

func Swap(s string) string {
  if len(s) == 1 {
    return s
  }
  return s[1:] + string(s[0])
}