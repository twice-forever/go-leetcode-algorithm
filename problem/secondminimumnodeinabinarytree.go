package problem

func findSecondMinimumValue(root *TreeNode) int {
	secondMin := -1
	var findValue func(node *TreeNode)
	findValue = func(node *TreeNode) {
		if node == nil || secondMin != -1 && root.Val >= secondMin {
			return
		}
		if node.Val > root.Val {
			secondMin = node.Val
		}
		findValue(node.Left)
		findValue(node.Right)
	}
	findValue(root)
	return secondMin
}
