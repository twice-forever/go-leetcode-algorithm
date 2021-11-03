package problem

import "fmt"

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	strSlice := binaryTreePaths(root.Left)
	strSlice = append(strSlice, binaryTreePaths(root.Right)...)
	for i := 0; i < len(strSlice); i++ {
		strSlice[i] = fmt.Sprintf("%d->%s", root.Val, strSlice[i])
	}
	if root.Right == nil && root.Left == nil {
		strSlice = append(strSlice, fmt.Sprintf("%d", root.Val))
	}
	return strSlice
}
