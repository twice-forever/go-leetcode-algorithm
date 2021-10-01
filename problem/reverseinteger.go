package problem

import (
	"strconv"
)

func Reverse(x int) int {
	negative := []byte("-")[0]
	s := strconv.Itoa(x)
	sLen := len(s)
	sIndex := sLen - 1
	newByte := make([]byte, 0, sLen)
	if s[0] == negative {
		newByte = append(newByte, negative)
	}

	for {
		if s[sIndex] == negative {
			break
		}

		newByte = append(newByte, s[sIndex])
		sIndex--

		if sIndex == -1 {
			break
		}
	}

	newS := string(newByte)
	newI, err := strconv.Atoi(newS)
	if err != nil {
		return 0
	}
	if int(int32(newI)) != newI {
		return 0
	}
	return newI
}
