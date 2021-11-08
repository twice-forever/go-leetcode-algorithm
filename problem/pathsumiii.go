package problem

func pathSumiii(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	num := pathSumiii(root.Left, targetSum)
	num += pathSumiii(root.Right, targetSum)
	num += rootSum(root, targetSum)
	return num
}

func rootSum(node *TreeNode, targetSum int) int {
	if node == nil {
		return 0
	}
	sum := rootSum(node.Left, targetSum-node.Val)
	sum += rootSum(node.Right, targetSum-node.Val)
	if targetSum == node.Val {
		return sum + 1
	}
	return sum
}
