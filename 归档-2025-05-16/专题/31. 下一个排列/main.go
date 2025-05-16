package main

import (
	"fmt"
)

// nextPermutation 将nums视为一个整体，返回当前输入nums的下一个字典序排列
func nextPermutation(nums []int) {
	// 逆序遍历，寻找比当前元素还大的
	// 要找到一个最小的，比当前元素大的，进行交换
	for i := len(nums) - 2; i >= 0; i-- {
		possibleIdx, possibleVal := 0, 0
		for j := 1; i+j < len(nums); j++ {
			if nums[i+j] > nums[i] {
				// 记录
				if possibleVal == 0 || possibleVal >= nums[i+j] {
					possibleIdx = i + j
					possibleVal = nums[i+j]
				}
			}
		}

		if possibleIdx != 0 {
			nums[i], nums[possibleIdx] = nums[possibleIdx], nums[i]
			// 交换完毕后，剩余部分一定是降序的，直接逆转
			for j, k := i+1, len(nums)-1; j < k; {
				nums[j], nums[k] = nums[k], nums[j]
				j++
				k--
			}
			return
		}
	}

	// 这是最后一个字典序，整体翻转
	for i, j := 0, len(nums)-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func main() {
	nums := []int{2, 3, 1, 3, 3}
	nextPermutation(nums)
	fmt.Println(nums)
}
