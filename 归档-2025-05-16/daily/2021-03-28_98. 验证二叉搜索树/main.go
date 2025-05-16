package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pre = math.MinInt64

func helper(root *TreeNode) bool {
	//return isValid(root, nil, nil)
	if root == nil {
		return true
	}

	if !helper(root.Left) {
		return false
	}

	if root.Val <= pre {
		return false
	}
	pre = root.Val

	return helper(root.Right)
}

func isValidBST(root *TreeNode) bool {
	//return isValid(root, nil, nil)
	pre = math.MinInt64
	return helper(root)
}

func isValid(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}

	if min != nil && root.Val <= min.Val {
		return false
	}

	if max != nil && root.Val >= max.Val {
		return false
	}

	return isValid(root.Left, min, root) && isValid(root.Right, root, max)
}

func main() {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:  2,
				Left: nil,
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
		},
	}

	res := isValidBST(tree)
	fmt.Println(res)
}
