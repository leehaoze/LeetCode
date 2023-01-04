package main

func runningSum(nums []int) []int {
	preSum := make([]int, len(nums))
	preSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i]
	}
	return preSum
}
