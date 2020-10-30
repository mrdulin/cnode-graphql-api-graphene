package kata

import (
  "strings"
  "strconv"
)

func Points(games []string) int {
  points := 0
  for _, game := range games {
    t := strings.Split(game, ":")
    x, _ := strconv.Atoi(t[0])
    y, _ := strconv.Atoi(t[1])
    if x > y {
      points += 3
    } else if x == y {
      points += 1
    } 
  }
  return points
}