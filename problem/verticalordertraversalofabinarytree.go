package problem

import (
	"math"
	"sort"
)

type data struct{ col, row, val int }

func verticalTraversal(root *TreeNode) [][]int {
	nodes := []data{}
	var dfs func(root *TreeNode, row, col int)
	dfs = func(root *TreeNode, row, col int) {
		if root == nil {
			return
		}
		dfs(root.Left, row+1, col-1)
		dfs(root.Right, row+1, col+1)
		nodes = append(nodes, data{col: col, row: row, val: root.Val})
	}
	dfs(root, 0, 0)

	sort.Slice(nodes, func(i, j int) bool {
		a, b := nodes[i], nodes[j]
		return a.col < b.col || a.col == b.col && (a.row < b.row || a.row == b.row && a.val < b.val)
	})

	lastCol := math.MinInt32
	ans := [][]int{}
	for _, v := range nodes {
		if v.col != lastCol {
			lastCol = v.col
			ans = append(ans, nil)
		}
		ans[len(ans)-1] = append(ans[len(ans)-1], v.val)
	}
	return ans
}
