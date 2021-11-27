package problem

import "math"

func findSecondMinimumValue(root *TreeNode) int {
	if root.Left == nil {
		return -1
	}
	secondMin := math.MaxInt32
	hasSecond := false
	var findValue func(node *TreeNode)
	findValue = func(node *TreeNode) {
		if node == nil {
			return
		}
		if secondMin >= node.Val && node.Val > root.Val {
			hasSecond = true
			secondMin = node.Val
		}
		findValue(node.Left)
		findValue(node.Right)
	}
	findValue(root)
	if !hasSecond {
		return -1
	}
	return secondMin
}
