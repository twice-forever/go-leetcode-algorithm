package problem

func GenerateTrees(n int) []*TreeNode {
	return help(1, n)
}

func help(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	nodeSlice := make([]*TreeNode, 0, end-start+1)
	for i := start; i <= end; i++ {
		left := help(start, i-1)
		right := help(i+1, end)

		for _, lv := range left {
			for _, rv := range right {
				treeNode := TreeNode{Val: i}
				treeNode.Left = lv
				treeNode.Right = rv
				nodeSlice = append(nodeSlice, &treeNode)
			}
		}
	}
	return nodeSlice
}
