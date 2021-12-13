package problem

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	if root == nil {
		return []int{-1}
	}
	var ans []int
	index := 0
	flag := execute(root, voyage, &index, &ans)
	if !flag || index < len(voyage)-1 {
		return []int{-1}
	}
	return ans
}

func execute(root *TreeNode, voyage []int, index *int, ans *[]int) bool {
	if root == nil {
		*index -= 1
		return true
	}
	if *index >= len(voyage) {
		return false
	}
	if root.Val != voyage[*index] {
		*index -= 1
		return false
	}

	*index += 1
	left := execute(root.Left, voyage, index, ans)
	if !left {
		flip(root)
		*index += 1
		left := execute(root.Left, voyage, index, ans)
		if !left {
			return false
		}
		*ans = append(*ans, root.Val)
	}
	*index += 1
	right := execute(root.Right, voyage, index, ans)
	if !right {
		return false
	}
	return true
}

func flip(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
}
