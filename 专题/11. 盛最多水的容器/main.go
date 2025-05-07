package main

import (
	"fmt"
)

// maxArea 数组中任意两项构成边界，可以存储水，此函数返回可以存储的最大水量
func maxArea(height []int) int {
	// left, right 分别是左右边界
	left, right := 0, len(height)-1
	// ret 最终解
	var ret int

	// 边界不能重叠
	for left < right {
		// 计算可以容纳的水量
		water := (right - left) * min(height[right], height[left])

		if water > ret {
			ret = water
		}

		// 移动较短边
		if height[right] > height[left] {
			left++
		} else {
			right--
		}
	}

	return ret
}

// min 返回较小值，不考虑相等的情况
func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
