package problem

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if root == nil {
		return nil
	}
	if depth == 1 {
		left := root
		root = &TreeNode{Val: val, Left: left}
		return root
	}
	if depth == 2 {
		left := root.Left
		root.Left = &TreeNode{Val: val, Left: left}
		right := root.Right
		root.Right = &TreeNode{Val: val, Right: right}
		return root
	}
	root.Left = addOneRow(root.Left, val, depth-1)
	root.Right = addOneRow(root.Right, val, depth-1)
	return root
}
