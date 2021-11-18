package problem

import (
	"fmt"
	"strconv"
)

func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}
	leftStr := tree2str(root.Left)
	rightStr := tree2str(root.Right)
	if leftStr == "" && rightStr == "" {
		return strconv.Itoa(root.Val)
	}
	if rightStr == "" {
		return fmt.Sprintf("%d(%s)", root.Val, leftStr)
	}
	return fmt.Sprintf("%d(%s)(%s)", root.Val, leftStr, rightStr)
}
