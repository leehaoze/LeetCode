package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	// 左闭右闭区间
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		// 根据 nums[mid] 和 nums[0]的关系，可知mid在左侧区间，还是右侧区间
		// 再根据 target 和 nums[0]的关系，可以知道我们应该如何收缩区间
		if nums[mid] == target {
			return mid
		} else if nums[left] <= nums[mid] {
			// left 到 mid 是递增的区间，如果target在这个区间，则收缩到这个区间
			// left 以外的已经排除了
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// mid 到 right 是递增的区间，如果target在这个区间，则收缩到这个区间
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(search([]int{3, 1}, 1))
}
