package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root *TreeNode, leftNodes []int) *TreeNode {
	for idx, item := range leftNodes {
		var newLeftNodes = make([]int, 0)

		if idx != len(leftNodes)-1 {
			newLeftNodes = append(leftNodes[:idx], leftNodes[idx+1:]...)
		}

		if item < root.Val {
			root.Left = helper(&TreeNode{item, nil, nil}, newLeftNodes)
		} else {
			root.Right = helper(&TreeNode{item, nil, nil}, newLeftNodes)
		}
	}

	return root
}

func generateArray(n int, exclude int) []int {
	var ret []int

	for i := 1; i <= n; i++ {
		if i == exclude {
			continue
		} else {
			ret = append(ret, i)
		}
	}

	return ret
}

func generateTrees(n int) []*TreeNode {
	var ret []*TreeNode
	for i := 1; i <= n; i++ {
		ret = append(ret, helper(&TreeNode{i, nil, nil}, generateArray(n, i)))
	}
	return ret
}

func main() {
	ret := generateTrees(3)
	fmt.Printf("%v", ret)
}