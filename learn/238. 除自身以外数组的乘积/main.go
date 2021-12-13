package main

import "fmt"

// func productExceptSelf(nums []int) []int {
// 	ret := make([]int, len(nums))
// 	for i := 0; i < len(nums); i++ {
// 		// 初始化 当前值
// 		ret[i] = 1
// 		for j := 0; j < i; j++ {
// 			ret[i] = nums[j] * ret[i]
// 			ret[j] = ret[j] * nums[i]
// 		}
// 	}
//
// 	return ret
// }
func productExceptSelf(nums []int) []int {
	l := make([]int, len(nums))
	r := make([]int, len(nums))
	l[0] = 1
	r[len(nums)-1] = 1

	for i := 1; i < len(nums); i++ {
		l[i] = nums[i-1] * l[i-1]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		r[i] = nums[i+1] * r[i+1]
	}

	// fmt.Println(l)
	// fmt.Println(r)
	ret := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ret[i] = l[i] * r[i]
	}

	return ret
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
