package problem

type CBTInserter struct {
	root  *TreeNode
	nodes []*TreeNode
}

func Constructoriiiiii(root *TreeNode) CBTInserter {
	inserter := CBTInserter{root: root}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		if nodes[0].Left != nil && nodes[0].Right != nil {
			nodes = append(nodes, nodes[0].Left, nodes[0].Right)
			nodes = nodes[1:]
			continue
		}
		break
	}
	inserter.nodes = nodes
	return inserter
}

func (this *CBTInserter) Insert(val int) int {
	num := this.nodes[0].Val
	if this.nodes[0].Left == nil {
		this.nodes[0].Left = &TreeNode{Val: val}
		return num
	}
	this.nodes[0].Right = &TreeNode{Val: val}
	this.nodes = append(this.nodes, this.nodes[0].Left, this.nodes[0].Right)
	this.nodes = this.nodes[1:]
	return num
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}
