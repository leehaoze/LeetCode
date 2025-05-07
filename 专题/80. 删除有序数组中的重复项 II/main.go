package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	newIdx, scanIdx := 0, 0

	for scanIdx < len(nums) {
		if newIdx < 2 || nums[newIdx-2] != nums[scanIdx] {
			nums[newIdx] = nums[scanIdx]
			newIdx++
		}

		scanIdx++
	}

	return newIdx
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(nums[:removeDuplicates(nums)])
}
