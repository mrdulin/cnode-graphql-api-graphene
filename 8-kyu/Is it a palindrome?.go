/*
Write function isPalindrome that checks if a given string (case insensitive) is a palindrome.

```racket
In Racket, the function is called palindrome?

(palindrome? "nope") ; returns #f
(palindrome? "Yay")  ; returns #t
```
*/
package kata

import "strings"

func IsPalindrome(str string) bool {
  return reverse(str) == strings.ToLower(str)
}

func reverse(s string) (r string) {
  for _, v := range s {
    r = strings.ToLower(string(v)) + r
  }
  return
}