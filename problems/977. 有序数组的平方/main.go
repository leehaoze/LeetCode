package main

import (
	"fmt"
)

func sortedSquares(nums []int) []int {
	left, right := 0, len(nums)-1
	ret := make([]int, len(nums))
	idx := len(nums) - 1
	for left <= right {
		l, r := nums[left]*nums[left], nums[right]*nums[right]
		if l > r {
			ret[idx] = l
			left++
		} else {
			ret[idx] = r
			right--
		}
		idx--
	}

	return ret
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}
