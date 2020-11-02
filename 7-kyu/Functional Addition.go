/*
Create a function add(n)/Add(n) which returns a function that always adds n to any number

Note for Java: the return type and methods have not been provided to make it a bit more challenging.

var addOne = Add(1)
addOne(3) // 4
*/
package kata

func Add(x int) func(int) int {
  return func(y int) int {
    return x + y
  }
}