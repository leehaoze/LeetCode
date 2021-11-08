package main

import "fmt"

type ListNode struct {
     Val int
     Next *ListNode
 }

// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
// 	head := &ListNode{}
// 	tail := head
//
// 	for l1 != nil && l2 != nil {
// 		if l1.Val > l2.Val {
// 			tail.Next = l2
// 			l2 = l2.Next
// 		} else {
// 			tail.Next = l1
// 			l1 = l1.Next
// 		}
//
// 		tail = tail.Next
// 	}
//
// 	for l1 != nil {
// 		tail.Next = l1
// 		tail = tail.Next
// 		l1 = l1.Next
// 	}
//
// 	for l2 != nil {
// 		tail.Next = l2
// 		tail = tail.Next
// 		l2 = l2.Next
// 	}
//
// 	return head.Next
// }

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

func main() {
	l1 := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	ret := mergeTwoLists(l1, l2)

	for ret != nil {
		fmt.Println(ret.Val)
		ret = ret.Next
	}
}