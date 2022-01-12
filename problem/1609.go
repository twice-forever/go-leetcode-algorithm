package problem

func isEvenOddTree(root *TreeNode) bool {
	nodes := []*TreeNode{root}
	depth := 0
	for len(nodes) > 0 {
		tmpNodes := []*TreeNode{}
		for k, v := range nodes {
			if depth%2 == 0 {
				if v.Val%2 == 0 || (k+1 < len(nodes) && v.Val >= nodes[k+1].Val) {
					return false
				}
			} else {
				if v.Val%2 != 0 || (k+1 < len(nodes) && v.Val <= nodes[k+1].Val) {
					return false
				}
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
	return true
}
