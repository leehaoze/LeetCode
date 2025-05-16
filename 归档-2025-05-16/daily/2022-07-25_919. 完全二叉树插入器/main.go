package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CBTInserter struct {
	root  *TreeNode
	nodes []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	return CBTInserter{
		root:  root,
		nodes: []*TreeNode{root},
	}
}

func (this *CBTInserter) Insert(val int) int {
	parentVal := 0
	idx := 0
	length := len(this.nodes)
	for i := 0; i < length; i++ {
		node := this.nodes[i]
		if node.Left == nil {
			node.Left = &TreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
			parentVal = node.Val
			idx = i
			this.nodes = append(this.nodes, node.Left)
			length++
			break
		}
		this.nodes = append(this.nodes, node.Left)
		length++

		if node.Right == nil {
			node.Right = &TreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
			parentVal = node.Val
			idx = i
			this.nodes = append(this.nodes, node.Right)
			length++
			// 这种情况应该把他踢出去

			break
		}

		this.nodes = append(this.nodes, node.Right)
		length++
	}

	this.nodes = this.nodes[idx:]
	return parentVal
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

/*
     1
   /   \
  2     3
 / \   / \
4   5 6
*/

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
