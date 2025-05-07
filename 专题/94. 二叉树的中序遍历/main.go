package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func inorderTraversal(root *TreeNode) []int {
// 	if root == nil {
// 		return []int{}
// 	}
// 	ret := make([]int, 0)
// 	ret = append(ret, inorderTraversal(root.Left)...)
// 	ret = append(ret, root.Val)
// 	ret = append(ret, inorderTraversal(root.Right)...)
// 	return ret
// }

func inorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)

	ret := make([]int, 0)
	// 对于每个节点，先压右，再压中，最后压左
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, root.Val)
		root = root.Right
	}

	return ret
}
