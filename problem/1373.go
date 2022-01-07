package problem

func maxSumBST(root *TreeNode) int {
	ans := 0
	var dfs func(node *TreeNode) (int, int, int, bool)
	dfs = func(node *TreeNode) (int, int, int, bool) {
		if node == nil {
			return 0, 0, 0, true
		}
		lMin, lMax, lSum, lIs := dfs(node.Left)
		rMin, rMax, rSum, rIs := dfs(node.Right)
		if !lIs || !rIs {
			return 0, 0, max(lSum, rSum), false
		}
		if lSum != 0 && lMax >= node.Val {
			return 0, 0, max(lSum, rSum), false
		}
		if rSum != 0 && rMin <= node.Val {
			return 0, 0, max(lSum, rSum), false
		}
		if lSum == 0 {
			lMin = node.Val
		}
		if rSum == 0 {
			rMax = node.Val
		}
		sum := lSum + rSum + node.Val
		if ans < sum {
			ans = sum
		}
		return lMin, rMax, sum, true
	}
	dfs(root)
	if ans < 0 {
		return 0
	}
	return ans
}
