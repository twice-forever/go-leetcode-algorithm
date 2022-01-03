package problem

func maxProduct(root *TreeNode) int {
	sum := 0
	nums := []int{}
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		cur := left + right + node.Val
		nums = append(nums, cur)
		return cur
	}
	sum = dfs(root)

	maxNum := 0
	for _, v := range nums {
		if v*(sum-v) > maxNum {
			maxNum = v * (sum - v)
		}
	}
	return maxNum % (1e9 + 7)
}
