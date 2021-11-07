package problem

func sumOfLeftLeaves(root *TreeNode) int {
	return sumLeft(root, false)
}

func sumLeft(node *TreeNode, isLeft bool) int {
	if node == nil {
		return 0
	}

	sum := sumLeft(node.Left, true)
	sum += sumLeft(node.Right, false)
	if node.Left == nil && node.Right == nil && isLeft == true {
		sum += node.Val
	}
	return sum
}
