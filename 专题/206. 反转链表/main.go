package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// o -> 1 -> 2 -> 3 -> 4
// o    1 -> 2    3 -> 4
// o    2 -> 1    3 -> 4
// o -> 2 -> 1 -> 3 -> 4

// // reverseList 将链表翻转后，返回头结点
// func reverseList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
//
// 	ret := reverseList(head.Next)
//
// 	head.Next.Next = head
// 	head.Next = nil
// 	return ret
// }

// reverseList 将链表翻转后，返回头结点
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

func main() {
	a := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}

	ret := reverseList(a)
	for ret != nil {
		fmt.Println(ret.Val)
		ret = ret.Next
	}
}
