package problem

import "fmt"

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	nodeMap := map[string]int{}
	nodes := []*TreeNode{}
	var makeStr func(node *TreeNode) string
	makeStr = func(node *TreeNode) string {
		if node == nil {
			return "#"
		}
		left := makeStr(node.Left)
		right := makeStr(node.Right)
		str := fmt.Sprintf("%d,%s,%s", node.Val, left, right)
		nodeMap[str]++
		if nodeMap[str] == 2 {
			nodes = append(nodes, node)
		}
		return str
	}
	makeStr(root)
	return nodes
}
