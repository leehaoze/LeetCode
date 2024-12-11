package main

import (
	"fmt"
)

// search 通过二分法在数组中查找指定值，返回目标值的下标，如果目标值不存在，则返回-1
// 当数组中有多个目标值时，返回下标最大的那一个
func search(nums []int, target int) int {
	// left, right 是搜索区间的边界值，搜索区间是左闭右闭
	left, right := 0, len(nums)-1

	// 当 left == right 时，搜索区间仍有一个值，需要继续对比
	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			// 当找到目标值时，需要继续缩小搜索区间
			// 由于我们要返回下标最大的位置（区间右侧的某个值），因此需要收缩区间的左侧
			// mid 值已经对比过了，需要排除在外
			left = mid + 1
		} else if nums[mid] > target {
			// target 在左侧，需要缩小右侧边界
			// 由于搜索区间是闭合的，mid已经对比过了，因此需要排除到区间外
			right = mid - 1
		} else {
			// 同上
			left = mid + 1
		}
	}

	// 结束循环时，由于 nums[mid] == target 时，left会+1 导致搜索区间结束，要找的解在 left -1的位置

	// 当没有目标值时，left有可能会越界
	if left-1 >= 0 && left-1 < len(nums) && nums[left-1] == target {
		return left - 1
	}

	return -1
}

func main() {
	fmt.Println(search([]int{0, 1, 3, 9, 12}, 0))
}
