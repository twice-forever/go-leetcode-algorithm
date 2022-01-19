package problem

func canJump(nums []int) bool {
	end := 0
	for i := 0; i < len(nums); i++ {
		if i+nums[i] > end {
			end = i + nums[i]
		}
		if i == end && i < len(nums)-1 {
			return false
		}
	}
	return true
}
