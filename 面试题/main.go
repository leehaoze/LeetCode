package main

import "fmt"

type Node struct {
	Val int
	Next *Node
}

func deleteTargetElement(head *Node, target int) *Node {
	// }
	// 1
	// 1 -> 2
	// 1 -> 2 -> 3
	tempHead := &Node{}
	bak := tempHead
	tempHead.Next = head
	for tempHead != nil && tempHead.Next != nil {
		if tempHead.Next.Val == target {
			tempHead.Next = tempHead.Next.Next
		}
		tempHead = tempHead.Next
	}

	return bak.Next
}

func main() {
	head := &Node{
		Val:  1,
		Next: &Node{
			Val:  2,
			Next: nil,
		},
	}

	ret := deleteTargetElement(head, 2)
	for ret != nil {
		fmt.Println(ret.Val)
		ret = ret.Next
	}
}
