package problem

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	leftEqual := isSubtree(root.Left, subRoot)
	rightEqual := isSubtree(root.Right, subRoot)
	if leftEqual || rightEqual {
		return true
	}
	if root.Val == subRoot.Val {
		return equalVal(root, subRoot)
	}
	return false
}

func equalVal(root, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if subRoot == nil || root == nil {
		return false
	}
	if root.Val != subRoot.Val {
		return false
	}
	leftEqual := equalVal(root.Left, subRoot.Left)
	rightEqual := equalVal(root.Right, subRoot.Right)
	return leftEqual && rightEqual
}
