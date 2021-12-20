package problem

func maxAncestorDiff(root *TreeNode) int {
	ans := 0
	var dfs func(root *TreeNode, max, min int)
	dfs = func(root *TreeNode, max, min int) {
		if root == nil {
			return
		}
		if ans < abs(max-root.Val) {
			ans = abs(max - root.Val)
		}
		if ans < abs(min-root.Val) {
			ans = abs(min - root.Val)
		}
		if min > root.Val {
			min = root.Val
		}
		if max < root.Val {
			max = root.Val
		}
		dfs(root.Left, max, min)
		dfs(root.Right, max, min)
	}
	dfs(root, root.Val, root.Val)
	return ans
}
