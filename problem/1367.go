package problem

func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil || head == nil {
		return false
	}
	var left, right bool
	if head.Val == root.Val {
		if head.Next == nil {
			return true
		}
		if root.Left != nil {
			left = dfs(head.Next, root.Left)
		}
		if root.Right != nil {
			right = dfs(head.Next, root.Right)
		}
		return left || right
	}
	return isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func dfs(head *ListNode, root *TreeNode) bool {
	var left, right bool
	if head.Val == root.Val {
		if head.Next == nil {
			return true
		}
		if root.Left != nil {
			left = dfs(head.Next, root.Left)
		}
		if root.Right != nil {
			right = dfs(head.Next, root.Right)
		}
		return left || right
	}
	return false
}
