package problem

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	xLeft, xRight := 0, 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		if node.Val == x {
			xLeft = left
			xRight = right
		}
		return 1 + left + right
	}
	dfs(root)
	other := n - xLeft - xRight - 1
	return other > xLeft+xRight || xRight > other+xLeft || xLeft > other+xRight
}
