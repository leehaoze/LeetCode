package main

import "fmt"

// ------------- 解1 -------------
// func search(nums []int, target int) int {
// 	if len(nums) == 1 {
// 		if nums[0] == target {
// 			return 0
// 		} else {
// 			return -1
// 		}
// 	}
//
// 	// step 1 找到旋转点
// 	rotatePos := 0
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] < nums[i-1] {
// 			rotatePos = i
// 		}
// 	}
//
// 	// 判断target在哪个部分
// 	// 4,5,6,1,2,3
// 	if rotatePos > 0 {
// 		if nums[0] > target { // 右半部分
// 			return binarySearch(nums, rotatePos, len(nums)-1, target)
// 		} else { // 左半部分
// 			return binarySearch(nums, 0, rotatePos-1, target)
// 		}
// 	} else {
// 		return binarySearch(nums, 0, len(nums)-1, target)
// 	}
//
// }
//
// func binarySearch(nums []int, start, end, target int) int {
// 	left, right := start, end
// 	for left <= right {
// 		mid := left + (right-left)/2
// 		if nums[mid] == target {
// 			return mid
// 		} else if nums[mid] > target {
// 			right = mid - 1
// 		} else if nums[mid] < target {
// 			left = mid + 1
// 		}
// 	}
//
// 	return -1
// }

// ------------- step 解2 错误 -------------
// func search(nums []int, target int) int {
// 	left, right := 0, len(nums)-1
// 	for left <= right {
// 		mid := left + (right - left) / 2
// 		if nums[left] < nums[right] { // 区间是严格递增的 继续查找
// 			if nums[mid] == target {
// 				return mid
// 			} else if nums[mid] < target {
// 				left = mid + 1
// 			} else if nums[mid] > target {
// 				right = mid - 1
// 			}
// 		} else { // 区间不是严格递增的，要根据mid情况判断
// 			if nums[mid] == target {
// 				return mid
// 			} else if nums[mid] < target {
// 				right =
// 			}
// 		}
// 	}
// }

// ------------- step 解2 -------------
// func search(nums []int, target int) int {
// 	if len(nums) == 1 {
// 		if nums[0] == target {
// 			return 0
// 		} else {
// 			return -1
// 		}
// 	}
//
// 	left, right := 0, len(nums)-1
// 	for left <= right {
// 		mid := left + (right-left)/2
// 		if nums[mid] == target {
// 			return mid
// 		}
// 		if nums[0] <= nums[mid] { // 左侧有序
// 			if nums[0] <= target && nums[mid] > target {
// 				right = mid - 1
// 			} else {
// 				left = mid + 1
// 			}
// 		} else { // 右侧有序
// 			if nums[mid] < target && nums[len(nums)-1] >= target {
// 				left = mid + 1
// 			} else {
// 				right = mid - 1
// 			}
// 		}
// 	}
//
// 	return -1
// }

// 2022-03-03 再一次做
// 4,5,6,7,0,1,2
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] { // 左侧有序 mid落入左半部分
			if nums[mid] < target || nums[0] > target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else { // 右侧有序 mid落入右半部分
			if nums[mid] < target && nums[len(nums)-1] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 1, 2, 3}, 1))
	fmt.Println(search([]int{4, 5, 6, 1, 2, 3}, 6))
	fmt.Println(search([]int{1, 3}, 5))
	fmt.Println(search([]int{3, 1}, 1))
}
