package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	virtualHead := ListNode{}
	cur := &virtualHead

	var exceed int
	for l1 != nil || l2 != nil || exceed == 1 {
		l1Val, l2Val := 0, 0
		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}

		ret := l1Val + l2Val + exceed

		if ret >= 10 {
			cur.Next = &ListNode{Val: ret % 10}
			exceed = 1
		} else {
			cur.Next = &ListNode{Val: ret}
			exceed = 0
		}

		cur = cur.Next
	}

	return virtualHead.Next
}
