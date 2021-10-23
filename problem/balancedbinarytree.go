package problem

func isBalanced(root *TreeNode) bool {
	_, isB := deepBool(root, 0)
	return isB
}

func deepBool(node *TreeNode, deep int) (int, bool) {
	if node == nil {
		return deep - 1, true
	}
	leftDeep, leftB := deepBool(node.Left, deep+1)
	rightDeep, rightB := deepBool(node.Right, deep+1)
	if !leftB || !rightB {
		return deep, false
	}
	if abs(leftDeep-rightDeep) > 1.0 {
		return deep, false
	}
	if leftDeep > rightDeep {
		return leftDeep, true
	}
	return rightDeep, true
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
