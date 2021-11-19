package problem

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	root := root1

	if root != nil && root2 != nil {
		root.Val += root2.Val
	} else if root == nil && root2 != nil {
		root = &TreeNode{Val: root2.Val}
	} else if root == nil && root2 == nil {
		return nil
	} else if root != nil && root2 == nil {
		return root
	}

	if root.Left == nil && root2.Left != nil {
		root.Left = &TreeNode{}
	}
	if root.Right == nil && root2.Right != nil {
		root.Right = &TreeNode{}
	}
	mergeTrees(root.Left, root2.Left)
	mergeTrees(root.Right, root2.Right)
	return root
}
