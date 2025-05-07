package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	scanIdx1, scanIdx2 := m-1, n-1
	newPos := m + n - 1

	for scanIdx1 >= 0 || scanIdx2 >= 0 {
		if scanIdx1 < 0 {
			nums1[newPos] = nums2[scanIdx2]
			newPos--
			scanIdx2--
			continue
		}

		if scanIdx2 < 0 {
			nums1[newPos] = nums1[scanIdx1]
			newPos--
			scanIdx1--
			continue
		}

		if nums1[scanIdx1] > nums2[scanIdx2] {
			nums1[newPos] = nums1[scanIdx1]
			scanIdx1--
		} else {
			nums1[newPos] = nums2[scanIdx2]
			scanIdx2--
		}
		newPos--
	}
}
