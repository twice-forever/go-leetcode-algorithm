package problem

func construct2DArray(original []int, m int, n int) [][]int {
	if m*n != len(original) {
		return [][]int{}
	}
	ans := make([][]int, 0, m)
	for i := 0; i < len(original); i += n {
		ans = append(ans, original[i:i+n])
	}
	return ans
}
