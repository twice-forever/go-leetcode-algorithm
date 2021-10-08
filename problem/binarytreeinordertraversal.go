package problem

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	slice := make([]int, 0, 10)
	if root.Left != nil {
		slice = append(slice, InOrderTraversal(root.Left)...)
	}

	slice = append(slice, root.Val)

	if root.Right != nil {
		slice = append(slice, InOrderTraversal(root.Right)...)
	}
	return slice
}
