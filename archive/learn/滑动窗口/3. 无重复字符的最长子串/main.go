package main

import "fmt"

// abcabcbb

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	m := make(map[uint8]int)
	start, end := 0, 0
	ret := 0
	for end < len(s) {
		m[s[end]]++
		repeated := true
		if val, _ := m[s[end]]; val == 1 {
			repeated = false
		}

		// for i := start; i <= end; i++ {
		// 	fmt.Print(string(s[i]))
		// }
		// fmt.Print("\n")
		for (repeated || end == len(s)-1) && (start < end) {
			if ret <= (end - start) {
				ret = end - start
				if !repeated && end == len(s)-1 {
					ret++
				}
			}

			m[s[start]]--
			if val, _ := m[s[start]]; val == 1 {
				repeated = false
			}
			start++
		}
		end++
	}

	return ret
}

func main() {
	// fmt.Println(lengthOfLongestSubstring("pwwkew"))
	// fmt.Println(lengthOfLongestSubstring("au"))
	// fmt.Println(lengthOfLongestSubstring("bbbbbbbbb"))
	fmt.Println(lengthOfLongestSubstring("aab"))
}
