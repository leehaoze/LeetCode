package main

// removeElement 移除数组中值为val的元素
func removeElement(nums []int, val int) int {
	newer, older := 0, 0
	for older < len(nums) {
		if nums[older] != val {
			nums[newer] = nums[older]
			newer++
		}

		older++
	}

	// 由于会自增一次，因此不需要再次+1
	return newer
}
