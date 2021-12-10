package main

import "fmt"

func subarraySum(nums []int, k int) int {
	preSum := make([]int, len(nums)+1)
	valMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
		valMap[preSum[i+1]]++
	}
	valMap[0] = 1

	count := 0
	for i := 0; i < len(preSum); i++ {
		if _, exists := valMap[preSum[i]+k]; exists {
			count++
		}
		// for j := i + 1; j < len(preSum); j++ {
		// 	if preSum[j]-preSum[i] == k {
		// 		count++
		// 	}
		// }
	}

	return count
}

func main() {
	// [0, 1, 2, 3]
	fmt.Println(subarraySum([]int{1}, 0))
}
