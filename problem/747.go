package problem

func dominantIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	num1, num2 := 0, -1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[num1] {
			num1, num2 = i, num1
		} else if num2 == -1 || nums[i] > nums[num2] {
			num2 = i
		}
	}
	if nums[num1]/2 >= nums[num2] {
		return num1
	}
	return -1
}
