package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	slow, fast := head, head
	length := 0
	// move the fast to make the fast one larger k nodes than the slow one
	for i := 0; i < k; i++ {
		// the last node
		if fast.Next == nil {
			// remember the length
			length = i + 1
			break
		}
		fast = fast.Next
	}

	if length != 0 {
		// if the length is not zero which means k is larger than the single link list
		k = k % length
		// remove the fast
		fast = head
		for i := 0; i < k; i++ {
			fast = fast.Next
		}
	}

	if fast == head {
		return head
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	newHead := slow.Next
	slow.Next = nil
	fast.Next = head
	return newHead
}

func main() {
	//head := &ListNode{
	//	Val: 0,
	//	Next: &ListNode{
	//		Val: 1,
	//		Next: &ListNode{
	//			Val:  2,
	//			Next: nil,
	//		},
	//	},
	//}
	head := &ListNode{
		Val:  1,
		Next: nil,
	}

	head = rotateRight(head, 1)

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
