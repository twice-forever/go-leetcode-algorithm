package problem

import (
	"strconv"
)

func printTree(root *TreeNode) [][]string {
	height := getHeight(root)
	line := (1 << height) - 1
	result := make([][]string, height)
	for i := range result {
		result[i] = make([]string, line)
	}
	fill(&result, root, 0, 0, line-1)
	return result
}

func fill(res *([][]string), root *TreeNode, i, l, r int) {
	if root == nil {
		return
	}
	mid := (l + r) / 2
	(*res)[i][mid] = strconv.Itoa(root.Val)
	fill(res, root.Left, i+1, l, mid)
	fill(res, root.Right, i+1, mid+1, r)
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(getHeight(root.Left), getHeight(root.Right))
}
