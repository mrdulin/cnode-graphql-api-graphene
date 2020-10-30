/*
_Friday 13th or Black Friday is considered as unlucky day. Calculate how many unlucky days are in the given year._

Find the number of Friday 13th in the given year.

__Input:__ Year as an integer.

__Output:__ Number of Black Fridays in the year as an integer.

__Examples:__

	unluckyDays(2015) == 3
	unluckyDays(1986) == 1

***Note:*** In Ruby years will start from 1593.
*/
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
