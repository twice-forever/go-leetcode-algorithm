package problem

import (
	"math"
)

func PathInZigZagTree(label int) []int {
	level, location := 1, 0
	for {
		sum := math.Pow(2, float64(level)) - 1
		if sum >= float64(label) {
			if level%2 == 0 {
				location = int(sum - float64(label))
			} else {
				location = int(float64(label) - math.Pow(2, float64(level-1)))
			}
			break
		}
		level++
	}
	ans := make([]int, level)
	ans[level-1] = label
	level = level - 1
	for ; level > 0; level-- {
		location = location / 2
		if level%2 == 0 {
			ans[level-1] = int(math.Pow(2, float64(level))-1) - location
		} else {
			ans[level-1] = int(math.Pow(2, float64(level-1))) + location
		}
	}
	return ans
}
