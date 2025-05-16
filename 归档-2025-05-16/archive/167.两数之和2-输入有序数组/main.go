package main

import (
	"fmt"
)

func binary_search(numbers []int, start int, target int) int {
	var ret = -1

	left, right := start, len(numbers)-1

	for left <= right {
		mid := left + (right-left)/2
		if numbers[mid] == target {
			ret = mid
			break
		} else if numbers[mid] < target {
			left = mid + 1
		} else if numbers[mid] > target {
			right = mid - 1
		}
	}

	return ret
}

func twoSum(numbers []int, target int) []int {
	for idx, item := range numbers {
		if targetPos := binary_search(numbers, idx+1, target-item); targetPos != -1 {
			return []int{idx + 1, targetPos + 1}
		}
	}

	return []int{}
}

func main() {
	fmt.Printf("%v", twoSum([]int{2, 7, 11, 15}, 13))
}
