package problem

func balanceBST(root *TreeNode) *TreeNode {
	nums := []int{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		nums = append(nums, node.Val)
		dfs(node.Right)
	}
	dfs(root)

	var build func(slice []int) *TreeNode
	build = func(slice []int) *TreeNode {
		l := len(slice)
		if l == 0 {
			return nil
		}
		node := &TreeNode{Val: slice[l/2]}
		node.Left = build(slice[:l/2])
		node.Right = build(slice[l/2+1:])
		return node
	}
	return build(nums)
}
