package main

import (
	"fmt"
)

// jump 返回到达最后一个位置的最小跳跃次数
func jump(nums []int) int {
	ret := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		// 可以向跳跃val的距离
		val := nums[i]
		for j := 1; j <= val && i+j < len(nums); j++ {
			if ret[i+j] == 0 {
				ret[i+j] = ret[i] + 1
			} else {
				if ret[i]+1 < ret[i+j] {
					ret[i+j] = ret[i] + 1
				}
			}
		}
	}

	// fmt.Println(ret)

	return ret[len(ret)-1]
}

func main() {
	// fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{2, 1}))
}
