package main

import (
	"fmt"
)

func minSubArrayLen(target int, nums []int) int {
	ret := len(nums) + 1
	start, end := 0, 0
	sum := 0
	for end < len(nums) {
		sum += nums[end]
		for sum >= target {
			if ret > end-start+1 {
				ret = end - start + 1
			}
			if ret == 1 {
				return 1
			}
			sum -= nums[start]
			start++
		}
		end++
	}

	if ret == len(nums)+1 {
		ret = 0
	}

	return ret
}

func main() {
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))
}
