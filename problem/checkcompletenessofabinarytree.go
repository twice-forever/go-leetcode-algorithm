package problem

func isCompleteTree(root *TreeNode) bool {
	nodes := []*TreeNode{root}
	hasNil := false
	for len(nodes) > 0 {
		tmpNodes := []*TreeNode{}
		for _, v := range nodes {
			if v != nil && hasNil {
				return false
			}
			if v == nil {
				hasNil = true
			}
			if v != nil {
				tmpNodes = append(tmpNodes, v.Left)
				tmpNodes = append(tmpNodes, v.Right)
			}
		}
		nodes = tmpNodes
	}
	return true
}
