package problem

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x >= 0 && x < 10 {
		return true
	}

	n := 0
	y := x
	for y > 0 {
		n = n*10 + y%10
		y = y / 10
	}

	return n == x
}
