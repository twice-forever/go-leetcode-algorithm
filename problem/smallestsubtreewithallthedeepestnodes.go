package problem

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	parents := map[int]*TreeNode{}
	deepest, nodes := 0, []*TreeNode{}
	var findDeepest func(node *TreeNode, depth int)
	findDeepest = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			if depth == deepest {
				nodes = append(nodes, node)
			}
			if depth > deepest {
				deepest = depth
				nodes = []*TreeNode{node}
			}
		}
		if node.Left != nil {
			parents[node.Left.Val] = node
			findDeepest(node.Left, depth+1)
		}
		if node.Right != nil {
			parents[node.Right.Val] = node
			findDeepest(node.Right, depth+1)
		}
	}
	findDeepest(root, 0)

	if len(nodes) == 1 {
		return nodes[0]
	}
	for len(nodes) > 1 {
		tmpNodes := []*TreeNode{}
		for i, v := range nodes {
			if i > 0 {
				if tmpNodes[len(tmpNodes)-1] == parents[v.Val] {
					continue
				}
			}
			tmpNodes = append(tmpNodes, parents[v.Val])
		}
		nodes = tmpNodes
	}
	return nodes[0]
}
