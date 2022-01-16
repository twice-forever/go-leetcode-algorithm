package problem

import "math/rand"

type Solution struct {
	head *ListNode
}

func Constructoriiiiiii(head *ListNode) Solution {
	return Solution{
		head: head,
	}
}

func (this *Solution) GetRandom() int {
	ans := 0
	for node, i := this.head, 1; node != nil; node = node.Next {
		if rand.Intn(i) == 0 {
			ans = node.Val
		}
		i++
	}
	return ans
}
