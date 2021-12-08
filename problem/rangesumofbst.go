package problem

func rangeSumBST(root *TreeNode, low int, high int) int {
	nums := []int{}
	var bst func(node *TreeNode)
	bst = func(node *TreeNode) {
		if node == nil {
			return
		}
		bst(node.Left)
		bst(node.Right)
		nums = append(nums, node.Val)
	}
	bst(root)

	sum, add := 0, false
	for i := 0; i < len(nums); i++ {
		if nums[i] == low {
			add = true
		}
		if !add {
			continue
		}
		sum += nums[i]
		if nums[i] == high {
			break
		}
	}
	return sum
}
