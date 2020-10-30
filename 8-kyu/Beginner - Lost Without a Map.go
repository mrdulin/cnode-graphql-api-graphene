/*
Given an array of integers, return a new array with each value doubled.

For example:

`[1, 2, 3] --> [2, 4, 6]`

For the beginner, try to use the `map` method - it comes in very handy quite a lot so is a good one to know.

~~~if:racket
```racket
;for racket you are given a list
(maps '(1 2 3)) ; returns '(2 4 6)
```
~~~
*/
package kata

func Maps(x []int) []int {
  var rval []int
  for _, v := range x {
    rval = append(rval, v * 2)
  }
  return rval
}