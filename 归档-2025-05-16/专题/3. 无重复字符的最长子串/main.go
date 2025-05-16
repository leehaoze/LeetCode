package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	// 窗口的左右边界
	left, right := 0, 0

	// 准备一个哈希表
	memory := map[uint8]struct{}{}

	check := func(char uint8) bool {
		_, exists := memory[char]
		return exists
	}

	var ret int

	// 外部循环扩充右边界
	for right < len(s) {
		// 内循环收缩左边界，直到符合条件
		for left <= right && check(s[right]) {
			// 操作map
			delete(memory, s[left])

			left++
		}

		// 计算长度
		if ret < (right - left + 1) {
			ret = right - left + 1
		}

		memory[s[right]] = struct{}{}
		right++
	}

	return ret
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
