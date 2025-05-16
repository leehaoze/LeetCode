package main

func isArraySpecial(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if !isSpecial(nums[i], nums[i+1]) {
			return false
		}
	}

	return true
}

func isSpecial(a, b int) bool {
	return (a % 2) != (b % 2)
}
