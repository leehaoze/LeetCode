package main

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		// 从 i 开始，放到正确的位置
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	// 找到第一个 nums[i] != i + 1 位置的
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}
