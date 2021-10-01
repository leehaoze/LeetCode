package main

import "fmt"

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	head, _ := flattenList(root)
	return head
}

// 扁平化一个链表, 返回链表的头和尾
func flattenList(root *Node) (*Node, *Node) {
	retHead := root
	var retTail *Node

	// 找有child的节点 然后扁平化这个子链表
	for root != nil {
		// 记录一下next 一会把root指向这个next
		nextNode := root.Next
		// 如果next是nil了,那么他就是尾
		if nextNode == nil {
			retTail = root
		}

		// 如果有孩子链表,那么给他打散
		if root.Child != nil {
			childHead, childTail := flattenList(root.Child)

			// root 要跟 childHead 连接  nextNode 要跟 childTail连接
			root.Next = childHead
			childHead.Prev = root

			childTail.Next = nextNode
			if nextNode != nil {
				nextNode.Prev = childTail
			}

			// 孩子给干掉
			root.Child = nil

			// 因为追加了一个孩子链表在后边,所以retTail是孩子的tail
			retTail = childTail
		}

		root = nextNode
	}

	return retHead, retTail
}

func main() {
	nodes := make([]*Node, 13)
	for i := range nodes {
		if i == 0 {
			continue
		}
		nodes[i] = &Node{
			Val:   i,
			Prev:  nil,
			Next:  nil,
			Child: nil,
		}
	}

	// nodes[1].Next = nodes[2]
	// nodes[2].Prev = nodes[1]
	// nodes[2].Next = nodes[3]
	// nodes[3].Prev = nodes[2]
	// nodes[3].Next = nodes[4]
	// nodes[4].Prev = nodes[3]
	// nodes[4].Next = nodes[5]
	// nodes[5].Prev = nodes[4]
	// nodes[5].Next = nodes[6]
	// nodes[6].Prev = nodes[5]
	// nodes[6].Next = nil
	//
	// nodes[3].Child = nodes[7]
	// nodes[7].Next = nodes[8]
	// nodes[8].Prev = nodes[7]
	// nodes[8].Next = nodes[9]
	// nodes[9].Prev = nodes[8]
	// nodes[9].Next = nodes[10]
	// nodes[10].Prev = nodes[9]
	//
	// nodes[8].Child = nodes[11]
	//
	// nodes[11].Next = nodes[12]
	// nodes[12].Prev = nodes[11]

	nodes[1].Next = nil
	nodes[1].Child = nodes[2]
	nodes[2].Next = nil
	nodes[2].Child = nodes[3]

	ret := flatten(nodes[1])

	for ret != nil {
		fmt.Println(ret.Val)
		ret = ret.Next
	}

}
