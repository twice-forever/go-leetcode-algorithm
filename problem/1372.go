package problem

func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}
	numLeft := dfsii(root, 0) - 1
	numRight := dfsii(root, 1) - 1
	ans := max(numLeft, numRight)
	left := longestZigZag(root.Left)
	right := longestZigZag(root.Right)
	ansii := max(left, right)
	return max(ans, ansii)
}

func dfsii(node *TreeNode, direction int) int {
	if node == nil {
		return 0
	}
	num := 0
	if direction == 0 {
		num = dfsii(node.Left, 1)
	} else {
		num = dfsii(node.Right, 0)
	}
	return num + 1
}
