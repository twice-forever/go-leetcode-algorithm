package problem

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	index := len(preorder)
	for i, v := range preorder {
		if v > preorder[0] {
			index = i
			break
		}
	}
	root := &TreeNode{Val: preorder[0]}
	root.Left = bstFromPreorder(preorder[1:index])
	root.Right = bstFromPreorder(preorder[index:])
	return root
}
