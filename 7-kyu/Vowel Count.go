/*
Return the number (count) of vowels in the given string.

We will consider a, e, i, o, u as vowels for this Kata (but not y).

The input string will only consist of lower case letters and/or spaces.

*/
package __kyu

import "strings"

func GetCount(str string) (count int) {
	var vowels = "aeiou"
	for _, v := range str {
		if strings.Contains(vowels, string(v)) {
			count++
		}
	}
	return
}
