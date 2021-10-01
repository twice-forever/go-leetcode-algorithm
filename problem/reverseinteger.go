package problem

func Reverse(x int) int {
	n := 0

	for x != 0 {
		n = n*10 + x%10
		x = x / 10
	}

	if n < -2147483648 || n > 2147483647 {
		return 0
	}
	return n
}
