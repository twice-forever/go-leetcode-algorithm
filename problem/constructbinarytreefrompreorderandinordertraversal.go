package problem

func BuildTree(preorder []int, inorder []int) *TreeNode {
	treeNode := &TreeNode{}
	rebuilder(preorder, inorder, treeNode)
	return treeNode
}

func rebuilder(preorder []int, inorder []int, node *TreeNode) {
	if len(preorder) == 0 {
		node = nil
		return
	}

	node.Val = preorder[0]
	for i, v := range inorder {
		if v == node.Val {
			leftInorder := inorder[:i]
			rightInorder := inorder[i+1:]
			leftPreorder := preorder[1 : len(leftInorder)+1]
			rightPreorder := preorder[len(leftInorder)+1:]
			if len(leftPreorder) > 0 {
				node.Left = &TreeNode{}
				rebuilder(leftPreorder, leftInorder, node.Left)
			}
			if len(rightPreorder) > 0 {
				node.Right = &TreeNode{}
				rebuilder(rightPreorder, rightInorder, node.Right)
			}
		}
	}
}
