package problem

func zigzagLevelOrder(root *TreeNode) [][]int {
	a := [][]int{}
	if root == nil {
		return a
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		a = append(a, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			a[i] = append(a[i], q[j].Val)

			if q[j].Left != nil {
				p = append(p, q[j].Left)
			}
			if q[j].Right != nil {
				p = append(p, q[j].Right)
			}
		}
		q = p
	}
	for i := 1; i < len(a); i += 2 {
		tmp := []int{}
		for j := len(a[i]) - 1; j >= 0; j-- {
			tmp = append(tmp, a[i][j])
		}
		a[i] = tmp
	}
	return a
}
