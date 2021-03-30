package main

import (
	"fmt"
	"sort"
)

var memory map[string]byte

func subsetsWithDup(nums []int) [][]int {
	memory = make(map[string]byte)
	return helper(nums, 0, len(nums))
}

func helper(nums []int, start, end int) [][]int {
	if start == end {
		return [][]int{make([]int, 0)}
	}

	subResult := helper(nums, start+1, end)
	newResult := make([][]int, 0, len(subResult))
	for _, childList := range subResult {
		newResult = append(newResult, childList)
		childList = append(childList, nums[start])
		sort.Ints(childList)
		//_, e := memory[fmt.Sprintf("%v", childList)]
		//fmt.Printf("%v, exists: %v\n", childList, e)
		if _, exists := memory[fmt.Sprintf("%v", childList)]; !exists {
			temp := make([]int, len(childList))
			copy(temp, childList)
			newResult = append(newResult, temp)
			memory[fmt.Sprintf("%v", childList)] = 1
		}

		//for _, arr := range newResult {
		//	fmt.Printf(" %v{@%p} ", arr, arr)
		//}
		//
		//fmt.Println()

	}
	return newResult
}

func main() {
	result := subsetsWithDup([]int{1, 1, 2, 2})
	//fmt.Println("------")
	for _, list := range result {
		fmt.Println(fmt.Sprintf("%v", list))
	}
}
