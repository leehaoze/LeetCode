package main

import (
	"fmt"
	"sort"
)

func smallestK(arr []int, k int) []int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	return arr[:k]
}

func main() {
	fmt.Println(smallestK([]int{1,3,5,7,2,4,6,8}, 4))
}
