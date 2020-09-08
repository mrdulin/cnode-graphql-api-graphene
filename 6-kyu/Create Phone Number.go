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
