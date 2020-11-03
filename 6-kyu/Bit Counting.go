/*
Write a function that takes an integer as input, and returns the number of bits that are equal to one in the binary representation of that number. You can guarantee that input is non-negative.

Example: The binary representation of 1234 is 10011010010, so the function should return 5 in this case
*/
package kata

import (
  "strconv"
)

func CountBits(n uint) int {
  b := strconv.FormatInt(int64(n), 2)
  r := 0
  for _, v := range b {
    if v == '1' {
      r++
    }
  }
  return r
}