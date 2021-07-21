package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	lenA, lenB := 0, 0
// 	hA, hB := headA, headB
// 	for hA != nil {
// 		hA = hA.Next
// 		lenA++
// 	}
//
// 	for hB != nil {
// 		hB = hB.Next
// 		lenB++
// 	}
//
// 	if lenA > lenB {
// 		hA = headA
// 		hB = headB
// 	} else {
// 		hA = headB
// 		hB = headA
// 		lenA, lenB = lenB, lenA
// 	}
//
// 	for i := 0; i < lenA-lenB; i++ {
// 		hA = hA.Next
// 	}
//
// 	return solve(hA, hB)
// }
//
// func solve(headA, headB *ListNode) *ListNode {
// 	if headA == nil || headB == nil {
// 		return nil
// 	}
//
// 	if headA == headB {
// 		return headA
// 	}
//
// 	return solve(headA.Next, headB.Next)
// }
//
func constructList(list []int) *ListNode {
	head := &ListNode{}
	ret := head
	for idx, val := range list {
		head.Val = val
		if idx != len(list)-1 {
			head.Next = &ListNode{}
			head = head.Next
		}
	}
	return ret
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	hA, hB := headA, headB

	for hA != hB {
		if hA == nil {
			hA = headB
		} else {
			hA = hA.Next
		}

		if hB == nil {
			hB = headA
		} else {
			hB = hB.Next
		}
	}

	return hA
}

func main() {
	ha := constructList([]int{4, 1, 8, 4, 5})
	hb := constructList([]int{5,0,1,8,4,5})
	fmt.Println(getIntersectionNode(ha, hb))
}
