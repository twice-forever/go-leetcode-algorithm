package problem

import "math"

func constructMaximumBinaryTree(nums []int) *TreeNode {
	numMap := map[int]int{}
	maxNum := math.MinInt64
	for i, v := range nums {
		numMap[v] = i
		if maxNum < v {
			maxNum = v
		}
	}
	node := &TreeNode{Val: maxNum}
	if numMap[maxNum] != 0 {
		node.Left = constructMaximumBinaryTree(nums[:numMap[maxNum]])
	}
	if numMap[maxNum] < len(nums)-1 {
		node.Right = constructMaximumBinaryTree(nums[numMap[maxNum]+1:])
	}
	return node
}
