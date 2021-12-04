package problem

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return root
	}
	leftTreeLen := 0
	for i, v := range postorder {
		if v == preorder[1] {
			leftTreeLen = i + 1
			break
		}
	}
	root.Left = constructFromPrePost(preorder[1:leftTreeLen+1], postorder[0:leftTreeLen])
	root.Right = constructFromPrePost(preorder[leftTreeLen+1:], postorder[leftTreeLen:len(postorder)-1])
	return root
}
