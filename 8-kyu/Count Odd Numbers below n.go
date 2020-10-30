/*
Given a number **n**, return the number of positive odd numbers below **n**, EASY!
```go
OddCount(7) // => 3, i.e [1, 3, 5]
OddCount(15) // => 7, i.e [1, 3, 5, 7, 9, 11, 13]
```
Expect large Inputs!
*/
package kata

func OddCount(n int) (count int) {
  var r = n % 2
  return (n - r) / 2
}