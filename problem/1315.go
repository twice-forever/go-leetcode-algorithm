package problem

func sumEvenGrandparent(root *TreeNode) int {
	ans := 0
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == 2 {
			ans += node.Val
			return
		}
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}

	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		tmpNodes := []*TreeNode{}
		for _, v := range nodes {
			if v.Val%2 == 0 {
				dfs(v, 0)
			}
			if v.Left != nil {
				tmpNodes = append(tmpNodes, v.Left)
			}
			if v.Right != nil {
				tmpNodes = append(tmpNodes, v.Right)
			}
		}
		nodes = tmpNodes
	}
	return ans
}
