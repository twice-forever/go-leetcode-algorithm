package problem

func Trap(height []int) int {
	target, subHeight := []int{}, []int{}
	hasTwo := false
	for _, v := range height {
		if v > 1 {
			hasTwo = true
		}
		if v >= 1 {
			target = append(target, 1)
			subHeight = append(subHeight, v-1)
		} else {
			target = append(target, v)
			subHeight = append(subHeight, v)
		}
	}
	ans, curNum := 0, 0
	hasLeft := false
	for _, v := range target {
		if v != 0 && !hasLeft {
			hasLeft = true
		} else if v != 0 && hasLeft {
			ans += curNum
			curNum = 0
		} else if v == 0 && hasLeft {
			curNum++
		}
	}
	if hasTwo {
		num := Trap(subHeight)
		ans += num
	}
	return ans
}
