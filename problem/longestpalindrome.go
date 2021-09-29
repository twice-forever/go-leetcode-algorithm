package problem

func LongestPalindrome(s string) string {
	head := 0
	stringLen := len(s)
	subStringLen := stringLen
	var isPalindrome bool

	for {
		if subStringLen == 1 {
			return string(s[0])
		}

		subString := s[head : subStringLen+head]
		compareLen := subStringLen / 2
		compareOne := subString[0:compareLen]
		compareTwo := subString[subStringLen-compareLen:]
		isPalindrome = true
		for i := 0; i < compareLen; i++ {
			if compareOne[i] != compareTwo[compareLen-1-i] {
				isPalindrome = false
				break
			}
		}
		if isPalindrome {
			return subString
		}
		if head+subStringLen == stringLen {
			head = 0
			subStringLen--
			continue
		}
		head++
	}
}
