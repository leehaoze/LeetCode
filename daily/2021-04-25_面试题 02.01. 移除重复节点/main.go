package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}


func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 这个链表已经去重了
	subListNode := removeDuplicateNodes(head.Next)
	head.Next = removeDumplicate(subListNode, head.Val)
	return head
}

func removeDumplicate(head *ListNode, val int) *ListNode {
	newHead := &ListNode{
		Next: head,
	}

	itreator := newHead
	for itreator != nil && itreator.Next != nil {
		if itreator.Next.Val == val {
			itreator.Next = itreator.Next.Next
		}
		itreator = itreator.Next
	}

	return newHead.Next
}


func main() {
	head := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  3,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
		},
	}

	head = removeDuplicateNodes(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}