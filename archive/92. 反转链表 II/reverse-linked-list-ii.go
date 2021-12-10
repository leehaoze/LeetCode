package leetcode

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

var next *ListNode

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		next = head.Next
		return head
	}

	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = next
	return last
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right)
	}

	head.Next = reverseBetween(head.Next, left - 1, right - 1)
	return head
}
