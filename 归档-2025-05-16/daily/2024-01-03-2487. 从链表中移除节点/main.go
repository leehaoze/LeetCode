package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNodes(head *ListNode) *ListNode {
	return helper(head)
}

// helper 返回一条符合题目规则的链表
// 题目规则：每个节点的右侧，没有比他大的节点
func helper(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 获取右侧链表
	head.Next = helper(head.Next)

	if head.Next == nil {
		return head
	}

	// 比较大小
	if head.Next.Val > head.Val {
		head = head.Next
	}

	return head
}
