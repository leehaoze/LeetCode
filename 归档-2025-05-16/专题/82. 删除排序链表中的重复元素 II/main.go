package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
//
// 	if head.Next != nil && head.Val == head.Next.Val {
// 		subList := deleteDuplicates(head.Next)
// 		for subList != nil && subList.Val == head.Val {
// 			subList = subList.Next
// 		}
//
// 		return subList
// 	} else {
// 		head.Next = deleteDuplicates(head.Next)
// 		return head
// 	}
// }

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	virtual := &ListNode{Next: head}
	cur := virtual

	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			// 一直删除
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}

	return virtual.Next
}
