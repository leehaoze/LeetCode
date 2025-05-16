package main

import (
	"fmt"
)

// merge 将两个传入的非递减数组，合并到 nums1
func merge(nums1 []int, m int, nums2 []int, n int) {
	// 直接合并牵扯到数值移动操作，因此我们进行逆序合并
	nums1Idx := m - 1
	nums2Idx := n - 1
	idx := m + n - 1

	for nums1Idx >= 0 || nums2Idx >= 0 {
		// nums1先走空
		if nums1Idx < 0 {
			nums1[idx] = nums2[nums2Idx]
			nums2Idx--
			idx--
			continue
		}

		// nums2 先走空
		if nums2Idx < 0 {
			nums1[idx] = nums1[nums1Idx]
			nums1Idx--
			idx--
			continue
		}

		// 都有值，比较大小
		if nums1[nums1Idx] > nums2[nums2Idx] {
			nums1[idx] = nums1[nums1Idx]
			nums1Idx--
		} else {
			nums1[idx] = nums2[nums2Idx]
			nums2Idx--
		}

		idx--
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	merge(nums1, 3, []int{2, 5, 6}, 3)
	fmt.Println(nums1)
}
