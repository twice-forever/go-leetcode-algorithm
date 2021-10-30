package problem

type BSTIterator struct {
	nums []int
}

func Constructor(root *TreeNode) BSTIterator {
	nums := build(root)
	return BSTIterator{
		nums: nums,
	}
}

func (this *BSTIterator) Next() int {
	num := this.nums[0]
	this.nums = this.nums[1:]
	return num
}

func (this *BSTIterator) HasNext() bool {
	if len(this.nums) > 0 {
		return true
	}
	return false
}

func build(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	nums := build(node.Left)
	nums = append(nums, node.Val)
	nums = append(nums, build(node.Right)...)
	return nums
}
