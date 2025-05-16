package main

func rotate(nums []int, k int) {
	ret := make([]int, len(nums))

	if k > len(nums) {
		k = k % len(nums)
	}

	idx := 0

	for i := len(nums) - k; i < len(nums); i++ {
		ret[idx] = nums[i]
		idx++
	}

	for idx < len(nums) {
		ret[idx] = nums[idx-k]
		idx++
	}

	for idx, val := range ret {
		nums[idx] = val
	}

	// fmt.Println(nums)
}

func main() {
	rotate([]int{-1}, 3)
}
