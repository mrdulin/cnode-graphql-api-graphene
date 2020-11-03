/*
Write a simple parser that will parse and run Deadfish.

Deadfish has 4 commands, each 1 character long:

i increments the value (initially 0)
d decrements the value
s squares the value
o outputs the value into the return array
Invalid characters should be ignored.

Parse("iiisdoso") == []int{8, 64}
*/
package kata

func Parse(data string) []int{
  var n int = 0
  var r = []int{}
  for _, ch := range data {
    command := string(ch)
    switch command {
      case "i":
        n++
      case "d":
        n--
      case "s":
        n *= n
      case "o":
        r = append(r, n)
    }
  }
  return r
}