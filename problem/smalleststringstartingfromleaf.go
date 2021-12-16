package problem

func smallestFromLeaf(root *TreeNode) string {
	var ans string
	var dfs func(node *TreeNode, words string)
	dfs = func(node *TreeNode, words string) {
		if node == nil {
			return
		}
		words = string(rune('a'+node.Val)) + words
		if node.Left == nil && node.Right == nil {
			if ans == "" {
				ans = words
				return
			}
			for i := 0; i < len(ans); i++ {
				if i == len(words) || words[i] < ans[i] {
					ans = words
					break
				} else if ans[i] < words[i] {
					break
				}
			}
		}
		dfs(node.Left, words)
		dfs(node.Right, words)
	}
	dfs(root, "")
	return ans
}
