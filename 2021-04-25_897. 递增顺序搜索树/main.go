package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pre *TreeNode
var head *TreeNode

func increasingBST(root *TreeNode) *TreeNode {
	pre = nil
	head = nil

	helper(root)
	return head
}

func helper(root *TreeNode) {
	if root == nil {
		return
	}

	helper(root.Left)
	//pre.Right = root
	if pre != nil {
		pre.Right = root
		pre.Left = nil
	}

	if head == nil {
		head = root
	}

	pre = root
	helper(root.Right)
	root.Left = nil
}

func main() {
	//tree := &TreeNode{
	//	Val: 5,
	//	Left: &TreeNode{
	//		Val: 3,
	//		Left: &TreeNode{
	//			Val: 2,
	//			Left: &TreeNode{
	//				Val:   1,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//			Right: nil,
	//		},
	//		Right: &TreeNode{
	//			Val:   4,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val:  6,
	//		Left: nil,
	//		Right: &TreeNode{
	//			Val: 8,
	//			Left: &TreeNode{
	//				Val:   7,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//			Right: &TreeNode{
	//				Val:   9,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//		},
	//	},
	//}
	
	tree := &TreeNode{
		Val:   2,
		Left:  &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   4,
			Left:  &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}

	ans := increasingBST(tree)
	fmt.Println(ans.Val)
}
