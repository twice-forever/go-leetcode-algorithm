package problem

func increasingBST(root *TreeNode) *TreeNode {
	vals := []int{}
	var bst func(node *TreeNode)
	bst = func(node *TreeNode) {
		if node == nil {
			return
		}
		bst(node.Left)
		vals = append(vals, node.Val)
		bst(node.Right)
	}
	bst(root)

	ans := &TreeNode{}
	curNode := ans
	for _, v := range vals {
		curNode.Right = &TreeNode{Val: v}
		curNode = curNode.Right
	}
	return ans.Right
}
