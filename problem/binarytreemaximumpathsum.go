package problem

func maxPathSum(root *TreeNode) int {
	_, max := maxPath(root, root.Val)
	return max
}

func maxPath(root *TreeNode, maxNum int) (int, int) {
	if root == nil {
		return 0, maxNum
	}
	maxLeft, maxNum := maxPath(root.Left, maxNum)
	maxRight, maxNum := maxPath(root.Right, maxNum)
	if maxLeft < 0 {
		maxLeft = 0
	}
	if maxRight < 0 {
		maxRight = 0
	}

	sum := maxLeft + maxRight + root.Val
	if sum > maxNum {
		maxNum = sum
	}

	return max(maxLeft+root.Val, maxRight+root.Val), maxNum
}
