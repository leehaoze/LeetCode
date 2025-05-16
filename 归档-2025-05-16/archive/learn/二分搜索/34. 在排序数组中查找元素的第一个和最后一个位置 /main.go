package main

import "fmt"

func searchFirstPos(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] >= target {
			right = mid - 1
		}
	}

	if left == len(nums) || nums[left] != target {
		return -1
	}

	return left
}

func searchLastPos(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] <= target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}

	if right == -1 || nums[right] != target {
		return -1
	}

	return right
}

func searchRange(nums []int, target int) []int {
	return []int{searchFirstPos(nums, target), searchLastPos(nums, target)}
}

func main() {
	fmt.Println(searchRange([]int{}, 6))
}
