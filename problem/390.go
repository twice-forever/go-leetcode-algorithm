package problem

func LastRemaining(n int) int {
	if n == 1 {
		return 1
	}
	if n < 6 {
		return 2
	}
	inter := 2
	if n%2 != 0 {
		n -= 1
	}
	scope := [...]int{2, n}
	for i := 0; scope[0] != scope[1]; i++ {
		if i%2 == 0 {
			if (scope[1]-scope[0])%(inter*2) == 0 {
				scope[0] = scope[0] + inter
			}
			scope[1] = scope[1] - inter
		} else {
			if (scope[1]-scope[0])%(inter*2) == 0 {
				scope[1] = scope[1] - inter
			}
			scope[0] = scope[0] + inter
		}
		inter *= 2
	}
	return scope[0]
}
