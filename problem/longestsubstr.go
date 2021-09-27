package problem

func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	var subStrlength int = 1
	var head = 0
	var tail = 1
	sLen := len(s)

	for tail < sLen {
		tempStr := s[head:tail]
		tailWord := s[tail]
		tempStrLen := len(tempStr)
		hasSame := false
		for i := 0; i < tempStrLen; i++ {
			if tempStr[i] == tailWord {
				if subStrlength < tempStrLen {
					subStrlength = tempStrLen
				}
				hasSame = true
				head++
				tail = head + 1
				break
			}
		}
		if !hasSame {
			tail++
		}
	}

	if tail == sLen {
		tempStr := s[head:tail]
		if subStrlength < len(tempStr) {
			subStrlength = len(tempStr)
		}
	}

	return subStrlength
}
