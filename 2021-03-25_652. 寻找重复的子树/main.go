package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var treeMap = make(map[string]int)
var repeatTreeRoot = make([]*TreeNode, 0)

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	treeMap = make(map[string]int)
	repeatTreeRoot = make([]*TreeNode, 0)
	traverse(root)
	return repeatTreeRoot
}

func traverse(root *TreeNode) string {
	if root == nil {
		return "#"
	}

	leftTree := traverse(root.Left)
	rightTree := traverse(root.Right)

	subTree := fmt.Sprintf("%s,%s,%d", leftTree, rightTree, root.Val)
	if _, exists := treeMap[subTree]; !exists {
		treeMap[subTree] = 1
	} else {
		treeMap[subTree] = treeMap[subTree] + 1
		if treeMap[subTree] == 2 {
			repeatTreeRoot = append(repeatTreeRoot, root)
		}
	}

	return subTree
}

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
	}

	findDuplicateSubtrees(tree)
}
