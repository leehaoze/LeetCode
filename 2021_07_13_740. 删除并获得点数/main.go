package main

import "fmt"

var mem = make(map[string]int)

func deleteAndEarn(nums []int) int {
	fmt.Println(nums)
	if len(nums) <= 0 {
		return 0
	}
	sub, earn := delete(nums, nums[0])

	a, b := 0, 0
	if val, exists := mem[fmt.Sprintf("%v", sub)]; exists {
		a = earn + val
	} else {
		a = earn + deleteAndEarn(sub)
	}

	if val, exists := mem[fmt.Sprintf("%v", nums[1:])]; exists {
		b = earn + val
	} else {
		b = deleteAndEarn(nums[1:])
	}

	mem[fmt.Sprintf("%v", nums)] = max(a, b)
	return max(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func delete(nums []int, target int) ([]int, int) {
	earn := 0
	a, b := target-1, target+1
	result := make([]int, 0)
	for _, num := range nums {
		if num == target {
			earn += num
		} else if num != a && num != b {
			result = append(result, num)
		}
	}

	return result, earn
}

func main() {
	// fmt.Println(deleteAndEarn([]int{3, 4, 2}))
	// fmt.Println(deleteAndEarn([]int{2,2,3,3,3,4}))
	fmt.Println(deleteAndEarn([]int{12, 32, 93, 17, 100, 72, 40, 71, 37, 92, 58, 34, 29, 78, 11, 84, 77, 90, 92, 35, 12, 5, 27, 92, 91, 23, 65, 91, 85, 14, 42, 28, 80, 85, 38, 71, 62, 82, 66, 3, 33, 33, 55, 60, 48, 78, 63, 11, 20, 51, 78, 42, 37, 21, 100, 13, 60, 57, 91, 53, 49, 15, 45, 19, 51, 2, 96, 22, 32, 2, 46, 62, 58, 11, 29, 6, 74, 38, 70, 97, 4, 22, 76, 19, 1, 90, 63, 55, 64, 44, 90, 51, 36, 16, 65, 95, 64, 59, 53, 93}))
}
