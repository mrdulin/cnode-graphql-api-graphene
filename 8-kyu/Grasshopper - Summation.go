package kata

func Summation(n int) int {
  var sum int
  for i := 0; i <= n; i++ {
    sum += i
  }
  return sum
}