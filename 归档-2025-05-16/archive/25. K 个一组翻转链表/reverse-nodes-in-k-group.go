package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// reverse 反转到end节点
//func reverse(head *ListNode, end *ListNode) *ListNode {
//	pre := head
//	cur := head
//	next := head
//	for cur != end {
//		next = cur.Next
//		cur.Next = pre
//		pre = cur
//		cur = next
//	}
//
//	return pre
//}

func reverse(head *ListNode, end *ListNode) *ListNode {
	// 反转到head就可以了
	if head.Next == end {
		return head
	}
	newHead := reverse(head.Next ,end)
	head.Next.Next = head
	head.Next = end
	return newHead
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil{
			return head
		}
		b = b.Next
	}

	newHead := reverse(head, b)
	a.Next= reverseKGroup(b, k)
	return newHead
}
