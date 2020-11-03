/*
Write a function that accepts an array of 10 integers (between 0 and 9), that returns a string of those numbers in the form of a phone number.

Example:
CreatePhoneNumber([10]uint{1,2,3,4,5,6,7,8,9,0})  // returns "(123) 456-7890"
The returned format must be correct in order to complete this challenge.
Don't forget the space after the closing parentheses!
*/
package __kyu

import (
	"fmt"
	"strconv"
	"strings"
)

func CreatePhoneNumber(numbers [10]uint) string {
	strs := []string{}
	for _, n := range numbers {
		strs = append(strs, strconv.FormatUint(uint64(n), 10))
	}
	a := strings.Join(strs[:3], "")
	b := strings.Join(strs[3:6], "")
	c := strings.Join(strs[6:], "")
	return fmt.Sprintf("(%s) %s-%s", a, b, c)
}
