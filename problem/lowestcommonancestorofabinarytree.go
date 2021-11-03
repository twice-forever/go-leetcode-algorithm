package problem

func lowestCommonAncestorii(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	node1 := lowestCommonAncestorii(root.Left, p, q)
	node2 := lowestCommonAncestorii(root.Right, p, q)
	if node1 == nil {
		return node2
	} else if node2 == nil {
		return node1
	}
	return root
}
