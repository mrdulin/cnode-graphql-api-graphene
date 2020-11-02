/*
Yor task is to write function factorial
*/
package kata

func Factorial(n int) int {
  if n == 0 || n == 1 {
    return 1
  }
  return Factorial(n - 1) * n
}