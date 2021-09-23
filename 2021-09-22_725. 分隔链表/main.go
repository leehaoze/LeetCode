package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	length := 0
	ptr := head
	for ptr != nil {
		length++
		ptr = ptr.Next
	}

	// 计算每组应该多少个节点
	averCount := length / k
	leftCount := length % k

	ret := make([]*ListNode, k, k)

	lengthOfGroup := make([]int, k, k)
	for i := 0; i < k; i++ {
		lengthOfGroup[i] = averCount
		if leftCount > 0 {
			lengthOfGroup[i] += 1
			leftCount -= 1
		}
	}

	ptr = head
	for idx, l := range lengthOfGroup {
		if l <= 0 {
			continue
		}

		ret[idx] = ptr
		for i := 0; i < l; i++ {
			if i == l-1 {
				next := ptr.Next
				ptr.Next = nil
				ptr = next
			} else {
				ptr = ptr.Next
			}
		}

	}

	return ret
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val: 7,
								Next: &ListNode{
									Val: 8,
									Next: &ListNode{
										Val: 9,
										Next: &ListNode{
											Val:  10,
											Next: nil,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	ret := splitListToParts(head, 3)

	for _, node := range ret {
		for node != nil {
			fmt.Print(node.Val, ",")
			node = node.Next
		}
		fmt.Println("")
	}
	fmt.Println("end")
}
