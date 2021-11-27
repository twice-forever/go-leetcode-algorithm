package problem

func longestUnivaluePath(root *TreeNode) int {
	num := 0
	if root == nil {
		return num
	}
	var getLongest func(node *TreeNode) int
	getLongest = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := getLongest(node.Left)
		right := getLongest(node.Right)
		if left != 0 && node.Left.Val != node.Val {
			left = 0
		}
		if right != 0 && node.Right.Val != node.Val {
			right = 0
		}
		num = max(num, 1+left+right)

		return max(1+left, 1+right)
	}
	getLongest(root)
	return num
}
