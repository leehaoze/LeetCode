package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	res := -1
	for {
		if root.Left == nil || root.Right == nil {
			return res
		}

	}
}
