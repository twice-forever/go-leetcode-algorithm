package problem

var slices [][]int

func LevelOrder(root *TreeNode) [][]int {
	slices = make([][]int, 0)
	makeSlice(root, 0)
	return slices
}

func makeSlice(node *TreeNode, i int) {
	if node == nil {
		return
	}
	if len(slices) == i {
		slice := make([]int, 0)
		slices = append(slices, slice)
	}
	slices[i] = append(slices[i], node.Val)
	makeSlice(node.Left, i+1)
	makeSlice(node.Right, i+1)
}
