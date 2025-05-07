package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// reverse 翻转 head -> end之间的链表
func reverse(head *ListNode, end *ListNode) *ListNode {
	prev := end.Next
	for head != end {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return end
}

// reverseKGroup 将传入的链表，每K个节点，翻转一次
// func reverseKGroup(head *ListNode, k int) *ListNode {
// // 虚节点
// virtualHead := &ListNode{Next: head}
//
// pre := virtualHead
//
// for head != nil {
// 	for i := 0; i < k; i++ {
// 		head = head.Next
// 		if head == nil {
// 			return virtualHead
// 		}
// 	}
//
// 	// 翻转这部分
// 	pre.Next = reverse(pre.Next, head)
// 	pre = head
// }
// }
