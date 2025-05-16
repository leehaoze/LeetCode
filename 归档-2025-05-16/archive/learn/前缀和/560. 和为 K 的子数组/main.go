package main

import "fmt"

func subarraySum(nums []int, k int) int {
	preSum := make([]int, len(nums)+1)
	valMap := make(map[int]int)
	valMap[0] = 1
	count := 0
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
		count += valMap[preSum[i+1]-k]
		valMap[preSum[i+1]]++
	}

	// count := 0
	// for i := 1; i < len(preSum); i++ {
	// 	for j := i - 1; j >= 0; j-- {
	// 		if preSum[i]-preSum[j] == k {
	// 			count++
	// 		}
	// 	}
	// }

	return count
}

func main() {
	// [0, 1, 2, 3]
	fmt.Println(subarraySum([]int{1}, 0))
}
