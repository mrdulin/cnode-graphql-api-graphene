/*
Take 2 strings s1 and s2 including only letters from ato z. Return a new sorted string, the longest possible, containing distinct letters,

each taken only once - coming from s1 or s2.
Examples:
a = "xyaabbbccccdefww"
b = "xxxxyyyyabklmopq"
longest(a, b) -> "abcdefklmopqwxy"

a = "abcdefghijklmnopqrstuvwxyz"
longest(a, a) -> "abcdefghijklmnopqrstuvwxyz"
*/
package kata

import (
  "strings"
  "sort"
)

func TwoToOne(s1 string, s2 string) string {
  strArr := strings.Split(s1 + s2, "")
  strs := removeDups(strArr)
  sort.Strings(strs)
  return strings.Join(strs, "")
}


func removeDups(strs []string) []string {
    result := make([]string, 0, len(strs))
    temp := map[string]struct{}{}
    for _, item := range strs {
        if _, ok := temp[item]; !ok {
            temp[item] = struct{}{}
            result = append(result, item)
        }
    }
    return result
}