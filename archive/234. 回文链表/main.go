package palindrome_linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

var left *ListNode

func isPalindrome(head *ListNode) bool {
	left = head
	return traverse(head)
}

func traverse(head *ListNode) bool {
	if head == nil {
		return true
	}

	res := traverse(head.Next)
	res = res && (left.Val == head.Val)
	left = left.Next
	return res;
}
