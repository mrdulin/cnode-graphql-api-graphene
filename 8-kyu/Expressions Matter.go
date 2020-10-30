package kata

func ExpressionMatter(a int, b int, c int) int {
   vals := []int{a+b+c, a*b*c, a*b+c, a*(b+c), a+b*c, (a+b)*c}
   max := vals[0]
   for _, v := range vals {
     if max < v {
       max = v
     }
   }
   return max
}