package problem

func hasPathSum(root *TreeNode, targetSum int) bool {
	hasSum := false
	if root == nil {
		return false
	}
	if root.Left != nil {
		hasSum = hasPathSum(root.Left, targetSum-root.Val)
	}
	if hasSum {
		return hasSum
	}
	if root.Right != nil {
		hasSum = hasPathSum(root.Right, targetSum-root.Val)
	}
	if root.Right == nil && root.Left == nil {
		return root.Val == targetSum
	}
	return hasSum
}
