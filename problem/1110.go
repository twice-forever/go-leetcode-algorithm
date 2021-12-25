package problem

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	ans := []*TreeNode{}
	var dfs func(node *TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		node.Left = dfs(node.Left)
		node.Right = dfs(node.Right)
		for i, v := range to_delete {
			if v == node.Val {
				if node.Left != nil {
					ans = append(ans, node.Left)
				}
				if node.Right != nil {
					ans = append(ans, node.Right)
				}
				to_delete = append(to_delete[:i], to_delete[i+1:]...)
				return nil
			}
		}
		return node
	}
	root = dfs(root)
	if root != nil {
		ans = append(ans, root)
	}
	return ans
}
