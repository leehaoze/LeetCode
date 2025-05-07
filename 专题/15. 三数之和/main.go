package main

import (
	"fmt"
	"slices"
)

func threeSum(nums []int) [][]int {
	var ret [][]int
	slices.Sort(nums)

	// 第一层循环，寻找第一个数
	for i := 0; i < len(nums); i++ {
		// 跳过重复的第一个数
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 第二层和第三层循环使用双指针实现
		left, right := i+1, len(nums)-1

		// 两个位置不能相等，相等会少一个数
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				ret = append(ret, []int{nums[i], nums[left], nums[right]})
				// 跳过重复的第二个数
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// 跳过重复的第三个数
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return ret
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{1, -1, -1, 0}))
	// [-1,-1,0,1]
}
