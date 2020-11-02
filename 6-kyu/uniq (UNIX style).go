/*
Implement a function which behaves like the uniq command in UNIX.

It takes as input a sequence and returns a sequence in which all duplicate elements following each other have been reduced to one instance.

Example:

["a", "a", "b", "b", "c", "a", "b", "c"]  =>  ["a", "b", "c", "a", "b", "c"]
*/
package uniq

func Uniq(a []string) []string {
  res := []string{}
  for i, v := range a {
    if i == 0 || v != string(a[i-1]) {
      res = append(res, v)
    }
  }
  return res
}