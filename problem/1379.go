package problem

func getTargetCopy(original, cloned, target *TreeNode) *TreeNode {
	var ans *TreeNode
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Val == target.Val {
			ans = root
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(cloned)
	return ans
}
