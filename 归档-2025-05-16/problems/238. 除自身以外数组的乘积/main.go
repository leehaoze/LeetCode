package main

import "fmt"

func productExceptSelf(nums []int) []int {
	prefixMulti := make([]int, len(nums))
	reversePrefixMulti := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			prefixMulti[i] = nums[i]
			reversePrefixMulti[i] = nums[len(nums)-1]
			continue
		}

		prefixMulti[i] = prefixMulti[i-1] * nums[i]
		reversePrefixMulti[i] = reversePrefixMulti[i-1] * nums[len(nums)-1-i]
	}

	// fmt.Println(prefixMulti)
	// fmt.Println(reversePrefixMulti)

	ret := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		prefix, suffix := 1, 1

		// fmt.Println(fmt.Sprintf("i: %v, i-1: %v, len(nums)-1-i: %v", i, i-1, len(nums)-1-i))
		if i != 0 {
			prefix = prefixMulti[i-1]
		}

		if i != len(nums)-1 {
			suffix = reversePrefixMulti[len(nums)-1-i-1]
		}

		ret[i] = prefix * suffix
	}

	return ret
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))

	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}
