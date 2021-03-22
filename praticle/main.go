package main

import "fmt"

type TreeNode struct {
	Val        int
	LeftChild  *TreeNode
	RightChild *TreeNode
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}

	traverse(root.LeftChild)
	traverse(root.RightChild)
	fmt.Println(root.Val)
}

func main() {
	root := &TreeNode{
		Val: 10,
		LeftChild: &TreeNode{
			Val: 8,
			LeftChild: &TreeNode{
				Val:        6,
				LeftChild:  nil,
				RightChild: nil,
			},
			RightChild: &TreeNode{
				Val:        9,
				LeftChild:  nil,
				RightChild: nil,
			},
		},
		RightChild: &TreeNode{
			Val: 12,
			LeftChild: &TreeNode{
				Val:        13,
				LeftChild:  nil,
				RightChild: nil,
			},
			RightChild: &TreeNode{
				Val:        11,
				LeftChild:  nil,
				RightChild: nil,
			},
		},
	}

	traverse(root)
}
