package main

import (
	"fmt"
	"slices"
)

func findValueOfPartition(nums []int) int {
	var ret = 0

	slices.Sort(nums)

	for i := 1; i < len(nums); i++ {
		if i == 1 {
			ret = nums[i] - nums[i-1]
		} else {
			if nums[i]-nums[i-1] < ret {
				ret = nums[i] - nums[i-1]
			}
		}
	}

	return ret
}

func main() {
	fmt.Println(findValueOfPartition([]int{4, 1, 4, 5, 1, 5}))
}
