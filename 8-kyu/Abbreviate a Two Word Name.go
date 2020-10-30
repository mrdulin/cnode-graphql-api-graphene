package kata

import (
  "strings"
  "fmt"
)

func AbbrevName(name string) string{
  words := strings.Split(name, " ")
  return fmt.Sprintf("%s.%s", strings.ToUpper(words[0][0:1]), strings.ToUpper(words[1][0:1]))
}