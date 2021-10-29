package problem

func sumNumbers(root *TreeNode) int {
	return pathNum(root, 0, 0)
}

func pathNum(node *TreeNode, num, sum int) int {
	if node == nil {
		return sum
	}

	num = num*10 + node.Val
	if node.Left == nil && node.Right == nil {
		sum += num
		return sum
	}

	sum = pathNum(node.Left, num, sum)
	sum = pathNum(node.Right, num, sum)
	return sum
}
