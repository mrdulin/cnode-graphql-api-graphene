/*
Your task, is to create NxN multiplication table, of size provided in parameter.

for example, when given size is 3:

1 2 3
2 4 6
3 6 9
for given example, the return value should be: [[1,2,3],[2,4,6],[3,6,9]]
*/
package multiplicationtable

func MultiplicationTable(size int) [][]int {
  res := [][]int{}
  for i := 1; i <= size; i ++ {
    v := []int{}
    for j := 1; j <= size; j++ {
      v = append(v, j*i)
    }
    res = append(res, v)
  }
  return res
}


