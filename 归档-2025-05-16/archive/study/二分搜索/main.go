package main

import "fmt"

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		}
	}

	return -1
}

func binarySearchLeftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] >= target {
			right = mid - 1
		}
	}

	if left == len(nums) {
		return -1
	} else {
		return left
	}
}

func binarySearchRightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] <= target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}

	if right == -1 {
		return -1
	} else {
		return right
	}
}

func main() {
	data := []int{1, 3, 5, 5, 6, 9}
	fmt.Println(binarySearchRightBound(data, 5))
	fmt.Println(binarySearchLeftBound(data, 5))
}
