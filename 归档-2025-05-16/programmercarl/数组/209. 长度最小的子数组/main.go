package main

import (
	"fmt"
)

// minSubArrayLen 返回 长度最小的，且和大于target的子数组 的长度
func minSubArrayLen(target int, nums []int) int {
	// 左开右闭区间
	left, right := 0, 0
	sum := 0
	ret := len(nums) + 1
	for right < len(nums) {
		right++
		sum += nums[right-1]

		for sum >= target {
			ret = min(ret, right-left)
			sum -= nums[left]
			left++
		}
	}

	if ret == len(nums)+1 {
		return 0
	}

	return ret
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}
