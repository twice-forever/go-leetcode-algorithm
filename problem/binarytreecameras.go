package problem

import "math"

const inf = math.MaxInt32 / 2

func minCameraCover(root *TreeNode) int {
	var dfs func(*TreeNode) (a, b, c int)
	dfs = func(node *TreeNode) (a int, b int, c int) {
		if node == nil {
			return inf, 0, 0
		}
		la, lb, lc := dfs(node.Left)
		ra, rb, rc := dfs(node.Right)
		a = lc + rc + 1
		b = min(a, min(la+rb, ra+lb))
		c = min(a, lb+rb)
		return
	}
	_, ans, _ := dfs(root)
	return ans
}
