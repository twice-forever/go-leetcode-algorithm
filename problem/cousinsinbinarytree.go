package problem

func isCousins(root *TreeNode, x int, y int) bool {
	var deep int
	ans := true
	var dfs func(node *TreeNode, depth int) bool
	dfs = func(node *TreeNode, depth int) bool {
		if node == nil {
			return false
		}
		left := dfs(node.Left, depth+1)
		right := dfs(node.Right, depth+1)
		if left && right {
			ans = false
		}
		if node.Val == x || node.Val == y {
			if deep == 0 {
				deep = depth
			} else if deep != depth {
				ans = false
			}
			return true
		}
		return false
	}
	dfs(root, 0)
	return ans
}
