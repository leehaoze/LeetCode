package main

import (
	"fmt"
)

// sortedSquares 将数组每一项进行平方，返回平方后的值组成的非递减数组
func sortedSquares(nums []int) []int {
	ret := make([]int, len(nums))
	idx := len(ret) - 1

	// 准备两个指针，从两侧开始扫描，平方后值比较大的移动到新数组，直到两个指针相遇
	left, right := 0, len(nums)-1

	// 区间是一个左闭右闭的区间，left == right 时，仍有一个值需要处理
	for left <= right {
		// 左侧更大
		if nums[left]*nums[left] > nums[right]*nums[right] {
			ret[idx] = nums[left] * nums[left]
			left++
		} else {
			ret[idx] = nums[right] * nums[right]
			right--
		}
		idx--
	}

	return ret
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}
