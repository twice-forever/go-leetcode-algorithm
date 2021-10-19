package problem

func BuildTree2(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	node := &TreeNode{Val: postorder[len(postorder)-1]}
	for i, v := range inorder {
		if v == node.Val {
			node.Left = BuildTree2(inorder[:i], postorder[:len(inorder[:i])])
			node.Right = BuildTree2(inorder[i+1:], postorder[len(inorder[:i]):len(postorder)-1])
			break
		}
	}
	return node
}
