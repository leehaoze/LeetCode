package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	// 6 5 3 1 0
	result := 0
	for idx, value := range citations {
		if idx+1 <= value {
			if idx+1 > result {
				result = idx + 1
			}
		} else {
			if value > result {
				result = value
			}
		}
	}

	return result
}

func main() {
	// result := hIndex([]int{3, 0, 6, 1, 5})
	result := hIndex([]int{1,3,1})
	// result := hIndex([]int{11, 15})
	fmt.Println(result)
}
