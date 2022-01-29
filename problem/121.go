package problem

func maxProfit(prices []int) int {
	ans, n := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		if n > prices[i] {
			n = prices[i]
		} else if ans < prices[i]-n {
			ans = prices[i] - n
		}
	}
	return ans
}
