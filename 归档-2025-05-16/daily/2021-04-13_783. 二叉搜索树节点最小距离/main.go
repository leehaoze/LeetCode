package main

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}

var minValue int
var preValue int
func minDiffInBST(root *TreeNode) int {
	minValue = 2147483647
	preValue = 0
	helper(root)
	return minValue
}

func helper(root *TreeNode) {
	if root == nil {
		return
	}

	helper(root.Left)
	if preValue != -1 {
		var tempVal int
		 if root.Val > preValue{
		 	tempVal = root.Val - preValue
			 if tempVal < minValue {
				 minValue = tempVal
			 }
		 } else if root.Val < preValue{
			 tempVal = preValue - root.Val
			 if tempVal < minValue {
				 minValue = tempVal
			 }
		 }

	}

	preValue = root.Val
	helper(root.Right)
}
