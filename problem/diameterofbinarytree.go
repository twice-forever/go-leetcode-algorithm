package problem

func diameterOfBinaryTree(root *TreeNode) int {
	maxPath := 0
	var dsf func(node *TreeNode) int
	dsf = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		nums := 0
		nums1 := dsf(node.Left)
		nums2 := dsf(node.Right)
		if nums1 > nums2 {
			nums = nums1 + 1
		} else {
			nums = nums2 + 1
		}
		if nums1+nums2 > maxPath {
			maxPath = nums1 + nums2
		}
		return nums
	}
	dsf(root)
	return maxPath
}
