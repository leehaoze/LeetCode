package main

import (
	"fmt"
	"strings"
)

// printArrayWithPointers 打印数组和指针位置
func printArrayWithPointers(nums []int, left, right int) {

	// 打印指针位置
	leftPointer := make([]string, len(nums))
	rightPointer := make([]string, len(nums))

	for i := 0; i < len(nums); i++ {
		if i == left {
			leftPointer[i] = "^"
		} else {
			leftPointer[i] = " "
		}

		if i == right {
			rightPointer[i] = "v"
		} else {
			rightPointer[i] = " "
		}
	}

	fmt.Println(" " + strings.Join(rightPointer, " "))
	// 打印数组
	fmt.Println(nums)
	fmt.Println(" " + strings.Join(leftPointer, " "))
	fmt.Println("")
}

func removeDuplicates(nums []int) int {
	left, right := 0, 0

	for right < len(nums) {
		if left-2 < 0 || nums[left-2] != nums[right] {
			nums[left] = nums[right]
			left++
		}

		right++
		// printArrayWithPointers(nums, left, right)
	}

	return left
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
}
