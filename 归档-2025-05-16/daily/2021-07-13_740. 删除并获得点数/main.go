package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func deleteAndEarn(nums []int) int {
	maxVal := 0
	for _, val := range nums {
		maxVal = max(maxVal, val)
	}
	sum := make([]int, maxVal+1)
	for _, val := range nums {
		sum[val] += val
	}

	return rob(sum)
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dp0, dp1 := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp0, dp1 = dp1, max(dp0+nums[i], dp1)
	}

	return dp1
}

func main() {
	// fmt.Println(deleteAndEarn([]int{3, 4, 2}))
	// fmt.Println(deleteAndEarn([]int{2,2,3,3,3,4}))
	// fmt.Println(deleteAndEarn([]int{12, 32, 93, 17, 100, 72, 40, 71, 37, 92, 58, 34, 29, 78, 11, 84, 77, 90, 92, 35, 12, 5, 27, 92, 91, 23, 65, 91, 85, 14, 42, 28, 80, 85, 38, 71, 62, 82, 66, 3, 33, 33, 55, 60, 48, 78, 63, 11, 20, 51, 78, 42, 37, 21, 100, 13, 60, 57, 91, 53, 49, 15, 45, 19, 51, 2, 96, 22, 32, 2, 46, 62, 58, 11, 29, 6, 74, 38, 70, 97, 4, 22, 76, 19, 1, 90, 63, 55, 64, 44, 90, 51, 36, 16, 65, 95, 64, 59, 53, 93}))
	fmt.Println(deleteAndEarn([]int{3, 1}))
}
