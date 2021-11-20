package problem

func averageOfLevels(root *TreeNode) []float64 {
	averages := []float64{}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		total := 0
		len := len(nodes)
		tmpNodes := []*TreeNode{}
		for _, v := range nodes {
			total += v.Val
			if v.Left != nil {
				tmpNodes = append(tmpNodes, v.Left)
			}
			if v.Right != nil {
				tmpNodes = append(tmpNodes, v.Right)
			}
		}
		average := float64(total) / float64(len)
		averages = append(averages, average)
		nodes = tmpNodes
	}
	return averages
}
