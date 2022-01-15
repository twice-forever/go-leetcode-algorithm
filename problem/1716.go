package problem

func TotalMoney(n int) int {
	num, ans := 0, 0
	index := 0
	for i := 0; i < n; i++ {
		if i%7 == 0 {
			num++
			index = 0
		}
		ans += num + index*1
		index++
	}
	return ans
}
