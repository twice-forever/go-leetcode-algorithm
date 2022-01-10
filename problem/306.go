package problem

import (
	"fmt"
	"strconv"
)

func IsAdditiveNumber(num string) bool {
	nLen := len(num)
	if nLen < 3 {
		return false
	}
	first, second := 0, 1
	for second < nLen {
		if first+1 == second {
			first = 1
			second++
		} else {
			first++
		}
		if first > 1 && num[0] == '0' || num[first] == '0' && second-first > 1 {
			continue
		}
		num1, _ := strconv.Atoi(num[:first])
		num2, _ := strconv.Atoi(num[first:second])
		sum := num1 + num2
		strSum := strconv.Itoa(sum)
		if second+len(strSum) > nLen {
			continue
		}
		if strSum == num[second:second+len(strSum)] {
			match := matchAdditiveNumber(num, num1, num2)
			if match {
				return match
			}
		}
	}
	return false
}

func matchAdditiveNumber(num string, num1, num2 int) bool {
	index := len(fmt.Sprintf("%d%d", num1, num2))
	nLen := len(num)
	for index < nLen {
		sum := num1 + num2
		strSum := strconv.Itoa(sum)
		if index+len(strSum) > nLen {
			return false
		}
		if strSum != num[index:index+len(strSum)] {
			return false
		} else {
			index += len(strSum)
			num1, num2 = num2, sum
		}
	}
	return true
}
