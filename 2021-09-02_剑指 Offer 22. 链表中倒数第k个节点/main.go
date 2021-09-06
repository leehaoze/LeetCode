package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	step := k
	fastHead := head
	for i := 0; i < step; i++ {
		fastHead = fastHead.Next
	}

	for fastHead != nil {
		head = head.Next
		fastHead = fastHead.Next
	}

	return head
}


func main() {
	head := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	fmt.Println(getKthFromEnd(head, 5).Val)
}
