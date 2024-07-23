package main

import "fmt"

// func removeElement(nums []int, val int) int {
// 	left, right := 0, len(nums)-1

// 	// left == right的情况可以不用考虑
// 	for left < right {
// 		// 尾指针移动到非target值的位置
// 		for right >= 0 && nums[right] == val {
// 			right--
// 		}

// 		// 头指针移动到target值的位置
// 		for left < len(nums) && nums[left] != val {
// 			left++
// 		}

// 		// 交换
// 		if left < right {
// 			nums[left], nums[right] = nums[right], nums[left]
// 			left++
// 			right--
// 		}
// 	}

// 	fmt.Println(left, right)

// 	if left > right {
// 		return left
// 	}

// 	if nums[left] == val {
// 		return left
// 	}

// 	return left + 1
// }

func removeElement(nums []int, val int) int {
	left, right := 0, 0

	for right < len(nums) {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}

		right++
	}

	return left
}

func main() {
	nums := []int{2, 2, 3}
	k := removeElement(nums, 2)
	fmt.Println(fmt.Sprintf("%#v, %#v", nums, k))
}
