package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scanln(&n)

	nums := make([]int, n)
	preSum := make([]int, n+1)

	for i := 0; i < n; i++ {
		_, _ = fmt.Scanln(&nums[i])
		preSum[i+1] = nums[i] + preSum[i]
	}

	// fmt.Println(preSum)

	for {
		var start, end int
		_, err := fmt.Scanf("%d %d\n", &start, &end)
		if err != nil {
			break
		}

		fmt.Println(preSum[end+1] - preSum[start])
	}
}
