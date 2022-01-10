package problem

import "math"

func goodNodes(root *TreeNode) int {
	ans := 0
	var dfs func(node *TreeNode, maxNum int)
	dfs = func(node *TreeNode, maxNum int) {
		if node == nil {
			return
		}
		if node.Val >= maxNum {
			maxNum = node.Val
			ans++
		}
		dfs(node.Left, maxNum)
		dfs(node.Right, maxNum)
	}
	dfs(root, math.MinInt32)
	return ans
}
