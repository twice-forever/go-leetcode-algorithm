package problem

func preorderTraversalii(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var slice []int
	slice = append(slice, root.Val)
	leftSlice := preorderTraversalii(root.Left)
	slice = append(slice, leftSlice...)
	rightSlice := preorderTraversalii(root.Right)
	slice = append(slice, rightSlice...)
	return slice
}
