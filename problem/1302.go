package problem

func deepestLeavesSum(root *TreeNode) int {
	depth, sum := 0, 0
	var dfs func(node *TreeNode, deep int)
	dfs = func(node *TreeNode, deep int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			if depth < deep {
				depth = deep
				sum = node.Val
			} else if depth == deep {
				sum += node.Val
			}
		}
		dfs(node.Left, deep+1)
		dfs(node.Right, deep+1)
	}
	dfs(root, 0)
	return sum
}
