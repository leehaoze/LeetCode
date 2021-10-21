package main

import "fmt"

func minMoves(nums []int) int {
	min := nums[0]
	for i := range nums {
		if nums[i] < min {
			min = nums[i]
		}
	}

	ret := 0
	for i := range nums {
		ret += nums[i] - min
	}

	return ret
}

func main() {
	fmt.Println(minMoves([]int{2,3,4}))
}
