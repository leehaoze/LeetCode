package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}

	// 初始化递推数组
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		// 长度为一的字符串一定是回文串
		dp[i][i] = true
	}

	var ret, begin int
	// 必须比1大才可以
	ret = 1

	// 枚举 子串长度
	for subLen := 2; subLen <= len(s); subLen++ {
		// 枚举左边界
		for left := 0; left < len(s); left++ {
			// 计算是否是回文串
			right := left + subLen - 1
			if right >= len(s) {
				// 右侧超出边界，退出当前循环，进入到下一个 子串长度枚举
				break
			}

			if s[left] == s[right] {
				// 两个值相等
				if right-left <= 2 {
					// 字符串长度为2，直接为true
					dp[left][right] = true
				} else {
					// 递推
					dp[left][right] = dp[left+1][right-1]
				}
			}

			// 如果是回文串，更新长度
			if dp[left][right] && ret < right-left+1 {
				ret = right - left + 1
				begin = left
			}
		}
	}

	// fmt.Println(begin, ret)
	return s[begin : begin+ret]
}

func main() {
	fmt.Println(longestPalindrome("babad"))
}
