package problem

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	root1Nodes, root2Nodes := []int{}, []int{}
	var dfs func(node *TreeNode, nodes *[]int)
	dfs = func(node *TreeNode, nodes *[]int) {
		if node == nil {
			return
		}
		dfs(node.Left, nodes)
		*nodes = append(*nodes, node.Val)
		dfs(node.Right, nodes)
	}
	dfs(root1, &root1Nodes)
	dfs(root2, &root2Nodes)

	ans := []int{}
	i, z := 0, 0
	for {
		if i == len(root1Nodes) {
			ans = append(ans, root2Nodes[z:]...)
			break
		} else if z == len(root2Nodes) {
			ans = append(ans, root1Nodes[i:]...)
			break
		}
		if root1Nodes[i] < root2Nodes[z] {
			ans = append(ans, root1Nodes[i])
			i++
		} else if root1Nodes[i] > root2Nodes[z] {
			ans = append(ans, root2Nodes[z])
			z++
		} else if root1Nodes[i] == root2Nodes[z] {
			ans = append(ans, root1Nodes[i], root2Nodes[z])
			i++
			z++
		}
	}
	return ans
}
