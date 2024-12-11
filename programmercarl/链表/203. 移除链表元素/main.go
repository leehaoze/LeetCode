package main

// ListNode 是链表中的一个节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// removeElements 删除链表中的指定元素，并返回新链表的头结点
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Val == val {
		return removeElements(head.Next, val)
	} else {
		head.Next = removeElements(head.Next, val)
		return head
	}
}
