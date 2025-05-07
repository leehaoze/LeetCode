package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	leftVal, rightVal := 0, 0
	aStart, bStart := 0, 0
	for i := 0; i <= (len(nums1)+len(nums2))/2; i++ {
		leftVal = rightVal
		if aStart < len(nums1) && (bStart >= len(nums2) || nums1[aStart] < nums2[bStart]) {
			rightVal = nums1[aStart]
			aStart++
		} else {
			rightVal = nums2[bStart]
			bStart++
		}
	}

	if (len(nums1)+len(nums2))%2 == 0 {
		return float64(leftVal+rightVal) / 2
	}

	return float64(rightVal)
}
