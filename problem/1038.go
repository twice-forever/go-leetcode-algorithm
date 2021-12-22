package problem

func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(root *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		root.Right = dfs(root.Right)
		sum += root.Val
		root.Val = sum
		root.Left = dfs(root.Left)
		return root
	}
	return dfs(root)
}
