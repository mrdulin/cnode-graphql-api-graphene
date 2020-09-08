package __kyu

import (
	"time"
)

func UnluckyDays(year int) int {
	var r int
	for i := 1; i <= 12; i++ {
		d := time.Date(year, time.Month(i), 13, 0, 0, 0, 0, time.Local)
		if d.Weekday() == time.Friday {
			r++
		}
	}
	return r
}
