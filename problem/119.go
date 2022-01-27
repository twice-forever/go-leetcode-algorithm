package problem

func getRow(rowIndex int) []int {
	ans := []int{1}
	for i := 1; i <= rowIndex; i++ {
		t := make([]int, i+1)
		t[0], t[i] = 1, 1
		for j := 1; j < i; j++ {
			t[j] = ans[j-1] + ans[j]
		}
		ans = t
	}
	return ans
}
