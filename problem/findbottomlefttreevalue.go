package problem

func findBottomLeftValue(root *TreeNode) int {
	deepLeft := 0
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		tmpNodes := []*TreeNode{}
		deepLeft = nodes[0].Val
		for _, v := range nodes {
			if v.Left != nil {
				tmpNodes = append(tmpNodes, v.Left)
			}
			if v.Right != nil {
				tmpNodes = append(tmpNodes, v.Right)
			}
		}
		nodes = tmpNodes
	}
	return deepLeft
}
