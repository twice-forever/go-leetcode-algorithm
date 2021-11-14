package problem

import "math"

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	maxNum := math.MinInt32
	nodes := []*TreeNode{root}
	nums := []int{}
	for len(nodes) > 0 {
		tmpNodes := []*TreeNode{}
		for _, v := range nodes {
			if v.Val > maxNum {
				maxNum = v.Val
			}
			if v.Left != nil {
				tmpNodes = append(tmpNodes, v.Left)
			}
			if v.Right != nil {
				tmpNodes = append(tmpNodes, v.Right)
			}
		}
		nums = append(nums, maxNum)
		maxNum = math.MinInt32
		nodes = tmpNodes
	}
	return nums
}
