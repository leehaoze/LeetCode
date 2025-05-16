package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

//
// func removeElements(head *ListNode, val int) *ListNode {
// 	vHead := &ListNode{}
// 	vHead.Next = head
//
// 	p := vHead
//
// 	for p.Next != nil {
// 		if p.Next.Val == val {
// 			p.Next = p.Next.Next
// 		} else {
// 			p = p.Next
// 		}
// 	}
//
// 	return vHead.Next
// }

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}

	return head
}
