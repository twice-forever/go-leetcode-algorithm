package problem

func containsNearbyDuplicate(nums []int, k int) bool {
	p := map[int]int{}
	for i, v := range nums {
		if index, ok := p[v]; ok && i-index <= k {
			return true
		}
		p[v] = i
	}
	return false
}
