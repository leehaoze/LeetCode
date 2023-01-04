package main

import "fmt"

// func minSubArrayLen(target int, nums []int) int {
// 	preSum := make([]int, len(nums)+1)
// 	for i, num := range nums {
// 		preSum[i+1] = preSum[i] + num
// 	}
//
// 	ret := len(nums) + 1
// 	for i := 0; i < len(preSum); i++ {
// 		for j := i + 1; j < len(preSum); j++ {
// 			if preSum[j]-preSum[i] >= target {
// 				if j-i < ret {
// 					ret = j - i
// 				}
// 			}
// 		}
// 	}
//
// 	if ret == len(nums)+1 {
// 		return 0
// 	}
// 	return ret
// }

// func minSubArrayLen(target int, nums []int) int {
// 	minLength := len(nums) + 1
// 	valMap := make(map[int]int) // 这个val就是下标了
// 	valMap[0] = 0               // 这个value里的0也是下标
// 	preSum := 0
// 	for i := 0; i < len(nums); i++ {
// 		preSum = preSum + nums[i]
// 		if pos, exists := valMap[preSum-target]; exists {
// 			if minLength > i-pos {
// 				minLength = i - pos
// 			}
// 		}
// 		valMap[preSum] = i
// 	}
//
// 	if minLength == len(nums)+1 {
// 		return 0
// 	}
//
// 	return minLength
// }

// func binarySearch(target int, nums []int, start int) int {
// 	left := start
// 	right := len(nums) - 1
// 	// 在 [start, end] 这个区间中寻找一个大于等于target的数，并且尽可能小的数
// 	idx := -1
// 	for left <= right {
// 		mid := left + (right-left)>>1
// 		if nums[mid] < target {
// 			left = mid + 1
// 		} else if nums[mid] >= target {
// 			right = mid - 1
// 			idx = mid
// 		}
// 	}
//
// 	return idx
// }
//
// func minSubArrayLen(target int, nums []int) int {
// 	preSum := make([]int, len(nums)+1)
// 	for i := 0; i < len(nums); i++ {
// 		preSum[i+1] = preSum[i] + nums[i]
// 	}
//
// 	ret := len(nums) + 1
// 	for i := 0; i < len(preSum); i++ {
// 		n := preSum[i] + target
// 		// 搜索这个n是否在 [i+1, len(preSum)] 之间存在
// 		if idx := binarySearch(n, preSum, i+1); idx != -1 {
// 			if ret > idx-i {
// 				ret = idx - i
// 			}
// 		}
// 	}
//
// 	if ret == len(nums)+1 {
// 		return 0
// 	}
//
// 	return ret
// }

func minSubArrayLen(target int, nums []int) int {
	start, end := 0, 0
	sum := 0
	ret := len(nums) + 1
	for end < len(nums) {
		sum += nums[end]
		for sum >= target {
			if ret > end-start+1 {
				ret = end - start + 1
			}

			sum -= nums[start]
			start++
		}
		end++
	}

	if ret == len(nums)+1 {
		return 0
	}

	return ret
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(11, []int{1, 2, 3, 4, 5}))
	// fmt.Println(binarySearch(6, []int{1, 2, 3, 4, 5}, 4))
}
