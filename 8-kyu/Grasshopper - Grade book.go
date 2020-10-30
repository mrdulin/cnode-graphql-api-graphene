/**
Grade book
Complete the function so that it finds the mean of the three scores passed to it and returns the letter value associated with that grade.

Numerical Score	Letter Grade
90 <= score <= 100	'A'
80 <= score < 90	'B'
70 <= score < 80	'C'
60 <= score < 70	'D'
0 <= score < 60	'F'
Tested values are all between 0 and 100. Theres is no need to check for negative values or values greater than 100.
*/
package kata

import (
  "unicode/utf8"
)

func GetGrade(a,b,c int) rune {
  avg := (a + b + c) / 3
  var grade string
  switch {
    case 90 <= avg && avg <= 100:
      grade = "A"
    case 80 <= avg && avg < 90:
      grade = "B"
    case 70 <= avg && avg < 80:
      grade = "C"
    case 60 <= avg && avg < 70:
      grade = "D"
    case 0 <= avg && avg < 60:
      grade = "F"
  }
  g, _ := utf8.DecodeRuneInString(grade)
  return g
}