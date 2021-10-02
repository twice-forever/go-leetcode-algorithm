package problem

var (
	stringSlice = []byte("0123456789-+ ")
	stringMap   = map[byte]int{
		stringSlice[0]: 0,
		stringSlice[1]: 1,
		stringSlice[2]: 2,
		stringSlice[3]: 3,
		stringSlice[4]: 4,
		stringSlice[5]: 5,
		stringSlice[6]: 6,
		stringSlice[7]: 7,
		stringSlice[8]: 8,
		stringSlice[9]: 9,
	}
)

func MyAtoi(s string) int {
	needWord := false
	needSymbol := true
	isNegative := false
	returnNumber := 0

	for i := 0; i < len(s); i++ {
		if needWord || !needSymbol {
			if s[i] == stringSlice[10] || s[i] == stringSlice[11] || s[i] == stringSlice[12] {
				break
			}
		}
		if s[i] == stringSlice[10] {
			isNegative = true
			needWord = true
			continue
		} else if s[i] == stringSlice[11] {
			needWord = true
			continue
		} else if s[i] == stringSlice[12] {
			continue
		}

		v, ok := stringMap[s[i]]
		if ok {
			needSymbol = false
			returnNumber = returnNumber*10 + v
		} else {
			break
		}

		if returnNumber > 2147483647 {
			if isNegative {
				return -2147483648
			}
			return 2147483647
		}
	}

	if isNegative {
		returnNumber = returnNumber * -1
	}
	if returnNumber < -2147483648 {
		return -2147483648
	} else if returnNumber > 2147483647 {
		return 2147483647
	}
	return returnNumber
}
