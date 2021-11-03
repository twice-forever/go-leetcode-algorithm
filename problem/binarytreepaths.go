package problem

import "fmt"

func binaryTreePaths(root *TreeNode) []string {
	var strSlice []string
	if root.Left != nil {
		strSlice = binaryTreePaths(root.Left)
	}
	if root.Right != nil {
		strSlice = append(strSlice, binaryTreePaths(root.Right)...)
	}
	for i := 0; i < len(strSlice); i++ {
		strSlice[i] = fmt.Sprintf("%d->%s", root.Val, strSlice[i])
	}
	if root.Right == nil && root.Left == nil {
		strSlice = append(strSlice, fmt.Sprintf("%d", root.Val))
	}

	return strSlice
}
