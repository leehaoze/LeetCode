package main

// findLength 返回nums1和nums2 公共的、长度最长的子数组长度
func findLength(nums1 []int, nums2 []int) int {
	// nums2 转为哈希表 val => [idx1, idx2]
	nums2Map := make(map[int][]int)
	for idx, val := range nums2 {
		nums2Map[val] = append(nums2Map[val], idx)
	}

	ret := 0
	// 遍历数组1，以i为起点，寻找公共子数组
	for i := 0; i < len(nums1); {
		// 以 i 为起点，开始搜索
		if _, exists := nums2Map[nums1[i]]; !exists {
			// nums2 中直接没有当前元素
			i++
			continue
		}

		ret = max(ret, 1)
		for _, nums2StartPos := range nums2Map[nums1[i]] {
			for j := 1; i+j < len(nums1) && nums2StartPos+j < len(nums2) && nums1[i+j] == nums2[nums2StartPos+j]; j++ {
				ret = max(ret, j+1)
			}
		}

		i++
	}

	return ret
}
