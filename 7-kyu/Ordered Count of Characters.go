/*
Count the number of occurrences of each character and return it as a list of tuples in order of appearance. For empty output return an empty list.

Example:

OrderedCount("abracadabra") == []Tuple{Tuple{'a', 5}, Tuple{'b', 2}, Tuple{'r', 2}, Tuple{'c', 1}, Tuple{'d', 1}}

// Where
type Tuple struct {
    Char  rune
    Count int
}

*/
package orderedcount

// Use the preloaded Tuple struct as return type
// type Tuple struct {
//  Char  rune
//  Count int
// }

func OrderedCount(text string) []Tuple {
  res := []Tuple{}
  for _, r := range text {
    i, t := FindByChar(res, r)
    if t == nil {
      res = append(res, Tuple{Char: r, Count: 1})
    } else {
      res[i].Count++
    }
  }
  return res
}

func FindByChar(tuples []Tuple, char rune) (int, *Tuple) {
  for i, t := range tuples {
    if t.Char == char {
      return i, &t
    }
  }
  return -1, nil
}