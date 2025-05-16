package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 1)
	res[0] = root.Val
	for _, c := range root.Children {
		res = append(res, preorder(c)...)
	}

	return res
}
