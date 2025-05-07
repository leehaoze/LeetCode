package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// deleteMiddle 删除链表中的中间节点后，返回头指针
func deleteMiddle(head *ListNode) *ListNode {
	virtualHead := &ListNode{Next: head}
	cur := virtualHead

	fast, slow := cur, cur
	length := 0
	slowIdx := 0
	for fast.Next != nil {
		fast = fast.Next
		length++

		for slowIdx < length/2 {
			slow = slow.Next
			slowIdx++
		}
	}

	// 删除
	slow.Next = slow.Next.Next
	return virtualHead.Next
}
