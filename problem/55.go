package problem

func canJump(nums []int) bool {
	end := 0
	for i := 0; i < len(nums); i++ {
		if i <= end {
			if i+nums[i] > end {
				end = i + nums[i]
			}
			if end >= len(nums)-1 {
				return true
			}
		}
	}
	return false
}
