package problem

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	nodes := []*TreeNode{root}
	sum := root.Val
	ans, level := 1, 1
	for len(nodes) > 0 {
		tmpNodes := []*TreeNode{}
		curSum := 0
		for _, v := range nodes {
			curSum += v.Val
			if v.Left != nil {
				tmpNodes = append(tmpNodes, v.Left)
			}
			if v.Right != nil {
				tmpNodes = append(tmpNodes, v.Right)
			}
		}
		if curSum > sum {
			sum = curSum
			ans = level
		}
		level++
		nodes = tmpNodes
	}
	return ans
}
