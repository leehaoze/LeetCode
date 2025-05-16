package main

import (
	"fmt"
)

func arrayNesting(nums []int) int {
	n := len(nums)
	ret := 0

	for i := 0; i < n; i++ {
		length := 0
		pos := i
		m := map[int]struct{}{}
		for pos < n {
			_, exists := m[nums[pos]]
			if !exists {
				length++
				m[nums[pos]] = struct{}{}
				pos = nums[pos]
			} else {
				break
			}
		}

		if ret < length {
			ret = length
		}
	}

	return ret
}

func main() {
	fmt.Println(arrayNesting([]int{5, 4, 0, 3, 1, 6, 2}))
}
