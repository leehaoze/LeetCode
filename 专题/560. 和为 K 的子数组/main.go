package main

import (
	"fmt"
)

// subarraySum 返回nums数组中，和为k的子数组数量
func subarraySum(nums []int, k int) int {
	// count是最终结果，pre是前缀和
	count, pre := 0, 0
	m := map[int]int{}
	// 前缀和为0的有一个子数组
	m[0] = 1

	for i := 0; i < len(nums); i++ {
		pre += nums[i]

		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		m[pre] += 1
	}

	return count
}

func main() {
	fmt.Println(subarraySum([]int{-1, -1, 1}, 0))
}
