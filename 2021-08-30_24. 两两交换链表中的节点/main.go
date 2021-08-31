package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
1 -> 2 -> 3 -> 4
*/

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	subList := swapPairs(head.Next.Next)
	head.Next.Next = head
	head.Next = subList
	return newHead
}

func main() {
	// url := "https://ui-avatars.com/api/?background=ffffff&color=000000&name=%v"
	idToNameMap := map[int]int{
		3634: 01,
		3635: 02,
		3636: 03,
		3637: 04,
		3638: 05,
		3639: 06,
		3640: 07,
		3641: 8,
		3642: 9,
		3643: 10,
		3644: 11,
		3645: 12,
		3646: 13,
		3647: 14,
		3648: 15,
		3649: 16,
		3650: 17,
		3651: 18,
		3652: 21,
		3653: 22,
		3654: 23,
		3655: 24,
		3656: 19,
		3657: 20,
	}

	newMap := make(map[int]int)
	for key, value := range idToNameMap {
		newMap[value] = key
	}

	name := 3634
	for i := 1; i <= 24; i++ {
		fmt.Printf("%02d\n", i)
		url := fmt.Sprintf("https://ui-avatars.com/api/?background=ffffff&color=000000&name=%02d", i)
		resp, _ := http.Get(url)
		body, _ := ioutil.ReadAll(resp.Body)
		out, _ := os.Create(fmt.Sprintf("face_img_%02d.png", newMap[i]))
		io.Copy(out, bytes.NewReader(body))
		name++
	}

	// list := &ListNode{
	// 	Val:  1,
	// 	Next: &ListNode{
	// 		Val:  2,
	// 		Next: &ListNode{
	// 			Val:  3,
	// 			Next: &ListNode{
	// 				Val:  4,
	// 				Next: &ListNode{
	// 					Val:  5,
	// 					Next: nil,
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	//
	// ret := swapPairs(list)
	// for ret != nil {
	// 	fmt.Println(ret.Val)
	// 	ret = ret.Next
	// }
}
