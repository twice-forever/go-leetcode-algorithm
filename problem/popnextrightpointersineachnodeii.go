package problem

func connectii(root *Node) *Node {
	if root == nil {
		return root
	}

	tmp := [][]*Node{}
	slice := []*Node{root}
	for len(slice) > 0 {
		tmp = append(tmp, slice)
		tmpSlice := []*Node{}
		for _, v := range slice {
			if v.Left != nil {
				tmpSlice = append(tmpSlice, v.Left)
			}
			if v.Right != nil {
				tmpSlice = append(tmpSlice, v.Right)
			}
		}
		slice = tmpSlice
	}

	for _, slice := range tmp {
		for i, v := range slice {
			if i+1 < len(slice) {
				v.Next = slice[i+1]
			}
		}
	}
	return root
}

func connectiii(root *Node) *Node {
	if root == nil {
		return root
	}

	q := []*Node{root}
	for len(q) > 0 {
		tmp := q
		q = nil
		for i, v := range tmp {
			if i+1 < len(tmp) {
				v.Next = tmp[i+1]
			}
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}
		}
	}
	return root
}
