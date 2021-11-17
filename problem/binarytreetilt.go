package problem

func findTilt(root *TreeNode) int {
	tilt := 0
	var countTilt func(node *TreeNode) int
	countTilt = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := countTilt(node.Left)
		right := countTilt(node.Right)
		tilt += abs(left - right)
		return node.Val + left + right
	}
	countTilt(root)
	return tilt
}
