package problem

func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	ansMap := map[string]int{}
	for _, v := range generateParenthesis(n - 1) {
		for i := 0; i < len(v); i++ {
			ansMap[v[:i]+"()"+v[i:]] = 1
		}
	}
	ans := []string{}
	for k := range ansMap {
		ans = append(ans, k)
	}
	return ans
}
