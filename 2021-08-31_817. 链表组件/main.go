package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, nums []int) int {
	bucket := make(map[int]struct{})
	for idx := range nums {
		bucket[nums[idx]] = struct{}{}
	}
	ret := 0
	newList := true
	for head != nil {
		if _, exists := bucket[head.Val]; !exists {
			newList = true
			head = head.Next
			continue
		}

		if newList {
			ret += 1
			newList = false
		}



		head = head.Next
	}

	return ret
}

func main() {
	head := &ListNode{
		Val:  3,
		Next: &ListNode{
			Val:  4,
			Next: &ListNode{
				Val:  0,
				Next: &ListNode{
					Val:  2,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}

	fmt.Println(numComponents(head, []int{4}))
}
