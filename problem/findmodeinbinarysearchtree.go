package problem

func findMode(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	numsMap := make(map[int]int)
	maxTime := 0

	var midSearch func(node *TreeNode)
	midSearch = func(node *TreeNode) {
		if node == nil {
			return
		}
		midSearch(node.Left)
		numsMap[node.Val]++
		if numsMap[node.Val] > maxTime {
			maxTime = numsMap[node.Val]
		}
		midSearch(node.Right)
	}
	midSearch(root)

	var nums []int
	for k, v := range numsMap {
		if v == maxTime {
			nums = append(nums, k)
		}
	}

	return nums
}
