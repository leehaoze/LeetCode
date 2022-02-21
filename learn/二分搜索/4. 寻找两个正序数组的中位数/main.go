package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	nums := make([]int, len1+len2)

}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
}
