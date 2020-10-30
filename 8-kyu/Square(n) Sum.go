package kata

import "math"

func SquareSum(numbers []int) int {
  var r float64
	for _, n := range numbers {
		r += math.Pow(float64(n), 2)
	}
	return int(r)
}