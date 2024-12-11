package main

import (
	"fmt"
)

// removeElement 将移除`nums`中，值为`val`的元素，返回值为`nums`中与`val`相同的元素数量
func removeElement(nums []int, val int) int {
	// newer 指向新数组位置，older指向旧数组的扫描位置
	newer, older := 0, 0

	// 开始扫描旧数组，如果不是 val，则复制到新数组（newer+1),如果是 val，则不需要复制到新数组(newer不动)
	for older < len(nums) {
		if nums[older] != val {
			nums[newer] = nums[older]
			newer++
		}
		older++
	}

	return newer
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	// 0, 1, 4, 2, 3, 0 ,2 ,2
	// 0, 1, 4, 0, 3, 2, 2, 2
	count := removeElement(nums, 2)
	for i := 0; i < count; i++ {
		fmt.Println(nums[i])
	}
}
