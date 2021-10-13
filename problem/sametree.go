package problem

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}

	left := IsSameTree(p.Left, q.Left)
	right := IsSameTree(p.Right, q.Right)

	if p.Val == q.Val && left && right {
		return true
	}
	return false
}
