package problem

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	nodes := []*TreeNode{}
	var findLeaf func(node *TreeNode)
	findLeaf = func(node *TreeNode) {
		if node == nil {
			return
		}
		findLeaf(node.Left)
		findLeaf(node.Right)
		if node.Left == nil && node.Right == nil {
			nodes = append(nodes, node)
		}
	}
	findLeaf(root1)

	var findSame func(node *TreeNode) bool
	findSame = func(node *TreeNode) bool {
		if node.Left != nil {
			isSame := findSame(node.Left)
			if !isSame {
				return isSame
			}
		}
		if node.Right != nil {
			isSame := findSame(node.Right)
			if !isSame {
				return isSame
			}
		}
		if node.Left == nil && node.Right == nil {
			if len(nodes) == 0 {
				return false
			}
			if node.Val != nodes[0].Val {
				return false
			}
			if node.Val == nodes[0].Val {
				nodes = nodes[1:]
			}
		}
		return true
	}

	return findSame(root2) && len(nodes) == 0
}
