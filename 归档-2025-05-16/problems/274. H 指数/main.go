package main

import "fmt"

func hIndex(citations []int) int {
	// 答案空间是 [0, 论文篇数]
	left, right := 0, len(citations)
	ret := 0

	// 左闭右闭区间
	for left <= right {
		mid := (right-left)/2 + left

		// mid是否满足k指数
		// 有 > mid篇论文，被引用了 > mid 次
		isK := func(mid int, citations []int) bool {
			count := 0
			for _, val := range citations {
				if val >= mid {
					count++

					if count >= mid {
						return true
					}
				}
			}

			return false
		}

		if isK(mid, citations) {
			if mid >= ret {
				ret = mid
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ret
}

func main() {
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
}
