package main

import "fmt"

// func nextGreaterElement(nums1 []int, nums2 []int) []int {
// 	type pos = int
// 	type val = int
// 	mapB := make(map[val]pos)
//
// 	for i, v := range nums2 {
// 		mapB[v] = i
// 	}
//
// 	ret := make([]int, len(nums1))
// 	for i, v := range nums1 {
// 		pos := mapB[v] + 1
// 		if pos >= len(nums2) {
// 			ret[i] = -1
// 		} else {
// 			found := false
// 			for j := pos; j < len(nums2); j++ {
// 				if nums2[j] > v {
// 					found = true
// 					ret[i] = nums2[j]
// 					break
// 				}
// 			}
// 			if !found {
// 				ret[i] = -1
// 			}
// 		}
// 	}
// 	return ret
// }

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	mp := make(map[int]int)
	stack := make([]int, 0)
	for i := len(nums2) - 1; i >= 0; i-- {
		n := nums2[i]
		for len(stack) > 0 && n >= stack[len(stack)-1] {
			// 出栈
			stack = stack[:len(stack)-1]
		}
		// 存储结果
		if len(stack) > 0 {
			mp[n] = stack[len(stack) - 1]
		} else {
			mp[n] = -1
		}

		stack = append(stack, n)
	}

	ret := make([]int, len(nums1))
	for i, v := range nums1 {
		ret[i] = mp[v]
	}
	return ret
}

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
}
