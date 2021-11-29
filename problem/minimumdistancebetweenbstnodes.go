package problem

import "math"

func minDiffInBST(root *TreeNode) int {
	minNum, pre := math.MaxInt32, -1
	var findMin func(node *TreeNode)
	findMin = func(node *TreeNode) {
		if node == nil {
			return
		}
		findMin(node.Left)
		if pre != -1 && minNum > node.Val-pre {
			minNum = node.Val - pre
		}
		pre = node.Val
		findMin(node.Right)
	}
	findMin(root)
	return minNum
}
