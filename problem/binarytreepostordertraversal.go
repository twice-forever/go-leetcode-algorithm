package problem

func postorderTraversal(root *TreeNode) []int {
	var vals []int
	var postorder func(root *TreeNode)
	postorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		postorder(root.Left)
		postorder(root.Right)
		vals = append(vals, root.Val)
	}
	postorder(root)
	return vals
}
