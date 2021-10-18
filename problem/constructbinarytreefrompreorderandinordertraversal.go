package problem

func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	treeNode := &TreeNode{
		Val: preorder[0],
	}

	for i, v := range inorder {
		if v == treeNode.Val {
			treeNode.Left = BuildTree(preorder[1:len(inorder[0:i])+1], inorder[:i])
			treeNode.Right = BuildTree(preorder[len(inorder[0:i])+1:], inorder[i+1:])
			break
		}
	}
	return treeNode
}
