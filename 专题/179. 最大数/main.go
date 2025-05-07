package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		a := nums[i]
		b := nums[j]
		return compare(a, b)
	})

	fmt.Println(nums)
	ret := ""
	for _, num := range nums {
		ret += fmt.Sprintf("%d", num)
	}

	return ret
}

// compare 如果 a > b，返回true，否则返回false
func compare(a, b int) bool {
	
}

func main() {
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
}
