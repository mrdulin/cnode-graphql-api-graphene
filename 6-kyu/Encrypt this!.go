/*
Acknowledgments:
I thank yvonne-liu for the idea and for the example tests :)

Description:
Encrypt this!

You want to create secret messages which can be deciphered by the Decipher this! kata. Here are the conditions:

Your message is a string containing space separated words.
You need to encrypt each word in the message using the following rules:
The first letter needs to be converted to its ASCII code.
The second letter needs to be switched with the last letter
Keepin' it simple: There are no special characters in input.
Examples:
EncryptThis("Hello") == "72olle"
EncryptThis("good") == "103doo"
EncryptThis("hello world") == "104olle 119drlo"
*/
package kata

import (
  "strconv"
  "strings"
)

func EncryptThis(text string) string {
  if len(text) == 0 {
    return text
  }
  tarr := strings.Split(text, " ")
  rarr := []string{}
  for _, t := range tarr {
    rs := []rune(t)
    firstLetter := strconv.Itoa(int(rs[0]))
    var r string
    if len(t) > 2 {
      r = firstLetter + string(rs[len(rs)-1]) + string(rs[2:len(rs)-1]) + string(rs[1])  
    } else {
      r = firstLetter + string(rs[1:])
    }
    rarr = append(rarr, r)
  }
  return strings.Join(rarr, " ")
}