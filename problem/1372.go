package problem

func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		_, left := dfs(node.Left)
		right, _ := dfs(node.Right)
		if right > ans {
			ans = right
		}
		if left > ans {
			ans = left
		}
		return left + 1, right + 1
	}
	dfs(root)
	return ans
}
