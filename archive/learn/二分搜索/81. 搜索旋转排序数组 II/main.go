package main

import "fmt"

// 1,1,1,2 => 1,2,1,1
func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return true
		}

		// 判断落入哪个区间
		if nums[left] < nums[mid] { // mid 落入左侧区间
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[left] == nums[mid] { // 无法判断
			left++
		} else if nums[mid] < nums[len(nums)-1] {
			if target >= nums[mid] && target <= nums[len(nums)-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else if nums[mid] == nums[right] {
			right--
		}

	}

	return false
}

func main() {
	// fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	// fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
	// fmt.Println(search([]int{1, 0, 1, 1, 1}, 0))
	fmt.Println(search([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 13, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 13))
}
