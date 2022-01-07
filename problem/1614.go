package problem

func maxDepth(s string) int {
	maxNum, curNum := 0, 0
	for _, v := range s {
		if v == '(' {
			curNum++
			if curNum > maxNum {
				maxNum = curNum
			}
		} else if v == ')' {
			curNum--
		}
	}
	return maxNum
}
