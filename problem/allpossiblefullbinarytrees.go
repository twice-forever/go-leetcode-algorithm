package problem

func allPossibleFBT(n int) []*TreeNode {
	if n == 1 {
		return []*TreeNode{{Val: 0}}
	}
	ans := []*TreeNode{}
	for i := 1; i < n; i += 2 {
		lefts := allPossibleFBT(i)
		rights := allPossibleFBT(n - i - 1)
		for _, lv := range lefts {
			for _, rv := range rights {
				n := &TreeNode{Val: 0}
				n.Left = lv
				n.Right = rv
				ans = append(ans, n)
			}
		}
	}
	return ans
}
