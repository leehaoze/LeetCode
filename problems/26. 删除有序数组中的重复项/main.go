package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	left, right := 0, 0
	for right < len(nums) {
		nums[left] = nums[right]
		for right+1 < len(nums) && nums[right+1] == nums[left] {
			right++
		}
		left++
		right++
	}
	for i := 0; i < left; i++ {
		fmt.Println(nums[i])
	}
	return left
}

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
