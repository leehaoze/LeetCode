package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	newHead := clone(head)
	headBak := head
	ret := newHead
	oldToNew := make(map[*Node]*Node)
	oldToNew[head] = newHead
	for head.Next != nil {
		newHead.Next = clone(head.Next)
		oldToNew[head.Next] = newHead.Next
		head = head.Next
		newHead = newHead.Next
	}

	newHeadBak := ret
	for headBak != nil {
		newHeadBak.Random = oldToNew[headBak.Random]
		headBak = headBak.Next
		newHeadBak = newHeadBak.Next
	}

	return ret
}

func clone(node *Node) *Node {
	return &Node{
		Val:    node.Val,
		Next:   nil,
		Random: nil,
	}
}
