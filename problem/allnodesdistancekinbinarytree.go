package problem

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	parents := map[int]*TreeNode{}
	var getParents func(node *TreeNode)
	getParents = func(node *TreeNode) {
		if node.Left != nil {
			parents[node.Left.Val] = node
			getParents(node.Left)
		}
		if node.Right != nil {
			parents[node.Right.Val] = node
			getParents(node.Right)
		}
	}
	getParents(root)

	nums := []int{}
	var findTarget func(node, from *TreeNode, k int)
	findTarget = func(node, from *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == k {
			nums = append(nums, node.Val)
		}
		if node.Left != from {
			findTarget(node.Left, node, depth+1)
		}
		if node.Right != from {
			findTarget(node.Right, node, depth+1)
		}
		if parents[node.Val] != from {
			findTarget(parents[node.Val], node, depth+1)
		}
	}
	findTarget(target, target, 0)
	return nums
}
