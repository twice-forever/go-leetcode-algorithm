package problem

import "math"

func distributeCoins(root *TreeNode) int {
	ans := 0
	executeii(root, &ans)
	return ans
}

func executeii(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	left := executeii(root.Left, ans)
	right := executeii(root.Right, ans)

	*ans += int(math.Abs(float64(left))) + int(math.Abs(float64(right)))
	return root.Val + left + right - 1
}
