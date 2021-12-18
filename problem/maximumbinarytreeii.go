package problem

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	var nums []int
	var maxNum = 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if maxNum < node.Val {
			maxNum = node.Val
		}
		nums = append(nums, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	nums = append(nums, val)
	return buildTree(nums)
}

func buildTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var maxI, maxV int
	for i, v := range nums {
		if maxV < v {
			maxI, maxV = i, v
		}
	}
	root := &TreeNode{Val: maxV}
	root.Left = buildTree(nums[:maxI])
	root.Right = buildTree(nums[maxI+1:])
	return root
}
