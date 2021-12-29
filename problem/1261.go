package problem

type FindElements struct {
	nums map[int]bool
}

func Constructoriiii(root *TreeNode) FindElements {
	f := FindElements{
		nums: make(map[int]bool),
	}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		f.nums[node.Val] = true
		if node.Left != nil {
			node.Left.Val = 2*node.Val + 1
			dfs(node.Left)
		}
		if node.Right != nil {
			node.Right.Val = 2*node.Val + 2
			dfs(node.Right)
		}
	}
	root.Val = 0
	dfs(root)
	return f
}

func (this *FindElements) Find(target int) bool {
	_, ok := this.nums[target]
	return ok
}
