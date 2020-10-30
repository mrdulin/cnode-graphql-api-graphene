/*
In this Kata, you will remove the left-most duplicates from a list of integers and return the result.

// Remove the 3's at indices 0 and 3
// followed by removing a 4 at index 1
solve([3, 4, 4, 3, 6, 3]); // => [4, 6, 3]
More examples can be found in the test cases.
*/
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
