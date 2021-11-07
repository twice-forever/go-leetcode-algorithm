package problem

func rob(root *TreeNode) int {
	sums := sumGet(root)
	return max(sums[0], sums[1])
}

func sumGet(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}

	l, r := sumGet(node.Left), sumGet(node.Right)
	selected := node.Val + l[1] + r[1]
	notSelected := max(l[0], l[1]) + max(r[0], r[1])
	return []int{selected, notSelected}
}
