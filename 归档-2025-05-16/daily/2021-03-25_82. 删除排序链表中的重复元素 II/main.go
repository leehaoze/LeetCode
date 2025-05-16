package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	head, _ = helper(head)
	return head
}

func helper(head *ListNode) (*ListNode, int) {
	if head == nil {
		return nil, -1000
	}

	if head.Next == nil {
		return head, -1000
	}

	nextHead, repeatVal := helper(head.Next)

	if nextHead != nil {
		if head.Val == nextHead.Val {
			return nextHead.Next, head.Val
		}
	}

	if head.Val == repeatVal {
		return nextHead, repeatVal
	}

	head.Next = nextHead
	return head, -1000
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
			},
		},
	}

	head = deleteDuplicates(head)

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
