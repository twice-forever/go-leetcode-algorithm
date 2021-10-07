package problem

func MaxArea(height []int) int {
	lenH := len(height)
	if len(height) == 0 {
		return 0
	}

	maxArea := 0
	head := 0
	end := lenH - 1

	for {
		area := (end - head) * min(height[head], height[end])
		maxArea = max(maxArea, area)
		if height[head] > height[end] {
			end--
		} else {
			head++
		}
		if end == head {
			break
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
