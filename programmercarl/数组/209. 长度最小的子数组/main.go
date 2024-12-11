package main

import (
	"fmt"
)

// minSubArrayLen 返回 长度最小的，且和大于target的子数组 的长度
func minSubArrayLen(target int, nums []int) int {
	// 计算前缀和
	preSum := make([]int, len(nums)+1)
	for idx, val := range nums {
		if idx == 0 {
			preSum[idx] = val
			continue
		}

		preSum[idx+1] = preSum[idx-1] + val
	}

	var ret int

	// 可以通过二分，查找 以 i 开始，满足 > target 条件的最短长度子数组
	for i := 0; i < len(nums); i++ {
		idx := binarySearchLeft(preSum, i, target-preSum[i])
		if idx != -1 {
			// 计算长度并更新
			length := idx - i + 1
			if ret == 0 || ret > length {
				ret = length
			}
		}
	}

	return ret
}

// binarySearchLeft 通过二分查找 大于等于 target 的第一个元素
func binarySearchLeft(nums []int, start, target int) int {
	// 左闭右闭区间
	left, right := start, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] >= target {
			// 向左收缩
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if left < len(nums) && nums[left] >= target {
		return left
	}

	return -1
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}
