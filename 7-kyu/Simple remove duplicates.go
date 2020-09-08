package __kyu

func Solve(arr []int) []int {
	var r []int
	for j := len(arr) - 1; j >= 0; j-- {
		rv := arr[j]
		for _, v := range arr {
			if rv == v && !Contains(r, v) {
				r = append([]int{v}, r...)
			}
		}
	}
	return r
}

func Contains(arr []int, v int) bool {
	for _, vv := range arr {
		if vv == v {
			return true
		}
	}
	return false
}
