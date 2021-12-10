package main

import (
	"fmt"
	"sort"
)

func minPairSum(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	length := len(nums) - 1
	ret, sum := 0, 0
	for i := 0; i < length / 2 + 1; i++ {
		sum = nums[i] + nums[length - i]
		if sum > ret {
			ret = sum
		}
	}

	return ret
}

func main() {
	fmt.Println(minPairSum([]int{3,5,2,3}))
	fmt.Println(minPairSum([]int{3,5,4,2,4,6}))
	fmt.Println(minPairSum([]int{1,1}))
	fmt.Println(minPairSum([]int{4,5,5,4,3,3,324,5,5,35,366}))
}