package problem

func Robiii(nums []int) int {
	n := len(nums)
	if n > 1 {
		return max(count(nums[:n-1]), count(nums[1:n]))
	}
	return nums[0]
}

func count(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}
	f, s := nums[0], max(nums[0], nums[1])
	for i := 2; i < l; i++ {
		f, s = s, max(f+nums[i], s)
	}
	return s
}
