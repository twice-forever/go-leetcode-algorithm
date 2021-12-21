package problem

func recoverFromPreorder(traversal string) *TreeNode {
	path, pos := []*TreeNode{}, 0
	for pos < len(traversal) {
		level := 0
		for traversal[pos] == '-' {
			level++
			pos++
		}
		value := 0
		for ; pos < len(traversal) && traversal[pos] != '-'; pos++ {
			value = value*10 + int(traversal[pos]-'0')
		}
		node := &TreeNode{Val: value}
		if level == len(path) {
			if len(path) > 0 {
				path[len(path)-1].Left = node
			}
		} else {
			path = path[:level]
			path[len(path)-1].Right = node
		}
		path = append(path, node)
	}
	return path[0]
}
