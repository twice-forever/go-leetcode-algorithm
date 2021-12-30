package problem

func deepestLeavesSum(root *TreeNode) int {
	nums := 0
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		nums = 0
		tmpNodes := []*TreeNode{}
		for _, v := range nodes {
			nums += v.Val
			if v.Left != nil {
				tmpNodes = append(tmpNodes, v.Left)
			}
			if v.Right != nil {
				tmpNodes = append(tmpNodes, v.Right)
			}
		}
		nodes = tmpNodes
	}
	return nums
}
