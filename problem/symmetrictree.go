package problem

func IsSymmetric(root *TreeNode) bool {
	return findDiffNode(root, root)
}

func findDiffNode(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && findDiffNode(left.Left, right.Right) && findDiffNode(left.Right, right.Left)
}
