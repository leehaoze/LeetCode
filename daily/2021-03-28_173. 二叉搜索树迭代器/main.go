package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	innerRoot *Node
}

type Node struct {
	Val  int
	Next *Node
}

var head *Node

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	traverse(root.Left)
	head.Next = &Node{
		Val:  root.Val,
		Next: nil,
	}
	head = head.Next
	traverse(root.Right)
}

func Constructor(root *TreeNode) BSTIterator {
	head = &Node{}
	hh := head
	traverse(root)
	return BSTIterator{innerRoot: hh}
}

func (this *BSTIterator) Next() int {
	this.innerRoot = this.innerRoot.Next
	return this.innerRoot.Val
}

func (this *BSTIterator) HasNext() bool {
	return this.innerRoot.Next != nil
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

func main() {
	tree := &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   20,
				Left:  nil,
				Right: nil,
			},
		},
	}

	iterator := Constructor(tree)
	fmt.Println(iterator.Next())
	fmt.Println(iterator.Next())
	fmt.Println(iterator.HasNext())
	fmt.Println(iterator.Next())
	fmt.Println(iterator.HasNext())
	fmt.Println(iterator.Next())
	fmt.Println(iterator.HasNext())
	fmt.Println(iterator.Next())
	fmt.Println(iterator.HasNext())
}
