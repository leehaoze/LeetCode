package main

import "sort"

func thirdMax(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	count := 1
	pre := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == pre {
			continue
		}

		count += 1
		pre = nums[i]

		if count == 3 {
			return nums[i]
		}
	}

	return nums[0]

}

func main() {

}
