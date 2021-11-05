package problem

func isValidSerialization(preorder string) bool {
	stk := []int{1}
	for i := 0; i < len(preorder); {
		if len(stk) == 0 {
			return false
		}
		if preorder[i] == ',' {
			i++
		} else if preorder[i] == '#' {
			stk[len(stk)-1]--
			if stk[len(stk)-1] == 0 {
				stk = stk[:len(stk)-1]
			}
			i++
		} else {
			for i < len(preorder) && preorder[i] != ',' {
				i++
			}
			stk[len(stk)-1]--
			if stk[len(stk)-1] == 0 {
				stk = stk[:len(stk)-1]
			}
			stk = append(stk, 2)
		}
	}
	return len(stk) == 0
}
