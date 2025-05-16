package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	list := make([]*TreeNode, 0)
	if root != nil {
		list = append(list, root)
	}

	for len(list) > 0 {
		nodeCount := len(list)
		currentNodeVal := make([]int, 0)
		for i := 0; i < nodeCount; i++ {
			currentNodeVal = append(currentNodeVal, list[i].Val)
			if list[i].Left != nil {
				list = append(list, list[i].Left)
			}
			if list[i].Right != nil {
				list = append(list, list[i].Right)
			}
		}
		ret = append(ret, currentNodeVal)

		list = list[nodeCount:]
	}

	return ret
}
