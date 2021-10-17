package problem

func maxDepth(root *TreeNode) int {
	return deepNum(root, 1)
}

func deepNum(node *TreeNode, deep int) int {
	if node == nil {
		return deep - 1
	}

	var num1, num2 int
	if node.Left != nil {
		num1 = deepNum(node.Left, deep+1)
	}
	if node.Right != nil {
		num2 = deepNum(node.Right, deep+1)
	}
	if num2 > num1 {
		return num2
	} else {
		return num1
	}
}
