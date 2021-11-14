package problem

import "math"

func getMinimumDifference(root *TreeNode) int {
	minVal := math.MaxInt64
	pre := -1
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 {
			minVal = min(minVal, node.Val-pre)
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return minVal
}
