package main

func moveZeroes(nums []int) {
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left] = nums[right]
			left++
		}
		right++
	}

	for i := left; i < len(nums); i++ {
		nums[i] = 0
	}

}

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
}
