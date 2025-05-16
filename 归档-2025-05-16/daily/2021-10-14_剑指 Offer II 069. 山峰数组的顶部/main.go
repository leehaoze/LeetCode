package main

import "fmt"

func peakIndexInMountainArray(arr []int) int {
	// 二分查找到一个数 他的左右都比他小
	// 二分查找时，遇到一个数，他的右侧比他小，左侧比他大， 那么需要在左侧继续查找

	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2
		if arr[mid] > arr[mid+1] && arr[mid] > arr[mid-1] {
			return mid
		} else if arr[mid] > arr[mid+1] && arr[mid] < arr[mid-1] {
			right = mid
		} else {
			left = mid
		}
	}

	return 0
}

func main() {
	fmt.Println(peakIndexInMountainArray([]int{0,1, 0}))
	fmt.Println(peakIndexInMountainArray([]int{1,3,5,3,2}))
	fmt.Println(peakIndexInMountainArray([]int{0,10,5,2}))
	fmt.Println(peakIndexInMountainArray([]int{3,4,5,1}))
	fmt.Println(peakIndexInMountainArray([]int{24,69,100,99,79,78,67,36,26,19}))
}
