package main

import (
	"fmt"
)

// func removeElement(nums []int, val int) int {
// 	length := len(nums)
// 	for i := 0; i < length; i++ {
// 		if nums[i] == val {
// 			// 后边覆盖前边的
// 			for j := i; j+1 < length; j++ {
// 				nums[j] = nums[j+1]
// 			}
// 			length--
// 			i--
// 		}
// 	}
//
// 	return length
// }

func removeElement(nums []int, val int) int {
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}
		right++
	}

	return left
}

func main() {
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
