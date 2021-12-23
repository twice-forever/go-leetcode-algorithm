package problem

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	if root == nil {
		return nil
	}
	var dfs func(*TreeNode, int) *TreeNode
	dfs = func(node *TreeNode, sum int) *TreeNode {
		if node == nil {
			return nil
		}
		if node.Left == nil && node.Right == nil {
			if sum+node.Val < limit {
				return nil
			}
			return node
		}
		if node.Left != nil {
			node.Left = dfs(node.Left, sum+node.Val)
		}
		if node.Right != nil {
			node.Right = dfs(node.Right, sum+node.Val)
		}
		if node.Left == nil && node.Right == nil {
			return nil
		}
		return node
	}
	return dfs(root, 0)
}
