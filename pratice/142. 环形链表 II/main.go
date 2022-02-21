package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	fast, slow := head.Next, head
	for fast != slow {
		if fast.Next == nil || slow == nil {
			return nil
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return fast
}
