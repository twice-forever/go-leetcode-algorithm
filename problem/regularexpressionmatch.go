package problem

func IsMatch(s string, p string) bool {
	sHead := 0
	sLen := len(s)
	pLen := len(p)
	pHead := 0
	reg := ".*"

	if p == reg {
		return true
	} else if s == "" && p == "" {
		return true
	} else if p == "" || s == "" {
		return false
	}

	for pHead < pLen && sHead < sLen {
		if pHead+1 < pLen && p[pHead+1] == reg[1] {
			if p[pHead] == reg[0] {
				if pHead+2 < pLen {
					tar := p[pHead+2]
					for {
						if s[sHead] == tar {
							pHead = pHead + 2
							break
						}
						sHead++
						if sHead == sLen {
							return false
						}
					}
				}
			} else if s[sHead] != p[pHead] {
				pHead = pHead + 2
			} else {
				tar := p[pHead]
				for {
					if s[sHead] != tar {
						pHead = pHead + 2
						break
					}
					sHead++
					if sHead == sLen {
						return true
					}
				}
			}
			// fmt.Println(string(p[0 : pHead+1]))
			// fmt.Println(string(s[0 : sHead+1]))
		} else if p[pHead] == reg[0] || p[pHead] == s[sHead] {
			sHead++
			pHead++
		} else {
			return false
		}

		if pHead == pLen && sHead < sLen {
			return false
		} else if sHead == sLen && pHead < pLen {
			return false
		}
	}
	return true
}
