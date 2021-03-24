package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return construct(nums, 0, len(nums))
}

func construct(nums []int, start, end int) *TreeNode {
	if start == end {
		return nil
	}

	rootIdx := findMaxValue(nums, start, end)
	root := &TreeNode{
		Val:   nums[rootIdx],
		Left:  construct(nums, start, rootIdx),
		Right: construct(nums, rootIdx+1, end),
	}
	return root
}

func findMaxValue(nums []int, start, end int) int {
	maxValue := -1
	idx := 0
	for i := start; i < end; i++ {
		if nums[i] > maxValue {
			maxValue = nums[i]
			idx = i
		}
	}

	return idx
}

func main() {
	root := constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
	travse(root)
}

func travse(root *TreeNode) {
	if root == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println(root.Val)
	travse(root.Left)
	travse(root.Right)
}
