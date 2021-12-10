package main

func flattenV2(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	rightTree := root.Right
	root.Right = root.Left
	root.Left = nil

	for root.Right != nil {
		root = root.Right
	}

	root.Right = rightTree
}
