package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := make([]*TreeNode, 0)
	ret := make([]int, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, root.Val)
		if root.Right != nil {
			stack = append(stack, root.Right)
		}

		if root.Left != nil {
			stack = append(stack, root.Left)
		}
	}

	return ret
}
