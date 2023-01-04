package main

import (
	"fmt"
)

func minWindow(s string, t string) string {
	m := make(map[rune]int)
	for _, c := range t {
		m[c]++
	}
	cnt := make(map[rune]int)

	check := func(m map[rune]int, cnt map[rune]int) bool {
		for k, v := range m {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}

	start, end := 0, 0
	retStart, retEnd := 0, 0
	for end < len(s) {
		if val, exists := m[rune(s[end])]; exists {
			if val > 0 {
				cnt[rune(s[end])]++
			}
		}

		for check(m, cnt) && start <= end {
			if end-start+1 < retEnd-retStart || retEnd == 0 {
				retStart = start
				retEnd = end + 1
			}

			// 缩小窗口
			if _, exists := m[rune(s[start])]; exists {
				cnt[rune(s[start])]--
			}

			start++
		}

		end++
	}

	return s[retStart:retEnd]

}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("ABCA", "AAB"))
	fmt.Println(minWindow("A", "A"))
	fmt.Println(minWindow("A", "AA"))
	fmt.Println(minWindow("aaaaaaaaaaaabbbbbcdd", "abcdd"))
	fmt.Println(minWindow("aab", "aab"))
}
