package main

//func merge(nums1 []int, m int, nums2 []int, n int) {
//	i, j := 0, 0
//
//	appendVal := 0
//	// 双指针遍历两个数组
//	for j < n {
//		// nums1 先走完，直接将nums2中的值拷贝进去
//		if i >= m+appendVal {
//			nums1[i] = nums2[j]
//			i++
//			j++
//			//fmt.Println(fmt.Sprintf("合并: %#v", nums1))
//			continue
//		}
//
//		if nums1[i] <= nums2[j] {
//			i++
//		} else {
//			// nums[i]及nums[i]以后的元素向后移动
//			for temp := m + appendVal - 1; temp >= i; temp-- {
//				nums1[temp+1] = nums1[temp]
//			}
//
//			// nums2的元素放到nums1中
//			nums1[i] = nums2[j]
//			i++
//			j++
//			appendVal++
//
//			//fmt.Println(fmt.Sprintf("后移: %#v", nums1))
//		}
//	}
//
//	fmt.Println(fmt.Sprintf("%#v", nums1))
//}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1

	// 合并后数组索引
	newIdx := len(nums1) - 1

	for i >= 0 || j >= 0 {
		if i < 0 {
			nums1[newIdx] = nums2[j]
			j--
			newIdx--
			continue
		}

		if j < 0 {
			nums1[newIdx] = nums1[i]
			i--
			newIdx--
			continue
		}

		if nums1[i] > nums2[j] {
			nums1[newIdx] = nums1[i]
			i--
		} else {
			nums1[newIdx] = nums2[j]
			j--
		}

		newIdx--
	}

	//fmt.Println(fmt.Sprintf("%#v", nums1))
}

func main() {
	merge([]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3)
}
