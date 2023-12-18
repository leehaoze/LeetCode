package main

func findPeakElement(nums []int) int {

	// 左闭右开的搜索区间
	left, right := 0, len(nums)

	// 搜索区间为空的情况：left == right
	for left < right {
		mid := left + (right-left)/2

		// 判断mid的位置，上升区间，下降区间，峰值
		if (mid-1 < 0 || nums[mid-1] < nums[mid]) &&
			(mid+1 >= len(nums) || nums[mid] > nums[mid+1]) {
			return mid
		} else if (mid-1 < 0 || nums[mid-1] < nums[mid]) &&
			(mid+1 >= len(nums) || nums[mid] < nums[mid+1]) {
			// 上升区间 mid > mid -1 并且 mid < mid + 1
			left = mid + 1
		} else if (mid-1 < 0 || nums[mid-1] > nums[mid]) &&
			(mid+1 >= len(nums) || nums[mid] > nums[mid+1]) {
			// 下降区间 mid < mid-1 并且 mid > mid + 1
			right = mid
		} else {
			// 山谷
			right = mid
		}
	}

	return -1
}
