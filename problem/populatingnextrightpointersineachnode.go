package problem

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	tmp1 := []*Node{root}
	doubleSlice := [][]*Node{}
	for len(tmp1) > 0 {
		doubleSlice = append(doubleSlice, tmp1)
		tmp2 := []*Node{}
		for _, value := range tmp1 {
			if value.Left != nil && value.Right != nil {
				tmp2 = append(tmp2, value.Left)
				tmp2 = append(tmp2, value.Right)
			}
		}
		tmp1 = tmp2
	}

	for _, slice := range doubleSlice {
		for i, v := range slice {
			if i+1 < len(slice) {
				v.Next = slice[i+1]
			}
		}
	}
	return root
}
