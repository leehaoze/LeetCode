package main

import "fmt"

func majorityElement(nums []int) int {
	ret, count := 0, 0
	for _, val := range nums {
		if count == 0 || val == ret {
			ret = val
			count++
		}

		if val != ret {
			count--
		}
	}

	return ret
}

// [2,2,1,1,1,2,2]

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
}
