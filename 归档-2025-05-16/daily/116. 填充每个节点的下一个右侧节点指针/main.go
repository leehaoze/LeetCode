package main


type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	connectTwoNode(root.Left, root.Right)
	return root
}

func connectTwoNode(nodeA *Node, nodeB *Node) {
	if nodeA == nil || nodeB == nil {
		return
	}

	nodeA.Next = nodeB
	connectTwoNode(nodeA.Left, nodeA.Right)
	connectTwoNode(nodeB.Left, nodeB.Right)
	connectTwoNode(nodeA.Right, nodeB.Left)
}
