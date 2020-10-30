/*
In this Kata, you will be given a string and two indexes (a and b). Your task is to reverse the portion of that string between those two indices inclusive.

solve("codewars",1,5) = "cawedors" -- elements at index 1 to 5 inclusive are "odewa". So we reverse them.
solve("cODEWArs", 1,5) = "cAWEDOrs" -- to help visualize.
Input will be lowercase and uppercase letters only.

The first index a will always be lower that than the string length; the second index b can be greater than the string length. More examples in the test cases. Good luck!
*/
package kata 

// import "fmt"

func Solve(s string, a, b int) string {
  if b > len(s) {
    b = len(s) - 1
  }
  return s[:a] + Reverse(s[a:b+1]) + s[b+1:]
}

func Reverse(s string) (r string) {
  for _, c := range s {
    r = string(c) + r
  } 
  return r
}