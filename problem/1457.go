package problem

var TestNodeii = &TreeNode{
	Val: 2,
	Left: &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 1,
		},
	},
	Right: &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 1,
		},
	},
}

func PseudoPalindromicPaths(root *TreeNode) int {
	ans := 0
	var dfs func(node *TreeNode, num [10]int, level int)
	dfs = func(node *TreeNode, num [10]int, level int) {
		if node == nil {
			return
		}
		num[node.Val]++
		if node.Left == nil && node.Right == nil {
			countSingle := 0
			if level%2 != 0 {
				countSingle = -1
			}
			for _, v := range num {
				if v%2 != 0 {
					countSingle++
				}
			}
			if countSingle == 0 {
				ans++
			}
		}
		dfs(node.Left, num, level+1)
		dfs(node.Right, num, level+1)
	}
	dfs(root, [10]int{}, 1)
	return ans
}
