package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var arr = make([]int, 0)

func kthSmallest(root *TreeNode, k int) int {
	arr = make([]int, 0)
	traverse(root)
	return arr[k-1]
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}

	traverse(root.Left)
	arr = append(arr, root.Val)
	traverse(root.Right)
}

func main() {
	tree := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
	}

	ret2 := kthSmallest(tree, 1)
	fmt.Println(ret2)
}
