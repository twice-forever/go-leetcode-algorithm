package problem

func findTarget(root *TreeNode, k int) bool {
	findMap := map[int]int{}
	var find func(root *TreeNode, k int) bool
	find = func(root *TreeNode, k int) bool {
		if root == nil {
			return false
		}
		if _, ok := findMap[k-root.Val]; ok {
			return true
		}
		findMap[root.Val] = 1
		return find(root.Left, k) || find(root.Right, k)
	}
	return find(root, k)
}
