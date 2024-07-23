package main

import "fmt"

// 0, 1, 1, 2, 3
func removeDuplicates(nums []int) int {
	left, right := 0, 0

	for right < len(nums) {
		if left-1 < 0 || nums[left-1] != nums[right] {
			nums[left] = nums[right]
			left++
		}

		right++
		// fmt.Println(nums)
	}

	// fmt.Println(nums)
	return left
}

func main() {
	ret := removeDuplicates([]int{1, 1, 2})
	fmt.Println(ret)
}
