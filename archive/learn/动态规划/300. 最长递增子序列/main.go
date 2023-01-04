package main

import "fmt"

// func lengthOfLIS(nums []int) int {
// 	ret := 0
// 	for i := int64(0); i < 1<<len(nums); i++ {
// 		preValue := 0
// 		first := true
// 		yes := true
// 		count := 0
// 		for j := 0; j < len(nums); j++ {
// 			if i>>j&1 == 1 {
// 				if first {
// 					preValue = nums[j]
// 					first = false
// 				} else {
// 					if nums[j] <= preValue {
// 						yes = false
// 						break
// 					} else {
// 						preValue = nums[j]
// 					}
// 				}
// 				count++
// 			}
// 		}
//
// 		if yes {
// 			// fmt.Println(fmt.Sprintf("%b", i))
// 			if count > ret {
// 				ret = count
// 			}
// 		}
// 	}
// 	return ret
// }

// func lengthOfLIS(nums []int) int {
// 	l, _ := solve(nums)
// 	return l
// }
//
// func solve(nums []int) (int, int) {
// 	if len(nums) == 1 {
// 		return 1, nums[0]
// 	}
//
// 	l, minVal := solve(nums[1:])
// 	if nums[0] < minVal {
// 		return l + 1, nums[0]
// 	}
// 	return l, minVal
// }
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	ret := 0
	for i := 0; i < len(dp); i++ {
		if dp[i] > ret {
			ret = dp[i]
		}
	}

	return ret
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7}))
}
