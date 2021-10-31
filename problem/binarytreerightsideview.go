package problem

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var nums []int
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		var tmp []*TreeNode
		nums = append(nums, nodes[len(nodes)-1].Val)
		for _, i := range nodes {
			if i.Left != nil {
				tmp = append(tmp, i.Left)
			}
			if i.Right != nil {
				tmp = append(tmp, i.Right)
			}
		}
		nodes = tmp
	}
	return nums
}
