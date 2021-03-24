package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	solve(root)
}

func solve(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	right := root.Right
	left := root.Left

	root.Left = nil
	if left != nil {
		root.Right = left
		leftLast := solve(left)
		leftLast.Right = right
		if right == nil {
			return leftLast
		}
	} else {
		root.Right = right
	}
	return solve(right)
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	traverse(root.Left)
	traverse(root.Right)
}

func main() {
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 12,
			Left: &TreeNode{
				Val:   13,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   11,
				Left:  nil,
				Right: nil,
			},
		},
	}

	traverse(root)

}
