package problem

func ZigzagConvert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	sLen := len(s)
	storeSlice := make([][]byte, numRows)
	for i := range storeSlice {
		storeSlice[i] = make([]byte, 0, sLen)
	}

	verticalIndex := 0
	sIndex := 0
	reachEnd := false
	for sIndex < sLen {
		storeSlice[verticalIndex] = append(storeSlice[verticalIndex], s[sIndex])
		sIndex++
		if verticalIndex == numRows-1 {
			reachEnd = true
		}
		if verticalIndex == 0 {
			reachEnd = false
		}
		if !reachEnd {
			verticalIndex++
		} else {
			verticalIndex--
		}
	}

	byteSlice := make([]byte, 0, sLen)
	for _, v := range storeSlice {
		byteSlice = append(byteSlice, v...)
	}
	return string(byteSlice)
}
