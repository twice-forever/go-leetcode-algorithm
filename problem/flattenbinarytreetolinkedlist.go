package problem

func flatten(root *TreeNode) {
	list := preorderTraversal(root)
	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left, prev.Right = nil, curr
	}
}

func preorderTraversal(node *TreeNode) []*TreeNode {
	list := []*TreeNode{}
	if node != nil {
		list = append(list, node)
		list = append(list, preorderTraversal(node.Left)...)
		list = append(list, preorderTraversal(node.Right)...)
	}
	return list
}
