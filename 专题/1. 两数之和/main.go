package main

func twoSum(nums []int, target int) []int {
	// 建立哈希表
	valToIdx := make(map[int]int)

	for i, val := range nums {
		if idx, exists := valToIdx[target-val]; exists {
			return []int{i, idx}
		}

		valToIdx[val] = i
	}

	return nil
}
