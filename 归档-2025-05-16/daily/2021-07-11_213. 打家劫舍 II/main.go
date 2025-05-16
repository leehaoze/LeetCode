package main

import "fmt"

/**
n = 1 直接偷
n = 2 选一个贵的偷就可以
...
n  if !stolenFirst[n-2] max(dp[n-2] + value[b], dp[n-1])
	else max(dp[n-2], dp[n-1])
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

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	dp0, dp1 := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums) - 1; i++ {
		dp0, dp1 = dp1, max(dp0+nums[i], dp1)
	}

	dp2, dp3 := nums[1], max(nums[1], nums[2])
	for i := 3; i < len(nums); i++ {
		dp2, dp3 = dp3, max(dp2+nums[i], dp3)
	}

	return max(dp1, dp3)
}

func main() {
	fmt.Println(rob([]int{2, 3, 2}))
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2,2,4,3,2,5}))
}
