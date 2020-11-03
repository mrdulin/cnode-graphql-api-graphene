/*
Complete the solution so that it splits the string into pairs of two characters. If the string contains an odd number of characters then it should replace the missing second character of the final pair with an underscore ('_').

Examples:

Solution("abc") //should return ["ab", "c_"]
Solution("abcdef") //should return ["ab", "cd", "ef"]
*/
package kata

func Solution(str string) []string {
  r := []string{}
  if len(str) % 2 != 0 {
    str = str + "_"
  }
  for i := 0; i < len(str); i += 2 {
    r = append(r, string([]rune(str)[i:i+2]))
  }
  return r
}