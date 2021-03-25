package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pre = 0

func convertBST(root *TreeNode) *TreeNode {
	pre = 0
	traverse(root)
	return root
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}

	traverse(root.Right)
	root.Val = root.Val + pre
	pre = root.Val
	traverse(root.Left)
}

func main() {
	root := &TreeNode{
		Val: 4,
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
			Val: 6,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:  7,
				Left: nil,
				Right: &TreeNode{
					Val:   8,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}

	traverse(root)
}
