package problem

func maxSubArray(nums []int) int {
	dp := []int{nums[0]}
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		dp = append(dp, max(nums[i], nums[i]+dp[i-1]))
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
