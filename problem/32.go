package problem

type OneWord struct {
	Index int
	Word  byte
}

func longestValidParentheses(s string) int {
	parenthesis := []OneWord{}
	for i, v := range s {
		switch v {
		case '(':
			parenthesis = append(parenthesis, OneWord{Index: i, Word: '('})
		case ')':
			if len(parenthesis) == 0 {
				parenthesis = append(parenthesis, OneWord{Index: i, Word: ')'})
			} else if parenthesis[len(parenthesis)-1].Word == '(' {
				parenthesis = parenthesis[:len(parenthesis)-1]
			} else {
				parenthesis = append(parenthesis, OneWord{Index: i, Word: ')'})
			}
		}
	}
	ans := len(s)
	if len(parenthesis) > 0 {
		ans = parenthesis[0].Index
		if parenthesis[len(parenthesis)-1].Index != len(s)-1 {
			parenthesis = append(parenthesis, OneWord{Index: len(s) - 1})
		}
	}
	for i := 1; i < len(parenthesis); i++ {
		if parenthesis[i].Index-parenthesis[i-1].Index-1 > ans {
			ans = parenthesis[i].Index - parenthesis[i-1].Index - 1
		}
	}
	return ans
}
