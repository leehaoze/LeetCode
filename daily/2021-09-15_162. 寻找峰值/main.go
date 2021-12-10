package main

import "fmt"

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return 0
		} else {
			return 1
		}
	}

	if nums[0] > nums[1] {
		return 0
	}

	for i := 2; i < len(nums); i++ {
		if nums[i-1] > nums[i] && nums[i-1] > nums[i-2] {
			return i - 1
		}
	}

	if nums[len(nums)-1] > nums[len(nums)-2] {
		return len(nums) - 1
	}

	return -1
}

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
	fmt.Println(findPeakElement([]int{3, 2, 1}))
}
