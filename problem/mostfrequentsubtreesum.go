package problem

func findFrequentTreeSum(root *TreeNode) []int {
	numsMap := map[int]int{}
	maxNum := 0

	var countSum func(node *TreeNode) int
	countSum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sum := countSum(node.Left)
		sum += countSum(node.Right)
		sum += node.Val
		numsMap[sum]++
		if numsMap[sum] > maxNum {
			maxNum = numsMap[sum]
		}
		return sum
	}
	countSum(root)

	nums := []int{}
	for k, v := range numsMap {
		if v == maxNum {
			nums = append(nums, k)
		}
	}
	return nums
}
