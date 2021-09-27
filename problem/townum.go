package problem

func TwoSum(nums []int, target int) []int {
	hashMap := make(map[int]int, len(nums))
	for i, v := range nums {
		if p, ok := hashMap[target-v]; ok {
			return []int{p, i}
		}
		hashMap[v] = i
	}
	return nil
}
