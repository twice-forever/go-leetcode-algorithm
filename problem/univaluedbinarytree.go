package problem

func isUnivalTree(root *TreeNode) bool {
	num := root.Val
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if node.Val != num {
			return false
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		return left && right
	}
	return dfs(root)
}
