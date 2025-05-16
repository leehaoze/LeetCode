package main

import "fmt"

/*
	[2,3,1,1,4]

	[2] 可以
	[2,3] nums[0] >= len(nums) 可以
	[2,3,1] nums[0] >= len(nums) 可以
	[2,3,1,1]  dp[n-1] >= len(nums) - n

	dp[i][j] = max(dp[0..i-1][j], nums[i]-1)
*/

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

func canJump(nums []int) bool {
	if nums[0] >= len(nums)-1 {
		return true
	}

	dp1 := make([]int, len(nums))
	dp2 := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if nums[0] >= i && i != 0 {
			dp1[i] = nums[0] - i
		} else {
			dp1[i] = -1
		}
	}
	// fmt.Println(dp1)
	for i := 1; i < len(nums); i++ {
		if dp1[i] < 0 {
			return false
		}

		for j := i; j < len(nums); j++ {
			if j == i {
				dp2[j] = dp1[i]
			} else {
				dp2[j] = max(dp1[j], nums[i] - j + i)
			}
		}

		// fmt.Println(fmt.Sprintf("%d %v", i, dp2))
		dp1, dp2 = dp2, dp1
	}


	return dp2[len(nums)-1] >= 0
}

// func canJump(nums []int) bool {
// 	if nums[0] >= len(nums) - 1{
// 		return true
// 	}
// 	dp := make([][]int, len(nums))
// 	for idx := range dp {
// 		dp[idx] = make([]int, len(nums))
// 		for i := range dp[idx] {
// 			dp[idx][i] = -1
// 		}
// 	}
//
// 	for i := 0; i < len(nums); i++ {
// 		for j := i; j < len(nums); j++ {
// 			if i > 0 {
// 				if dp[i-1][i] < 0 {
// 					continue
// 				}
// 				if i == j {
// 					dp[i][j] = dp[i-1][j]
// 				} else {
// 					dp[i][j] = max(dp[i-1][j], nums[i]-j+i)
// 				}
// 			} else {
// 				if i == j {
// 					continue
// 				} else {
// 					dp[i][j] = max(dp[i][j], nums[i]-j+i)
// 				}
// 			}
// 		}
// 	}
// 	// fmt.Println(dp)
// 	return dp[len(nums) - 1][len(nums) - 1] >= 0
// }

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump([]int{0}))
	fmt.Println(canJump([]int{1}))
	fmt.Println(canJump([]int{1, 2, 3}))
	fmt.Println(canJump([]int{0, 2, 3}))
}
