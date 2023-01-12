package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	rotatePoint := nums[0]
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		inLeft := nums[mid] >= rotatePoint

		if (inLeft && nums[mid] > target) || (!inLeft && nums[mid] < target) {
			if target >= rotatePoint {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] > target {
				right = mid - 1
			} else if nums[mid] < target {
				left = mid + 1
			}
		}

	}

	return -1
}

func main() {
	fmt.Println(search([]int{1, 3}, 3))
}
