package problem

func robii(nums []int) int {
	l, ans := len(nums), nums[0]
	if l >= 2 {
		nums[1] = max(ans, nums[1])
		ans = nums[1]
	}
	for i := 2; i < l; i++ {
		nums[i] = max(nums[i-1], nums[i]+nums[i-2])
		if ans < nums[i] {
			ans = nums[i]
		}
	}
	return ans
}
