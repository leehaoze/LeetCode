package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
	}

	ret := nums[0]
	for i := 0; i < len(dp); i++ {
		if dp[i] > ret {
			ret = dp[i]
		}
	}

	return ret
}
