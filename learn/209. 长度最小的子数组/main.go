package main

func minSubArrayLen(target int, nums []int) int {
	preSum := make([]int, len(nums)+1)
	for i, num := range nums {
		preSum[i+1] = preSum[i] + num
	}

	ret := len(nums) + 1
	for i := 0; i < len(preSum); i++ {
		for j := i + 1; j < len(preSum); j++ {
			if preSum[j]-preSum[i] >= target {
				if j-i < ret {
					ret = j - i
				}
			}
		}
	}

	if ret == len(nums)+1 {
		return 0
	}
	return ret
}
