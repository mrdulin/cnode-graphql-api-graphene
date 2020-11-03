/*
There is an array with some numbers. All numbers are equal except for one. Try to find it!

findUniq([ 1, 1, 1, 2, 1, 1 ]) === 2
findUniq([ 0, 0, 0.55, 0, 0 ]) === 0.55
It’s guaranteed that array contains at least 3 numbers.

The tests contain some very huge arrays, so think about performance.
*/
package kata

func FindUniq(arr []float32) float32 {
  var uniq float32
  hash := make(map[float32]int)
  for _, v := range arr {
    hash[v] += 1
  }
  for k, v := range hash {
    if v == 1 {
      uniq = k
      break
    }
  }
  return uniq
}