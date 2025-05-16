package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fast, slow := head, head

	for fast != nil && slow != nil {
		fast = fast.Next
		slow = slow.Next

		if fast != nil {
			fast = fast.Next
		}

		if fast == slow && fast != nil {
			return true
		}
	}

	return false
}

func main() {

}
