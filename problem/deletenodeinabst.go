package problem

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		return RebuildTree(root)
	}
	root.Left = deleteNode(root.Left, key)
	root.Right = deleteNode(root.Right, key)
	return root
}

func RebuildTree(node *TreeNode) *TreeNode {
	if node.Left == nil {
		return node.Right
	}
	if node.Right == nil {
		return node.Left
	}

	left := node.Left
	node = node.Right
	leaf := node.Left
	if leaf == nil {
		node.Left = left
		return node
	}
	for {
		if leaf.Left == nil {
			leaf.Left = left
			break
		}
		leaf = leaf.Left
	}
	return node
}
