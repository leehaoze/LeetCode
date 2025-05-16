package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeLinkList 合并两个链表，并返回合并后的新链表
func mergeLinkList(a, b *ListNode) *ListNode {
	virtualHead := &ListNode{}
	currentNode := virtualHead

	for a != nil || b != nil {
		if a == nil {
			currentNode.Next = b
			break
		} else if b == nil {
			currentNode.Next = a
			break
		} else {
			if a.Val < b.Val {
				currentNode.Next = a
				currentNode = currentNode.Next
				a = a.Next
			} else {
				currentNode.Next = b
				currentNode = currentNode.Next
				b = b.Next
			}
		}
	}

	return virtualHead.Next
}

// mergeTwoList 将两个有序链表，合并到一起，返回新链表的头结点
func mergeTwoList(a, b *ListNode) *ListNode {
	// 终止条件
	if a == nil {
		return b
	}

	if b == nil {
		return a
	}

	// 如果 a链表的当前值小于b链表的当前值
	// 则当前节点使用a链表节点，剩余的进行合并
	if a.Val < b.Val {
		a.Next = mergeTwoList(a.Next, b)
		return a
	} else {
		b.Next = mergeTwoList(b.Next, a)
		return b
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}

// merge 将数组中指定范围(start, end)的链表，合并成一个有序链表，返回链表的头结点
func merge(lists []*ListNode, start, end int) *ListNode {
	if start == end {
		return lists[start]
	}
	if end-start == 1 {
		return mergeTwoList(lists[end], lists[start])
	}

	mid := start + (end-start)/2
	return mergeTwoList(merge(lists, start, mid), merge(lists, mid+1, end))
}
