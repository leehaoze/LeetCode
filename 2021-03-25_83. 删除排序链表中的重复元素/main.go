package _021_03_25_83__删除排序链表中的重复元素

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	childList := deleteDuplicates(head.Next)
	if childList == nil || head.Val != childList.Val {
		head.Next = childList
	} else {
		head.Next = childList.Next
	}

	return head
}
