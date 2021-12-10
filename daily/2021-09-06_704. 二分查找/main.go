package main

import "fmt"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + ((right - left) >> 2)
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func main() {
	fmt.Println(search([]int{1, 2, 3, 4, 5}, 3))
	fmt.Println(search([]int{1, 2, 3, 4, 5}, 4))
	fmt.Println(search([]int{1, 2, 3, 4, 5}, 5))
	fmt.Println(search([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(search([]int{1, 2, 3, 4, 5}, -1))

}
