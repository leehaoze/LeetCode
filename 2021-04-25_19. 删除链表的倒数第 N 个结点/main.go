package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//head, _ = helper(head, n)
	//return head
	newHead := &ListNode{
		Next: head,
	}
	slow, fast := newHead, newHead

	for i := 0; i < n; i++ {
		fast = fast.Next
	}


	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next
	return newHead.Next
}

func helper(head *ListNode, n int) (*ListNode, int) {
	if head.Next == nil {
		if n == 1 {
			return nil, 1
		}
		return head, 1
	}

	next, count := helper(head.Next, n)
	head.Next = next
	if count == n-1 {
		return next, count + 1
	} else {
		return head, count + 1
	}
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	removeNthFromEnd(head, 2)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
