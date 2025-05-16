package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for slow != nil && fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		} else {
			return false
		}

		if slow == fast {
			return true
		}
	}

	return false
}
