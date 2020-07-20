package main

import "fmt"

func searchInsert(nums []int, target int) int {
	// 不就是二分么。。。。
	left, right := 0, len(nums) - 1
	for left < right {
		middle := (left + right) / 2

		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			right = middle - 1
		} else if nums[middle] < target {
			left = middle + 1
		}
	}

	if nums[left] == target {
		return left
	} else if nums[left] < target {
		return left + 1
	} else {
		return left
	}
}

func main() {
	fmt.Printf("result %d\n", searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Printf("result %d\n", searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Printf("result %d\n", searchInsert([]int{1, 3, 5, 6}, 7))
	fmt.Printf("result %d\n", searchInsert([]int{1, 3, 5, 6}, 0))
	fmt.Printf("result %d\n", searchInsert([]int{1, 3, 5, 6}, 4))
}
