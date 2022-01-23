package problem

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, 2)
	dp[0], dp[1] = 1, 2
	for i := 2; i < n; i++ {
		dp[1], dp[0] = dp[1]+dp[0], dp[1]
	}
	return dp[1]
}
