package main

import "fmt"

func hIndex(citations []int) int {
	left, right := 0, len(citations)
	mid := 0
	for left < right {
		// fmt.Println(left, right)
		mid = (left + right + 1) >> 1
		if citations[len(citations) - mid] >= mid {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return right
}

func check(citations []int, h int) bool {
	return citations[len(citations) - h] >= h
}

func main() {
	fmt.Println(hIndex([]int{0,1,3,5,6}))
	fmt.Println(hIndex([]int{1}))
	fmt.Println(hIndex([]int{0}))
}