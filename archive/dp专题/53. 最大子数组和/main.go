package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func maxSubArray(nums []int) int {
// 	dp := make([]int, len(nums))
// 	for idx, num := range nums {
// 		if idx == 0 {
// 			dp[idx] = num
// 		} else {
// 			dp[idx] = max(dp[idx-1]+num, num)
// 		}
// 	}
//
// 	max := dp[0]
// 	for _, i := range dp {
// 		if max < i {
// 			max = i
// 		}
// 	}
// 	return max
// }

func maxSubArray(nums []int) int {
	dpPre := 0
	ret := nums[0]

	for _, num := range nums {
		dpPre = max(dpPre+num, num)
		ret = max(ret, dpPre)
	}
	return ret
}
