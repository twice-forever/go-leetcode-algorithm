package problem

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	subTree := []*TreeNode{}
	var getSubTree func(root *TreeNode)
	getSubTree = func(root *TreeNode) {
		if root.Left != nil {
			subTree = append(subTree, root.Left)
			getSubTree(root.Left)
		}
		if root.Right != nil {
			subTree = append(subTree, root.Right)
			getSubTree(root.Right)
		}
	}
	getSubTree(root)

	sameSubTree := []*TreeNode{}
	for len(subTree) > 1 {
		str := subTree[0]
		subTree = subTree[1:]
		hasSame := false
		for _, v := range sameSubTree {
			if compareTree(str, v) {
				hasSame = true
				break
			}
		}
		if !hasSame {
			for j := 0; j < len(subTree); j++ {
				if compareTree(str, subTree[j]) {
					sameSubTree = append(sameSubTree, str)
					break
				}
			}
		}
	}
	return sameSubTree
}

func compareTree(src, target *TreeNode) bool {
	if src == nil && target == nil {
		return true
	}
	if src == nil || target == nil {
		return false
	}
	if src.Val != target.Val {
		return false
	}
	left := compareTree(src.Left, target.Left)
	right := compareTree(src.Right, target.Right)
	if left == true && right == true {
		return true
	}
	return false
}
