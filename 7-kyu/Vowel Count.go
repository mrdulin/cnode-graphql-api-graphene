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
