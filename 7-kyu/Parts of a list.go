/*
Write a function partlist that gives all the ways to divide a list (an array) of at least two elements into two non-empty parts.

Each two non empty parts will be in a pair (or an array for languages without tuples or a structin C - C: see Examples test Cases - )
Each part will be in a string
Elements of a pair must be in the same order as in the original array.
Examples of returns in different languages:
a = ["az", "toto", "picaro", "zone", "kiwi"] -->

[["az", "toto picaro zone kiwi"], ["az toto", "picaro zone kiwi"], ["az toto picaro", "zone kiwi"], ["az toto picaro zone", "kiwi"]]

or

a = {"az", "toto", "picaro", "zone", "kiwi"} -->

{{"az", "toto picaro zone kiwi"}, {"az toto", "picaro zone kiwi"}, {"az toto picaro", "zone kiwi"}, {"az toto picaro zone", "kiwi"}}

or

a = ["az", "toto", "picaro", "zone", "kiwi"] -->

[("az", "toto picaro zone kiwi"), ("az toto", "picaro zone kiwi"), ("az toto picaro", "zone kiwi"), ("az toto picaro zone", "kiwi")]

or

a = [|"az", "toto", "picaro", "zone", "kiwi"|] -->

[("az", "toto picaro zone kiwi"), ("az toto", "picaro zone kiwi"), ("az toto picaro", "zone kiwi"), ("az toto picaro zone", "kiwi")]

or

a = ["az", "toto", "picaro", "zone", "kiwi"] -->

"(az, toto picaro zone kiwi)(az toto, picaro zone kiwi)(az toto picaro, zone kiwi)(az toto picaro zone, kiwi)"

You can see other examples for each language in "Your test cases"
*/
package kata

import "strings"

func PartList(arr []string) string {
  r := []string{}
  for i := 0; i < len(arr)-1; i++ {
    v := append(arr[:i+1], arr[i+1:]...)
    var s string
    for j, vv := range v {
      if i == j {
        s += vv + ", "
      } else {
        s += vv + " "
      }
    }
    s = "(" + strings.TrimSpace(s) + ")"
    r = append(r, s)
  }
  return strings.Join(r, "")
}