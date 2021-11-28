package problem

type KthLargest struct {
	nums []int
	k    int
}

func Constructoriiiii(k int, nums []int) KthLargest {
	if len(nums) == 0 {
		return KthLargest{
			nums: nums,
			k:    k,
		}
	}
	sortNums := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		for ti, v := range sortNums {
			if nums[i] > v {
				sortNums = append(sortNums[:ti], append([]int{nums[i]}, sortNums[ti:]...)...)
				break
			}
			if ti == len(sortNums)-1 {
				sortNums = append(sortNums, nums[i])
			}
		}
	}
	return KthLargest{
		nums: sortNums,
		k:    k,
	}
}

func (this *KthLargest) Add(val int) int {
	if len(this.nums) == 0 {
		this.nums = append(this.nums, val)
	}
	for ti, v := range this.nums {
		if val > v {
			this.nums = append(this.nums[:ti], append([]int{val}, this.nums[ti:]...)...)
			break
		}
		if ti == len(this.nums)-1 {
			this.nums = append(this.nums, val)
		}
	}
	return this.nums[this.k-1]
}
