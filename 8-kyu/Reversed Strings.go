/*
Complete the solution so that it reverses the string passed into it. 

```
'world'  =>  'dlrow'
```
*/
package kata

func Solution(word string) (out string) {
  for _, s := range word {
    out = string(s) + out
  }
  return
}