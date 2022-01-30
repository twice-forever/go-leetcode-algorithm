package problem

func maxProfitii(prices []int) int {
	ans := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}
