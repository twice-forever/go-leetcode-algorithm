package problem

func sumRootToLeaf(root *TreeNode) int {
	var sum int
	var dfs func(root *TreeNode, num int)
	dfs = func(root *TreeNode, num int) {
		if root == nil {
			return
		}
		num = num<<1 + root.Val
		dfs(root.Left, num)
		dfs(root.Right, num)
		if root.Left == nil && root.Right == nil {
			sum += num
			return
		}
	}
	dfs(root, 0)
	return sum
}
