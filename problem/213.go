package problem

func Robiii(nums []int) int {
	n := len(nums)
	if n > 1 {
		n1 := robii(nums[:n-1])
		n2 := robii(nums[1:n])
		return max(n1, n2)
	}
	return nums[0]
}
