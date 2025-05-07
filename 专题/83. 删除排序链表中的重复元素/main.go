package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// deleteDuplicates 删除一个排序链表中的重复元素，只保留一个
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	ret := deleteDuplicates(head.Next)

	if ret == nil {
		return head
	}

	if head.Val == ret.Val {
		return ret
	}

	head.Next = ret
	return head
}
