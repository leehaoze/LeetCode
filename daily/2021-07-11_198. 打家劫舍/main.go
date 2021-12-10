package main

import "fmt"

/**
dp[i] = max(dp[i-2] + value[i], dp[i-1])
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
	fmt.Println(rob([]int{1,2,3,1}))
	fmt.Println(rob([]int{2,7,9,3,1}))
}