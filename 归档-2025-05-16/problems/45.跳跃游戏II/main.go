package main

import "fmt"

/*
给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。

每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:

0 <= j <= nums[i]
i + j < n
返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。
*/

func jump(nums []int) int {
	// dp[i] 到达此节点的最小跳跃次数，dp[0] = 0
	dp := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		dp[i] = len(nums) + 1
	}

	for i := 0; i < len(nums); i++ {

		for j := 1; j <= nums[i]; j++ {
			// fmt.Println(fmt.Sprintf("i: %d, j: %d", i, j))
			if i+j >= len(nums) {
				break
			}
			dp[i+j] = min(dp[i+j], dp[i]+1)
		}
	}

	// fmt.Println(dp)
	return dp[len(nums)-1]
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}
