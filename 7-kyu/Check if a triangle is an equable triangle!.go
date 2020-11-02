/*
A triangle is called an equable triangle if its area equals its perimeter. Return true, if it is an equable triangle, else return false. You will be provided with the length of sides of the triangle. Happy Coding!
*/
package kata

import "math"

func EquableTriangle(a, b, c int) bool {
  p := a + b + c
  p2 := p / 2
  area := int(math.Sqrt(float64(p2*(p2-a)*(p2-b)*(p2-c))))
  return area == p
}