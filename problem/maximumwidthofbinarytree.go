package problem

type LevelNode struct {
	TreeNode *TreeNode
	Pos      int
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	nodes := []LevelNode{LevelNode{root, 0}}
	maxWidth := 1
	for len(nodes) > 0 {
		tmpNodes := []LevelNode{}
		for _, v := range nodes {
			maxWidth = max(maxWidth, v.Pos-nodes[0].Pos+1)
			if v.TreeNode.Left != nil {
				tmpNodes = append(tmpNodes, LevelNode{v.TreeNode.Left, v.Pos * 2})
			}
			if v.TreeNode.Right != nil {
				tmpNodes = append(tmpNodes, LevelNode{v.TreeNode.Right, v.Pos*2 + 1})
			}
		}
		nodes = tmpNodes
	}
	return maxWidth
}
