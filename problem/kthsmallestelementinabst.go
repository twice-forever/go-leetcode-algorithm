package problem

func kthSmallest(root *TreeNode, k int) int {
	var smallest func(node *TreeNode)
	var returnVal int
	smallest = func(node *TreeNode) {
		if node == nil {
			return
		}
		smallest(node.Left)
		k--
		if k == 0 {
			returnVal = node.Val
		}
		smallest(node.Right)
	}
	smallest(root)
	return returnVal
}
