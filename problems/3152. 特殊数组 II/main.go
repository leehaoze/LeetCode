package main

import "fmt"

func isArraySpecial(nums []int, queries [][]int) []bool {
	points := make([]int, len(nums))
	// 计算断点数组
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]%2 == nums[i+1]%2 {
			// 同时计算前缀和
			points[i+1] = points[i] + 1
		} else {
			points[i+1] = points[i]
		}
	}

	// 求解
	ret := make([]bool, len(queries))

	for i, query := range queries {
		if points[query[1]]-points[query[0]] == 0 {
			ret[i] = true
		}
	}

	return ret
}

func main() {
	fmt.Println(isArraySpecial([]int{4, 3, 1, 6}, [][]int{{2, 3}}))
}
