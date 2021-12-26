package problem

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	ans := root
	maxDeep := 0
	var dfs func(node *TreeNode, depth int) int
	dfs = func(node *TreeNode, depth int) int {
		if node == nil {
			return depth - 1
		}
		if node.Left == nil && node.Right == nil {
			if maxDeep < depth {
				maxDeep = depth
				ans = node
			}
			return depth
		}
		left := dfs(node.Left, depth+1)
		right := dfs(node.Right, depth+1)
		if left == maxDeep && right == maxDeep {
			ans = node
		}
		return max(left, right)
	}
	dfs(root, 0)
	return ans
}
