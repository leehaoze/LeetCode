package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaf1 := traverse(root1)
	leaf2 := traverse(root2)

	if len(leaf1) != len(leaf2) {
		return false
	}

	for i := range leaf1 {
		if leaf1[i] != leaf2[i] {
			return false
		}
	}

	return true
}

func traverse(root *TreeNode) []int {
	if root.Left == nil && root.Right == nil {
		// 叶子节点
		return []int{root.Val}
	}

	ans := make([]int, 0)
	if root.Left != nil {
		leftAns := traverse(root.Left)
		if len(leftAns) != 0 {
			ans = append(ans, leftAns...)
		}
	}

	if root.Right != nil {
		rightAns := traverse(root.Right)
		if len(rightAns) != 0 {
			ans = append(ans, rightAns...)
		}
	}
	return ans
}

func main() {
	//root := &TreeNode{
	//	Val:   3,
	//	Left:  &TreeNode{
	//		Val:   5,
	//		Left:  &TreeNode{
	//			Val:   6,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: &TreeNode{
	//			Val:   2,
	//			Left:  &TreeNode{
	//				Val:   7,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//			Right: &TreeNode{
	//				Val:   4,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val:   1,
	//		Left:  &TreeNode{Val: 9},
	//		Right: &TreeNode{Val: 8},
	//	},
	//}
	//
	//fmt.Println(traverse(root))
	root1 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: nil,
	}

	root2 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: nil,
	}

	fmt.Println(leafSimilar(root1, root2))
}
