package problem

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

var L11 = ListNode{
	Val:  2,
	Next: &L12,
}

var L12 = ListNode{
	Val:  4,
	Next: &L13,
}

var L13 = ListNode{
	Val: 3,
}

var L21 = ListNode{
	Val:  5,
	Next: &L22,
}

var L22 = ListNode{
	Val:  6,
	Next: &L23,
}

var L23 = ListNode{
	Val: 4,
}

func AddTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		l1 = &ListNode{
			Val: 0,
		}
	}
	if l2 == nil {
		l2 = &ListNode{
			Val: 0,
		}
	}

	var result ListNode
	result.Val = l1.Val + l2.Val
	if result.Val > 9 {
		result.Val = result.Val - 10
		if l1.Next == nil {
			l1.Next = &ListNode{
				Val: 0,
			}
		}
		l1.Next.Val++
	}
	if val := AddTwoNumbers1(l1.Next, l2.Next); val != nil {
		result.Next = val
	}

	return &result
}

func AddTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := 10
	carry := sum / 10
	fmt.Println(carry)
	return nil
}
